package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bsedg/tasker"
)

func main() {
	var (
		port        = os.Getenv("PORT")
		versionFile = os.Getenv("VERSION_FILE")
	)

	ctx := &tasker.TaskerContext{
		Tasks: tasker.NewTaskStore(),
	}

	// Register helper HTTP endpoints with handlers.
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong\n")
	})
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, versionFile)
	})

	// Register tasker HTTP endpoints with handlers.
	http.HandleFunc("/tasks", tasker.NewHandler(ctx, tasker.TasksHandler))

	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
