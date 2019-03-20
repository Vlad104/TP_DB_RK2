package main

import (
	"time"
	"net/http"
)

type Forum struct {
	Posts   int64  `json:"posts,omitempty"`
	Slug    string `json:"slug"`
	Threads int32  `json:"threads,omitempty"`
	Title   string `json:"title"`
	User    string `json:"user"`
}

func (f *Forum) SetForum(r *http.Request) int {
	f.Slug = r.ForumValue("slug")
	f.Title = r.ForumValue("title")
	f.User = r.ForumValue("user")
}

type Post struct {
	Author   string    `json:"author"`
	Created  time.Time `json:"created,omitempty"`
	Forum    string    `json:"forum,omitempty"`
	Id       int64     `json:"id,omitempty"`
	IsEdited bool      `json:"isEdited,omitempty"`
	Message  string    `json:"message"`
	Parent   int64     `json:"parent,omitempty"`
	Thread   int32     `json:"thread,omitempty"`
}

type PostFull struct {
	Author *User   `json:"author,omitempty"`
	Forum  *Forum  `json:"forum,omitempty"`
	Post   *Post   `json:"post,omitempty"`
	Thread *Thread `json:"thread,omitempty"`
}

type PostUpdate struct {
	Message string `json:"message,omitempty"`
}

type Posts []*Post

type Status struct {
	Forum  int32 `json:"forum"`
	Post   int64 `json:"post"`
	Thread int32 `json:"thread"`
	User   int32 `json:"user"`
}

type Thread struct {
	Author  string    `json:"author"`
	Created time.Time `json:"created,omitempty"`
	Forum   string    `json:"forum,omitempty"`
	Id      int32     `json:"id,omitempty"`
	Message string    `json:"message"`
	Slug    string    `json:"slug,omitempty"`
	Title   string    `json:"title"`
	Votes   int32     `json:"votes,omitempty"`
}

type ThreadUpdate struct {
	Message string `json:"message,omitempty"`
	Title   string `json:"title,omitempty"`
}

type Threads []*Thread

type User struct {
	About 		string	`json:"about,omitempty"`
	Email 		string	`json:"email"`
	Fullname 	string	`json:"fullname"`
	Nickname 	string	`json:"nickname,omitempty"`
}

type UserUpdate struct {
	About 		string `json:"about,omitempty"`
	Fullname 	string `json:"fullname,omitempty"`
	Email 		string `json:"email,omitempty"`
}

type Users []*User

type Vote struct {
	Nickname string `json:"nickname"`
	Voice    int32  `json:"voice"`
}