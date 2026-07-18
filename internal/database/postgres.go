package database

import(
	"fmt"
	"context"
	"github.com/joho/godotenv"
	"os"
	"github.com/jackc/pgx/v5"
	"gofs/internal/types"
)

func ConnectDB() (conn *pgx.Conn){
	if err := godotenv.Load(); err != nil {
		fmt.Println("AN ERROR OCCURED LOADING THE ENV VARIABLES: ", err)
	}
	conn , err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	fmt.Println(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func GetUsers(){
	conn := ConnectDB()
	result, err := conn.Query(context.Background() , "SELECT id, username, email, password_hash FROM users")
	if err != nil {
		fmt.Println("AN ERROR OCCURED: ", err)
	}
	fmt.Println("ROWS QUERIED HERE IS THE CURSOR: ~> \n", result)

	fmt.Println("Getting Values via pgx.Scan()")
	var users []types.User
	
	for result.Next() {
		var user types.User	   
		err := result.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.PasswordHash,
		)

		if err != nil {
			fmt.Println("AN ERROR OCCURED WHILE SCANNING ROWS: ", err)
		}
		users = append(users, user)
	}
	fmt.Println("BELOW ARE USERS QUERIED!")
	fmt.Println(users)
}

func CreateUser(UserCreateRequest types.User){
	//TODO
}