package main

import (
	"github.com/danielwelch/mednotes/Godeps/_workspace/src/github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}

	r := handlers.LoggingHandler(os.Stdout, Router())

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":"+port, r))

}
