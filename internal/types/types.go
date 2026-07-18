package types

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID					uuid.UUID
	Username			string
	Email				string
	PasswordHash 		string
	CreatedAt			time.Time
	UpdatedAt			time.Time
}

type UserCreateRequest struct {
	ID					uuid.UUID
	Username			string
	Email				string
	PasswordHash 		string
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