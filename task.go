package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID    primitive.ObjectID
	Title string
	Body  string
}
