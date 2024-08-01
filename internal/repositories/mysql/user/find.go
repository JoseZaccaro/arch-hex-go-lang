package mysql_user

import (
	"api/autentiacion/internal/domain"
	"api/autentiacion/internal/repositories"
	"fmt"
	"log"
)

// import (
//
//	"api/autentiacion/internal/domain"
//	"context"
//	"errors"
//	"fmt"
//
// )

func (r UserRepository) FindAllUsers() ([]*domain.UserDB, error) {
	//TODO: find all users
	users := []*domain.UserDB{}
	result := r.DB.Raw("SELECT * FROM users")

	if result.Error != nil {
		return nil, result.Error
	}

	rows, _ := result.Rows()
	// columnNames, _ := rows.Columns()

	for rows.Next() {
		user := domain.UserDB{}
		rows.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
		mapper := &repositories.Mapper{}
		user.ID = mapper.ToInt64(user.ID)
		users = append(users, &user)
	}

	// fmt.Println(result)
	// fmt.Println(users)
	return users, nil
}
func (r UserRepository) FindUserByUsername(username string) (*domain.UserDB, error) {
	return nil, nil
}
func (r UserRepository) FindUserByEmail(email string) (*domain.UserDB, error) {
	result := r.DB.Raw("SELECT * FROM users WHERE LOWER(email) = LOWER(?)", email)

	if result.Error != nil {
		return nil, result.Error
	}

	rows, _ := result.Rows()

	if rows.Next() {
		user := &domain.UserDB{}
		rows.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Email, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
		log.Println(user)
		mapper := &repositories.Mapper{}
		user.ID = mapper.ToInt64(user.ID)
		return user, nil // User found
	}

	return nil, fmt.Errorf("user not found") // User not found
}
func (r UserRepository) FindUserById(id string) (*domain.UserDB, error) {
	return nil, nil
}
func (r UserRepository) ExistsByEmail(email string) (bool, error) {
	result := r.DB.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email)

	if result.Error != nil {
		return false, result.Error
	}

	rows, _ := result.Rows()
	if rows.Next() {
		var exists bool
		rows.Scan(&exists)
		return exists, nil
	}
	return true, nil // User found
}
func (r UserRepository) ExistsById(id string) (bool, error) {
	return true, nil
}
func (r UserRepository) ExistsByUsername(username string) (bool, error) {

	result := r.DB.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", username)
	// result := r.DB.Table("users").Where("username = ?", username).First(&user)

	if result.Error != nil {
		return false, result.Error
	}

	rows, _ := result.Rows()
	if rows.Next() {
		var exists bool
		rows.Scan(&exists)
		return exists, nil
	}
	return true, nil // User found
}
