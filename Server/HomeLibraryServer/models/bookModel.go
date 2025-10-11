package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Genre struct {
	GenreID   int    `bson:"genre_id" json:"genre_id" validate:"required"`
	GenreName string `bson:"genre_name" json:"genre_name" validate:"required,min=2,max=100"`
}

type Book struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	BookID      string        `bson:"book_id" json:"book_id" validate:"required"`
	Title       string        `bson:"title" json:"title" validate:"required,min=2,max=500"`
	Author      string        `bson:"author" json:"author" validate:"required"`
	Price       int           `bson:"price" json:"price" validate:"required"`
	CoverPath   string        `bson:"cover_path" json:"cover_path" validate:"required"`
	Language    string        `bson:"language" json:"language" validate:"required"`
	Publication string        `bson:"publication" json:"publication" validate:"required"`
	Genre       []Genre       `bson:"genre" json:"genre" validate:"required,dive"`
}
