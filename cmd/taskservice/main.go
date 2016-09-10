package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func setup() {
	log.Println("taskservice")
}

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/ping", PongHandler)
	http.HandleFunc("/version", VersionHandler)
	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

type TaskerContext struct {
	VersionFile string
}

func PongHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong\n")
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	versionFile := os.Getenv("VERSION_FILE")
	http.ServeFile(w, r, versionFile)
}
