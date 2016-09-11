package tasker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type TaskerContext struct {
	Tasks *TaskStore
}

type ErrResponse struct {
	Message string `json:"message"`
}

type TaskerHandler func(*TaskerContext, http.ResponseWriter, *http.Request) (interface{}, int, error)

func NewHandler(ctx *TaskerContext, th TaskerHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		preHandler := time.Now()
		data, status, err := th(ctx, w, r)
		handlerLatency := time.Since(preHandler)

		ms := int64(handlerLatency / time.Millisecond)
		log.Printf("HTTP %d %s %s | %d ms", status, r.Method, r.URL.Path, ms)

		w.WriteHeader(status)

		if err != nil {
			errMsg := &ErrResponse{err.Error()}
			errData, err := json.Marshal(errMsg)
			if err != nil {
				log.Printf("error encoding error response: %s", errMsg)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if _, err = w.Write(errData); err != nil {
				log.Printf("error writing error response: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		if data != nil {
			dataJSON, err := json.Marshal(data)
			if err != nil {
				log.Printf("error encoding data: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if _, err = w.Write(dataJSON); err != nil {
				log.Printf("error writing data: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	})
}

// TasksHandler handles management of tasks.
func TasksHandler(c *TaskerContext, w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	switch r.Method {
	case http.MethodGet:
		return c.Tasks.GetAll(), http.StatusOK, nil
	case http.MethodPost:
		task := &Task{}
		if err := json.NewDecoder(r.Body).Decode(task); err != nil {
			return nil, http.StatusBadRequest, fmt.Errorf("error decoding task")
		}

		if err := task.Valid(); err != nil {
			return nil, http.StatusBadRequest, err
		}

		t, err := c.Tasks.Save(task)
		if err != nil {
			return t, http.StatusInternalServerError, err
		}

		return t, http.StatusCreated, err
	default:
		log.Printf("%s /tasks", r.Method)
		return nil, http.StatusNotImplemented, nil
	}
}
