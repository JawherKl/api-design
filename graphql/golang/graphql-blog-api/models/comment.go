package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PostID  primitive.ObjectID `bson:"postId" json:"postId"`
	Author  string             `bson:"author" json:"author"`
	Content string             `bson:"content" json:"content"`
}
