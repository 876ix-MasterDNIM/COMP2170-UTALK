package repository

import (
	"U-Talk/server"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Repository facilitates database transactions
type Repository struct {
	dbName         string
	collectionName string
	ipAddress      string
	port           string
}

// Repository constructor
func (r *Repository) Repository(dbName string, collectionName string) {
	r.dbName = dbName
	r.collectionName = collectionName
}

// UserData fetches user data from database
func (r Repository) UserData(username string, user *datastructures.User) error {
	session, err := mgo.Dial(r.ipAddress)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	dbuser := DbUser{}
	collection := session.DB(r.dbName).C(r.collectionName)
	err = collection.Find(bson.M{"username": username}).One(&dbuser)
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}
	user.User(dbuser.Email, string(dbuser.Password), dbuser.Username, dbuser.ImageURL)
	return nil
}

// StoreUser stores user to database
func (r Repository) StoreUser(user *datastructures.User) {
	session, err := mgo.Dial(r.ipAddress)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("u-talk").C("users")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password()), bcrypt.DefaultCost)
	err = collection.Insert(&DbUser{Username: user.Username(), UserType: user.UserType(), Email: user.Email(), Password: hashedPassword, ImageURL: user.ImageURL()})
	if err != nil {
		log.Fatal(err)
	}
}

// DbUser represents user object in database
type DbUser struct {
	ImageURL string
	Username string
	Email    string
	Password []byte
	UserType string
}

// DbThread represents thread object in database
type DbThread struct {
	Description string
	Moderator   string
	IconURL     string
	Posts       []DbPost
	Created     time.Time
}

// DbPost represents post object in database
type DbPost struct {
	Author  string
	Content string
	Edited  bool
	Created time.Time
}

// DbCategory represents category object in database
type DbCategory struct {
	Thread []DbThread
}
