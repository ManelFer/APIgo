package repositories

import (
	"github.com/ManelFer/APIgo/in/config/database"
	"github.com/ManelFer/APIgo/in/models"
)

func GetAllusers() ([]models.User, error) {
	rows, err := database.DB.Query("SELECT id, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
