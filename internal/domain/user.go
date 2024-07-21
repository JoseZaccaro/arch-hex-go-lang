package domain

import "time"

//? Entity
type User struct {
	Username     string      `json:"username" bson:"username"`
	Email        string      `json:"email" bson:"email"`
	PasswordHash string      `json:"-" bson:"password_hash"`
	RoleID       interface{} `json:"role_id" bson:"role_id"`
	CreatedAt    time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at" bson:"updated_at"`
}

//? Mongo DTO
type UserMongo struct {
	// * Autogenerated ID / _ID
	ID           interface{} `json:"id" bson:"_id"`
	Username     string      `json:"username" bson:"username"`
	Email        string      `json:"email" bson:"email"`
	PasswordHash string      `json:"-" bson:"password_hash"`
	RoleID       interface{} `json:"role_id" bson:"role_id"`
	CreatedAt    time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at" bson:"updated_at"`
}
type UserLogin struct {
	// * Autogenerated ID / _ID
	ID       interface{} `json:"id" bson:"_id"`
	Username string      `json:"username" bson:"username"`
	Email    string      `json:"email" bson:"email"`
}

type UserCreateParams struct {
	Username     string `json:"username" bson:"username"`
	Email        string `json:"email" bson:"email"`
	PasswordHash string `json:"password" bson:"password_hash"`
	RoleID       string `json:"role_id" bson:"role_id"`
}
