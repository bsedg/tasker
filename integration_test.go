// +build integration

package tasker

import (
	"net/http"
	"testing"

	"github.com/bsedg/irest"
)

var baseURL string

func init() {
	baseURL = "http://taskservice:80"
}

func GetAllTasksTestHelper(t *irest.Test) *irest.Test {
	return t.AddHeader("Content-Type", "application/json").
		Get(baseURL, "/tasks")
}

func CreateTaskTestHelper(t *irest.Test, task *Task) *irest.Test {
	return t.AddHeader("Content-Type", "application/json").
		Post(baseURL, "/tasks", task)
}

func TestTaskService(t *testing.T) {
	inTask := &Task{
		Name:          "test",
		Action:        "query something",
		ScheduledTime: "* 1 * * *",
	}

	overallTest := irest.NewTest("tasks test")
	getTest := overallTest.NewTest("get test")
	GetAllTasksTestHelper(getTest).
		MustStatus(http.StatusOK)

	for _, each := range overallTest.Tests {
		if each.Error != nil {
			t.Errorf("%s: %s for %s", each.Name, each.Error, each.Endpoint)
		}
	}
}
