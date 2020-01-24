package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatRoom struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type      string             `json:"type"`
	AvatarURL string             `json:"avatar_url"`
	From      string             `json:"from"`
	Messages  []Message          `json:"messages"`
	Time      time.Time          `json:"time"`
}
