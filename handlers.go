package tasker

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// PongHandler simply responds with "pong".
func PongHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong\n")
}

// VersionHandler returns a version file if supplied by the server
// to show build time and build information.
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	versionFile := os.Getenv("VERSION_FILE")
	http.ServeFile(w, r, versionFile)
}

// TasksHandler handles management of tasks.
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data, err := json.Marshal(CreateTask("task", "GET /version", "every 5 minutes"))
		if err != nil {
			log.Printf("error encoding json: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err = w.Write(data); err != nil {
			log.Printf("error writing response: %v", data)
		}
	default:
		log.Printf("%s /tasks", r.Method)
		w.WriteHeader(http.StatusNotImplemented)
	}
}
