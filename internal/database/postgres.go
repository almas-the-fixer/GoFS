package database

import(
	"fmt"
	"context"
	"os"
	"github.com/jackc/pgx/v5"
	"gofs/internal/types"
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

func CreateUser(req types.UserCreateRequest){
	//TODO
}