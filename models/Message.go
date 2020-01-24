package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID      primitive.ObjectID
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}
