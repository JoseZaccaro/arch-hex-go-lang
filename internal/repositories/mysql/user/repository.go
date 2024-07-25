package mysql_user

import "gorm.io/gorm"

type UserRepository struct {
	DB *gorm.DB
	// Collection *mongo.Collection
}
