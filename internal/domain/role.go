package domain

import "time"

type Role struct {
	ID          interface{} `json:"id" bson:"_id"`
	Name        string      `json:"name" bson:"name"`
	Description string      `json:"description" bson:"description"`
	CreatedAt   time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" bson:"updated_at"`
}
type CreateRole struct {
	// ID          interface{} `json:"id" bson:"_id"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
