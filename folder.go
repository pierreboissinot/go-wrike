package wrike

import (
	"fmt"
)

// FolderService endpoint, see Wrike API docs: https://developers.wrike.com/documentation/api/methods/get-folder-tree
type FolderService struct {
	client *Client
}

// Metadata struct
type Metadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// CustomField struct
type CustomField struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// Project struct
type Project struct {
	AuthorID      string   `json:"authorId"`
	OwnerIds      []string `json:"ownerIds"`
	Status        string   `json:"status"`
	StartDate     string   `json:"startDate"`
	EndDate       string   `json:"endDate"`
	CreatedDate   string   `json:"createdDate"`
	CompletedDate string   `json:"completedDate"`
}

// Folder struct, also known as Project
type Folder struct {
	Kind string `json:"kind"`
	Data []struct {
		ID             string        `json:"id"`
		AccountID      string        `json:"accountId"`
		Title          string        `json:"title"`
		CreatedDate    string        `json:"createdAt"`
		UpdatedDate    string        `json:"updatedAt"`
		Description    string        `json:"description"`
		SharedIds      []string      `json:"sharedIds,omitempty"`
		ParentIds      []string      `json:"parentIds,omitempty"`
		ChildIds       []string      `json:"childIds,omitempty"`
		SuperParentIds []string      `json:"superParentIds,omitempty"`
		Scope          string        `json:"scope"`
		HasAttachments bool          `json:"hasAttachments"`
		Permalink      string        `json:"permalink"`
		WorkflowID     string        `json:"workflowId"`
		Metadata       []Metadata    `json:"metadata,omitempty"`
		CustomFields   []CustomField `json:"customFields"`
		Project        Project       `json:"project"`
	} `json:"data,omitempty"`
}

// GetFolder from id, see Wrike API: https://developers.wrike.com/documentation/api/methods/get-folder
func (s *FolderService) GetFolder(id string) (*Folder, *Response, error) {
	u := fmt.Sprintf("folders/%s", id)
	req, err := s.client.NewRequest("GET", u)
	if err != nil {
		return nil, nil, err
	}

	f := new(Folder)
	resp, err := s.client.Do(req, f)

	if err != nil {
		return nil, resp, err
	}
	return f, resp, err
}
