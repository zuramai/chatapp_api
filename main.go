package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/zuramai/whatsapp_api/controllers"
)

const (
	port = "8080"
)

func main() {
	fmt.Println("Starting server at port ", port, "..")
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/chat", controllers.ChatIndex).Methods("GET")
	router.HandleFunc("/api/v1/chat", controllers.ChatStore).Methods("POST")

	router.HandleFunc("/api/v1/message", controllers.MessageIndex).Methods("GET")
	router.HandleFunc("/api/v1/message", controllers.MessageStore).Methods("POST")

	corsObj := handlers.AllowedOrigins([]string{"*"})
	router.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("."+"/static/"))))

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(corsObj)(router)))

}
