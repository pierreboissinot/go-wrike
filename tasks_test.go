package wrike

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestTasksService(t *testing.T) {
	client := NewClient(nil, os.Getenv("WRIKE_API_TOKEN"))
	want := &TasksResponse{Kind: "tasks"}
	permalink := os.Getenv("WRIKE_TEST_PERMALINK")
	fmt.Println("Permalink " + permalink)
	fields := "[\"description\"]"
	tasks, _, err := client.Tasks.QueryTasks(QueryTasksOptions{permalink, fields})
	if err != nil {
		t.Fatalf("Tasks.QueryTasks returns an error: %v", err)
	}
	fmt.Println(tasks.Data[0].Description)

	if !reflect.DeepEqual(want, tasks) {
		t.Errorf("Tasks.QueryTasks returned %+v, want %+v", tasks, want)
	}

	// GetTask
	task, _, err := client.Tasks.GetTask(os.Getenv("WRIKE_TEST_TASK_ID"), GetTaskOptions{})
	fmt.Println("Task " + task.Data[0].Description)
}
