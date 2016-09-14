package tasker

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type TaskStore struct {
	DB *sql.DB
}

// InitDatabase initialized data tables.
// TODO: this may need to live elsewhere.
func InitDatabase(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS tasks")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE tasks (id BIGINT(18) NOT NULL AUTO_INCREMENT, name VARCHAR(255) DEFAULT NULL, action VARCHAR(255) DEFAULT NULL, time VARCHAR(255) DEFAULT NULL, created DATETIME DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (`id`))")
	return err
}

func (ts *TaskStore) Save(t *Task) (*Task, error) {
	var cmd string
	if t.ID == 0 {
		// Create a new task.
		cmd = fmt.Sprintf("INSERT INTO tasks (name, action, time) VALUES ('%s', '%s', '%s')",
			t.Name, t.Action, t.ScheduledTime)
	} else {
		cmd = fmt.Sprintf("UPDATE tasks t SET t.name='%s', t.action='%s', t.time='%s' WHERE id = %d",
			t.Name, t.Action, t.ScheduledTime, t.ID)
	}

	insert, err := ts.DB.Exec(cmd)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	id, err := insert.LastInsertId()
	if t.ID == 0 && err != nil {
		log.Println(err)
	} else if t.ID == 0 {
		t.ID = id
	}

	return t, nil
}

func (ts *TaskStore) GetAll() ([]*Task, error) {
	cmd := fmt.Sprintf("SELECT * from tasks")
	rows, err := ts.DB.Query(cmd)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	tasks := []*Task{}
	for rows.Next() {
		t := &Task{}
		if err := rows.Scan(&t.ID, &t.Name, &t.Action, &t.ScheduledTime, &t.Created); err != nil {
			log.Println(err)
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (ts *TaskStore) Get(id int64) *Task {
	return nil
}

func (ts *TaskStore) Delete(id int64) {

}

// Task encapsulates a named, scheduled task action.
type Task struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Action        string    `json:"action"`
	ScheduledTime string    `json:"time"`
	Created       time.Time `json:"created"`
}

// Valid validates input for a task.
func (t *Task) Valid() error {
	if t.Name == "" || t.Action == "" || t.ScheduledTime == "" {
		return fmt.Errorf("name, action, and scheduled time required")
	}

	return nil
}

// CreateTask simply creates a task with provided parmaters
// and generates a psuedo random number.
func CreateTask(name, action, schedTime string) *Task {
	return &Task{
		ID:            rand.Int63n(10000),
		Name:          name,
		Action:        action,
		ScheduledTime: schedTime,
	}
}
