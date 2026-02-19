package repositories

import (
	"github.com/ManelFer/APIgo/in/config/database"
	"github.com/ManelFer/APIgo/in/models"
)

func CreateUser(user models.User) error {
	query := "INSERT INTO users (email, password, name) VALUES ($1, $2, $3)"
	_, err := database.DB.Exec(query, user.Email, user.Password, user.Name)
	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.QueryRow("SELECT id, email, password, name FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllusers() ([]models.User, error) {
	rows, err := database.DB.Query("SELECT id, email, password, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
