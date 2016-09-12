package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bsedg/tasker"

	_ "github.com/go-sql-driver/mysql"
)

const (
	TaskerAuthHeader = "X-Tasker-Authentication"
)

func main() {
	var (
		authKey       = os.Getenv("AUTH_KEY")
		port          = os.Getenv("PORT")
		mysqlDatabase = os.Getenv("MYSQL_DATABASE")
		mysqlPassword = os.Getenv("MYSQL_PASSWORD")
		mysqlUser     = os.Getenv("MYSQL_USER")
		mysqlHost     = os.Getenv("MYSQL_HOST")
		mysqlPort     = os.Getenv("MYSQL_PORT")
		versionFile   = os.Getenv("VERSION_FILE")
	)

	// For parsing time properly from sql to go, add parseTime=true.
	// https://github.com/go-sql-driver/mysql#timetime-support
	dbConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)
	db, err := sql.Open("mysql", dbConnStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	ctx := &tasker.TaskerContext{
		Tasks: &tasker.TaskStore{db},
	}

	// Register helper HTTP endpoints with handlers.
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		err = db.Ping()
		if err != nil {
			log.Println(err)
			io.WriteString(w, "error: connection to database\n")
			return
		}

		io.WriteString(w, "pong\n")
	})

	// TODO: use proper status codes, responses, etc.
	http.HandleFunc("/db/init", func(w http.ResponseWriter, r *http.Request) {
		token, ok := r.Header[TaskerAuthHeader]
		if !ok {
			io.WriteString(w, "missing authentication: "+TaskerAuthHeader+"\n")
			return
		}

		if token[0] != authKey {
			io.WriteString(w, "wrong authentication token\n")
			return
		}

		if err := tasker.InitDatabase(db); err != nil {
			log.Println(err)
			io.WriteString(w, "error: connection to database\n")
			return
		}
		io.WriteString(w, "database initialized\n")
	})
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, versionFile)
	})

	// Register tasker HTTP endpoints with handlers.
	http.HandleFunc("/tasks", tasker.NewHandler(ctx, tasker.TasksHandler))

	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
