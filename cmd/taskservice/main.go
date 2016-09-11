package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bsedg/tasker"

	"gopkg.in/redis.v4"
)

func main() {
	var (
		port        = os.Getenv("PORT")
		redisHost   = os.Getenv("REDIS_HOST")
		redisPort   = os.Getenv("REDIS_PORT")
		versionFile = os.Getenv("VERSION_FILE")
	)

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})

	ctx := &tasker.TaskerContext{
		DBClient: client,
		Tasks:    tasker.NewTaskStore(),
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
