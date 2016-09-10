package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bsedg/tasker"
)

func main() {
	log.Println("taskservice")

	port := os.Getenv("PORT")

	// Register HTTP endpoints with handlers.
	http.HandleFunc("/ping", tasker.PongHandler)
	http.HandleFunc("/version", tasker.VersionHandler)
	http.HandleFunc("/tasks", tasker.TasksHandler)

	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
