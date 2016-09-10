package tasker

import (
	"math/rand"
)

// Task encapsulates a named, scheduled task action.
type Task struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Action        string `json:"action"`
	ScheduledTime string `json:"time"`
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
