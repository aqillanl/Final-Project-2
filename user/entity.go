package user

import "time"

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
	Age          int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
