package main

import (
	"github.com/gorilla/mux"
	"pp-api/controller"
	"net/http"
	"pp-api/config"
	"github.com/rs/cors"
	"log"
)

const configFile = "resources/config.yml"

func main() {

	config := config.LoadConfig(configFile)

	r := mux.NewRouter()

	controller.HandleRequests(r)
	log.Printf("Starting api endpoint on port%s", config.ListeningAddr)

	handler := cors.Default().Handler(r)

	err := http.ListenAndServe(config.ListeningAddr, handler)
	log.Fatal("Error starting server: ", err.Error())
}
