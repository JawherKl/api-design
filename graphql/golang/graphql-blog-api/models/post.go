package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID      string   `json:"id" bson:"_id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Author  string   `json:"author"`
	//Comments []*Comment `json:"comments,omitempty"`
}

type NewPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}