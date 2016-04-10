package datastructures

// User type representing a user
type User struct {
	imageURL string
	username string
	email    string
	password string
	userType string
}

// Username user's name
func (u User) Username() string {
	return u.username
}

// ImageURL user's image url
func (u User) ImageURL() string {
	return u.imageURL
}

// Email user's email
func (u User) Email() string {
	return u.email
}

// Password user's password
func (u User) Password() string {
	return u.password
}

// UserType user's type
func (u User) UserType() string {
	return u.userType
}

//User constructor
func (u *User) User(email string, password string, username string, imageURL string) {
	u.username = username
	u.password = password
	u.email = email
	u.imageURL = imageURL
	u.userType = "regular"
}
