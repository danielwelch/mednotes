package main

import (
	"github.com/danielwelch/mednotes/Godeps/_workspace/src/github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	r := handlers.LoggingHandler(os.Stdout, Router())

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
