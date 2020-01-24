package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zuramai/whatsapp_api/config"
	"github.com/zuramai/whatsapp_api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func MessageIndex(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	var messages []models.Message

	cursor, err := config.Database().Collection("messages").Find(context.TODO(), bson.M{})

	if err != nil {
		fmt.Println(err)
		return
	}

	for cursor.Next(context.TODO()) {
		cursor.Decode(&message)
		messages = append(messages, message)
		message = models.Message{}
	}

	writeJSON(w, 200, "Success get all messages", messages)
}

func MessageStore(w http.ResponseWriter, r *http.Request) {

}
