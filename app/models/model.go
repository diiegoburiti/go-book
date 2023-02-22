package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Isbn  string             `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Title string             `json:"title" bson:"title,omitempty"`
}
