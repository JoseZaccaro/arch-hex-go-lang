package domain

import "time"

type User struct {
	ID           interface{} `json:"id" bson:"_id"`
	Username     string      `json:"username" bson:"username"`
	Email        string      `json:"email" bson:"email"`
	PasswordHash string      `json:"-" bson:"password_hash"`
	CreatedAt    time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at" bson:"updated_at"`
}

type UserCreateParams struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password_hash"`
}

type UserUpdateParams struct {
	Username string
	Email    string
	Password string
}
