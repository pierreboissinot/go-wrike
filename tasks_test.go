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
	tasks, _, err := client.Tasks.QueryTasks(permalink)
	if err != nil {
		t.Fatalf("Tasks.QueryTasks returns an error: %v", err)
	}
	fmt.Println(tasks.Data[0].Title)

	if !reflect.DeepEqual(want, tasks) {
		t.Errorf("Tasks.QueryTasks returned %+v, want %+v", tasks, want)
	}
}
