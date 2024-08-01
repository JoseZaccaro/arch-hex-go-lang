package domain

import "time"

//? Entity
type User struct {
	Username     string      `json:"username" bson:"username" sql:"username"`
	Email        string      `json:"email" bson:"email" sql:"email"`
	PasswordHash string      `json:"-" bson:"password_hash" sql:"password_hash"`
	RoleID       interface{} `json:"role_id" bson:"role_id" sql:"role_id" gorm:"type:varchar"`
	CreatedAt    time.Time   `json:"created_at" bson:"created_at" sql:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at" bson:"updated_at" sql:"updated_at"`
}

//? DB DTO
type UserDB struct {
	// * Autogenerated ID / _ID
	ID           interface{} `gorm:"primaryKey;type:int" json:"id" bson:"_id" sql:"id"`
	Username     string      `json:"username" bson:"username" sql:"username"`
	Email        string      `json:"email" bson:"email" sql:"email"`
	PasswordHash string      `json:"-" bson:"password_hash" sql:"password_hash"`
	RoleID       interface{} `json:"role_id" bson:"role_id" sql:"role_id" gorm:"type:varchar"`
	CreatedAt    time.Time   `json:"created_at" bson:"created_at" sql:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time   `json:"updated_at" bson:"updated_at" sql:"updated_at" gorm:"autoUpdateTime"`
}
type UserLogin struct {
	// * Autogenerated ID / _ID
	ID       interface{} `gorm:"primaryKey;type:int" json:"id" bson:"_id" sql:"id"`
	Username string      `json:"username" bson:"username" sql:"username"`
	Email    string      `json:"email" bson:"email" sql:"email"`
}

type UserCreateParams struct {
	Username     string      `json:"username" bson:"username" sql:"username"`
	Email        string      `json:"email" bson:"email" sql:"email"`
	PasswordHash string      `json:"password" bson:"password_hash" sql:"password_hash"`
	RoleID       interface{} `json:"role_id" bson:"role_id" sql:"role_id" gorm:"type:varchar"`
}
