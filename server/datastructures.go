package datastructures

import "time"

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

// Post represents Posts in the forum
type Post struct {
	author  string
	content string
	edited  bool
	created time.Time
}

// Author returns the name of the author of the post
func (p Post) Author() string {
	return p.author
}

// Content returns the content of the post
func (p Post) Content() string {
	return p.content
}

// SetContent edits the content contained within a post
func (p *Post) SetContent(newContent string) {
	p.content = newContent
}

// Created returns the datetime of post creation
func (p Post) Created() time.Time {
	return p.created
}

// Edited sets whether the post has been edited or not
func (p *Post) Edited(edited bool) {
	p.edited = edited
}

// WasEdited returns whether post has been edited or not
func (p Post) WasEdited() bool {
	return p.edited
}

// Thread represents a thread in the forum
type Thread struct {
	description string
	moderator   string
	iconURL     string
	topic       string
	posts       []Post
	created     time.Time
}

// Thread constructor
func (t *Thread) Thread(description string, moderator string, iconURL string, topic string) {
	t.description = description
	t.moderator = moderator
	t.iconURL = iconURL
	t.created = time.Now()
	t.posts = []Post{}
	t.topic = topic
}

// TotalPosts returns the number of posts in the thread
func (t Thread) TotalPosts() int {
	return len(t.posts)
}

// Moderator returns the moderator of thread
func (t Thread) Moderator() string {
	return t.moderator
}

// Description return the description of thread
func (t Thread) Description() string {
	return t.description
}

// Created returns the datetime of thread creation
func (t Thread) Created() time.Time {
	return t.created
}

// Topic return topic
func (t Thread) Topic() string {
	return t.topic
}

// IconURL returns the url of thread's icon
func (t Thread) IconURL() string {
	return t.iconURL
}

// Posts returns the thread's posts
func (t Thread) Posts() []Post {
	return t.posts
}

// Category represents a category in the forum
type Category struct {
	threads []Thread
	iconuRL string
	name    string
}

// Threads returns the threads in a category
func (c Category) Threads() []Thread {
	return c.threads
}

// IconURL returns the url of category's icon
func (c Category) IconURL() string {
	return c.iconuRL
}

// Name returns the name of the category
func (c Category) Name() string {
	return c.name
}
