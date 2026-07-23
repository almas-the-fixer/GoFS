package types

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID					uuid.UUID		`json:"id"`
	Username			string			`json:"username"`
	Email				string			`json:"email"`
	PasswordHash 		string			`json:"-"`
	CreatedAt			time.Time		`json:"created_at"`
	UpdatedAt			time.Time		`json:"updated_at"`
}

type UserCreateRequest struct {
	Username			string			`json:"username"`
	Email				string			`json:"email"`
	Password 			string			`json:"password"`
	ConfirmPassword		string			`json:"confirm_password"`
}

type File struct {
	ID					uuid.UUID
	OwnerID				uuid.UUID
	OriginalName		string
	StoredName			string
	Size				int64
	MimeType			string
	StoragePath			string
}

type LoginRequest struct {
	Email				string				`json:"email"`
	Password			string				`json:"password"`
}