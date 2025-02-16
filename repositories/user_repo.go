package repositories

import (
	"database/sql"
	"errors"
	"log"
	"task-manager-api/config"
	"task-manager-api/models"
)

// CreateUser creates a new user
func CreateUser(user *models.User) error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := config.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = uint(id)
	return nil
}

// GetUserByEmail fetches a user by email
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE email = ? LIMIT 1"
	err := config.DB.Get(&user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("ℹ️ User not found:", email)
			return nil, nil
		}
		log.Println("❌ Error querying user:", err)
		return nil, err
	}
	return &user, nil
}
