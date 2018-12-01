// +build example

package wrike

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func queryTasksExample() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	wrikec := wrike.NewClient(nil, os.Getenv("WRIKE_API_TOKEN"))
	tasks, _, _ := wrikec.Tasks.QueryTasks("https://www.wrike.com/open.htm?id=294160255")
	fmt.Print(task.Data[0].Title)
}
