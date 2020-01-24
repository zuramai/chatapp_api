package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/zuramai/whatsapp_api/config"
	"github.com/zuramai/whatsapp_api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChatIndex(w http.ResponseWriter, r *http.Request) {

	var chatRoom models.ChatRoom
	var chatRooms []models.ChatRoom

	cursor, err := config.Database().Collection("chat_room").Find(context.TODO(), bson.M{})

	if err != nil {
		fmt.Println(err)
		return
	}

	for cursor.Next(context.TODO()) {
		cursor.Decode(&chatRoom)
		chatRooms = append(chatRooms, chatRoom)
		chatRoom = models.ChatRoom{}
	}

	writeJSON(w, 200, "Success get all chatRooms", chatRooms)
}

func ChatStore(w http.ResponseWriter, r *http.Request) {
	var chatRoom models.ChatRoom
	err := json.NewDecoder(r.Body).Decode(&chatRoom)
	if err != nil {
		panic(err)
	}
	chatRoom.Time = time.Now()
	res, err := config.Database().Collection("chat_room").InsertOne(context.TODO(), chatRoom)

	if err != nil {
		panic(err)
	}

	writeJSON(w, 200, "Success create new chat room", res)
}
