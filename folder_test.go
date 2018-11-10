package wrike

import (
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestFolderService_GetFolder(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/folders/"+os.Getenv("WRIKE_FOLDER"), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
	})

	want := &Folder{Kind: "folders"}
	folder, _, err := client.Folders.GetFolder(os.Getenv("WRIKE_FOLDER"))
	if err != nil {
		t.Fatalf("Folders.GetFolder returns an error: %v", err)
	}

	if !reflect.DeepEqual(want, folder) {
		t.Errorf("Folders.GetFolder returned %+v, want %+v", folder, want)
	}
}
