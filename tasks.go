package tasker

import (
	"fmt"
	"math/rand"

	"gopkg.in/redis.v4"
)

type TaskStore struct {
	Tasks map[int64]*Task
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		Tasks: make(map[int64]*Task),
	}
}

func (ts *TaskStore) Save(t *Task) (*Task, error) {
	if t.ID == 0 {
		t = CreateTask(t.Name, t.Action, t.ScheduledTime)
	}

	ts.Tasks[t.ID] = t

	return t, nil
}

func (ts *TaskStore) GetAll() []*Task {
	tasks := make([]*Task, len(ts.Tasks))

	i := 0
	for _, task := range ts.Tasks {
		tasks[i] = task
		i++
	}
	return tasks
}

func (ts *TaskStore) Get(id int64) *Task {
	t, _ := ts.Tasks[id]

	return t
}

func (ts *TaskStore) Delete(id int64) {
	delete(ts.Tasks, id)
}

// Task encapsulates a named, scheduled task action.
type Task struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Action        string `json:"action"`
	ScheduledTime string `json:"time"`
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
