package mysql_user

import (
	"api/autentiacion/internal/domain"
	"fmt"
)

// import (
// 	"api/autentiacion/internal/domain"
// 	"context"
// 	"errors"
// 	"fmt"
// )

func (r UserRepository) FindAllUsers() ([]*domain.UserDB, error) {
	//TODO: find all users
	result := r.DB.Table("users")
	fmt.Println(result)
	return nil, nil
}
func (r UserRepository) FindUserByUsername(username string) (*domain.UserDB, error) {
	return nil, nil
}
func (r UserRepository) FindUserByEmail(email string) (*domain.UserDB, error) {
	return nil, nil
}
func (r UserRepository) FindUserById(id string) (*domain.UserDB, error) {
	return nil, nil
}
func (r UserRepository) ExistsByEmail(email string) (bool, error) {
	return true, nil
}
func (r UserRepository) ExistsById(id string) (bool, error) {
	return true, nil
}
func (r UserRepository) ExistsByUsername(username string) (bool, error) {
	return true, nil
}
