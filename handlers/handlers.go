package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/davidpolme/momentvision-admin-backend/config"
	"github.com/davidpolme/momentvision-admin-backend/controllers"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/api/programmers", controllers.CreateProgram).Methods("POST")
	router.HandleFunc("/api/highlights", controllers.CreateHighlights).Methods("POST")

	//set port
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = config.PORT
	}
	//allow cors
	fmt.Println("Server started on port " + PORT)
	handler := cors.AllowAll().Handler(router)

	if err := http.ListenAndServe(":"+PORT, handler); err != nil {
		log.Fatal(err)
	}

}
