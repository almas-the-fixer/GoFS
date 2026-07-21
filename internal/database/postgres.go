package database

import (
	"context"
	"errors"
	"fmt"
	"gofs/internal/types"
	"gofs/internal/validation"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

func ConnectDB() (conn *pgx.Conn, err error) {

	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	fmt.Println(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func DeleteUser(conn *pgx.Conn, userID string) error {
	_, err := conn.Exec(context.Background(), `DELETE FROM users WHERE id = $1`, userID)
	if err != nil{
		fmt.Println("An Error Occured when Deleting a User: ", err)
		return err
	}
	return nil
}

func GetUser(conn *pgx.Conn, userID string)(user types.User,err error){
	err = conn.QueryRow(context.Background(), `SELECT id, username, email, password_hash, created_at, updated_at FROM users WHERE id = $1`,userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		fmt.Println("An ERROR Occured: ", err)
		return user, err
	}
	return user, nil
}

func GetUsers(conn *pgx.Conn) (users []types.User) {
	result, err := conn.Query(context.Background(), "SELECT id, username, email, password_hash, created_at, updated_at FROM users")
	if err != nil {
		fmt.Println("AN ERROR OCCURED: ", err)
	}
	defer result.Close()

	for result.Next() {
		var user types.User
		err := result.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			fmt.Println("AN ERROR OCCURED WHILE SCANNING ROWS: ", err)
		}
		users = append(users, user)
	}
	return users
}

// Some Notes Ill moove Somewhere else later:
// Exec()     -> Use when you DON'T expect rows back
// QueryRow() -> Use when you expect EXACTLY ONE row back
// Query()    -> Use when you expect MULTIPLE rows back
// PseudoCode:
// 1: Validate User Create Req 2: Insert into Table 3: return user
func CreateUser(conn *pgx.Conn, userReq types.UserCreateRequest) (string, error) {
	var userID string
	err := validation.UserCreateRequestValidator(userReq)
	if err != nil {
		return "", err
	}
	fmt.Println("Handler Validation Passed")

	// Hashing Password Before Storing it
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("An Error Occures while HASHING password: ", err)
		return "", err
	}

	err = conn.QueryRow(context.Background(), `INSERT INTO users(username, email, password_hash) VALUES($1, $2, $3) RETURNING id`, userReq.Username, userReq.Email, hashedPass).Scan(&userID)

	//Exec() returns a CommandTag, which contains information like:
	// INSERT 0 1
	// UPDATE 3
	// DELETE 1
	// _, err = conn.Exec(context.Background(), `INSERT INTO users(username, email, password_hash) VALUES($1, $2, $3) RETURNING id`, userReq.Username, userReq.Email, hashedPass)

	if err != nil {
		var pgError *pgconn.PgError	
		if errors.As(err, &pgError){
			fmt.Println(pgError.Message)
			fmt.Println(pgError.Code)
		}
		return "", err
	}
	return  userID,nil
}