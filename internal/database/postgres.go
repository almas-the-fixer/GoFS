package database

import (
	"context"
	"fmt"
	"gofs/internal/types"
	"gofs/internal/validation"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() (conn *pgx.Conn,err error){
	
	conn , err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	fmt.Println(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func GetUsers(conn *pgx.Conn)(users []types.User){
	result, err := conn.Query(context.Background() , "SELECT * FROM users")
	if err != nil {
		fmt.Println("AN ERROR OCCURED: ", err)
	}
	
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

// PseudoCode:
// 1: Validate User Create Req 2: Insert into Table 3: return user
func CreateUser(conn *pgx.Conn, userReq types.UserCreateRequest)(error){
	err := validation.UserCreateRequestValidator(userReq)
	if err != nil {
		return err
	}
	fmt.Println("Validation Passed")

	_ , err = conn.Exec(context.Background(), `INSERT INTO users(username, email, password_hash) VALUES($1, $2, $3)`, userReq.Username, userReq.Email, userReq.Password)

	if err != nil{
		return err
	}
	return nil
}