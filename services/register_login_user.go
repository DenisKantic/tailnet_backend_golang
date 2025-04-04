package services

import (
	"tailnet_backend_go/config/database"
	"tailnet_backend_go/models"
)

func RegisterNewUser(user models.UserRegister) error {

	db, err := database.DB_connect()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO users(username, email, password, profile_image,location, bio) 
			VALUES($1, $2, $3, $4, $5, $6)`
	_, err = db.Exec(query, user.Username, user.Email, user.Password, user.ProfileImageURL,
		user.Location, user.Bio)

	return err
}
