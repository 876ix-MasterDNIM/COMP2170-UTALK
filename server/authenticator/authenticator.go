package authenticator

import (
	"U-Talk/server"
	"U-Talk/server/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var db = repository.Repository{}

// AuthenticateUser authenticates a user
func AuthenticateUser(username string, password string) error {
	db.Repository("u-talk", "users")
	user := new(datastructures.User)
	err := db.UserData(username, user)
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}
	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(password))
	if errPassword != nil {
		return fmt.Errorf("Error: %s", err)
	}
	return nil
}
