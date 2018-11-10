// +build example

package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-wrike"
	"log"
	"os"
)

func getProjectExample() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	wrikec := wrike.NewClient(nil, os.Getenv("WRIKE_API_TOKEN"))
	folder, _, _ := wrikec.Folders.GetFolder(os.Getenv("WRIKE_FOLDER"))
	fmt.Print(folder.Data[0].Title)
}
