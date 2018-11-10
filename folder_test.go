package wrike

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestFolderService_GetFolder(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/v4/folders/IEAAT56UI4HBWMRG", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":"IEAAT56UI4HBWMRG"}`)
	})
	want := &Folder{ID: "IEAAT56UI4HBWMRG"}

	folder, _, err := client.Folders.GetFolder("IEAAT56UI4HBWMRG")
	if err != nil {
		t.Fatalf("Folders.GetFolder returns an error: %v", err)
	}

	if !reflect.DeepEqual(want, folder) {
		t.Errorf("Folders.GetFolder returned %+v, want %+v", folder, want)
	}
}
