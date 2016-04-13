package repository

import (
	"U-Talk/server"
	"U-Talk/server/utilities/sessions"
	"fmt"
	"log"
	"net/http"
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

// AddThread adds a thread to db
func (r Repository) AddThread(thread *datastructures.Thread, categoryName string, request *http.Request) {
	session, _ := mgo.Dial(r.ipAddress)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("u-talk").C("forum")
	query := bson.M{"name": categoryName}
	update := bson.M{"$push": bson.M{"threads": bson.M{"topic": thread.Topic(), "description": thread.Description(), "moderator": sessions.UserName(request), "posts": thread.Posts(), "created": thread.Created(), "iconurl": thread.IconURL()}}}
	err := collection.Update(query, update)
	if err != nil {
		fmt.Println("Adding thread")
		log.Fatal(err)
	}
}

// Threads returns the threads in a category
func (r Repository) Threads(categoryName string) []DbThread {
	session, _ := mgo.Dial(r.ipAddress)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	category := DbCategory{}
	collection := session.DB("u-talk").C("forum")
	err := collection.Find(bson.M{"name": categoryName}).One(&category)
	if err != nil {
		log.Fatal(err)
	}
	return category.Threads
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
	Topic       string
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
	Threads []DbThread
	Name    string
	IconURL string
}

// TotalPosts total posts
func (d DbThread) TotalPosts() int {
	return len(d.Posts)
}
