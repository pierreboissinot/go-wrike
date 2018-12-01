package wrike

import "fmt"

// TaskService is Tasks endpoint, see https://developers.wrike.com/documentation/api/methods/query-tasks
type TaskService struct {
	client *Client
}

// TasksResponse is Response from /tasks query
type TasksResponse struct {
	Kind string `json:"kind"`
	Data []Task `json:"data,omitempty"`
}

// Task is a task
type Task struct {
	ID             string `json:"id"`
	AccountID      string `json:"accountId"`
	Title          string `json:"title"`
	Status         string `json:"status"`
	Importance     string `json:"importance"`
	CreatedDate    Time   `json:"createdDate"`
	UpdatedDate    Time   `json:"updatedDate"`
	Dates          Dates  `json:"dates"`
	Scope          string `json:"scope"`
	CustomStatusID string `json:"customStatusId"`
	Permalink      string `json:"permalink"`
	Priority       string `json:"priority"`
	Description    string `json:"description"`
}

// Dates are dates
type Dates struct {
	Type     string `json:"type"`
	Duration int    `json:"duration"`
	Start    string `json:"start"`
	Due      string `json:"due"`
}

// QueryTasksOptions to add params to query
type QueryTasksOptions struct {
	Permalink string `url:"permalink"`
	Fields    string `url:"fields"`
}

// QueryTasks query tasks
func (t *TaskService) QueryTasks(options QueryTasksOptions) (*TasksResponse, *Response, error) {
	path := "tasks"
	req, err := t.client.NewRequest("GET", path, options)
	if err != nil {
		return nil, nil, err
	}

	tasksResponse := new(TasksResponse)
	response, err := t.client.Do(req, tasksResponse)
	if err != nil {
		return nil, response, err
	}

	return tasksResponse, response, err
}

// TaskResponse represents /tasks/{id} response
type TaskResponse struct {
	Kind string         `json:"kind"`
	Data []DetailedTask `json:"data"`
}

// DetailedTask represents task return fron /tasks/{id} endpoint
type DetailedTask struct {
	ID               string        `json:"id"`
	AccountID        string        `json:"accountID"`
	Title            string        `json:"title"`
	Description      string        `json:"description"`
	BriefDescription string        `json:"briefDescription"`
	ParentIDs        []string      `json:"parentIds"`
	SuperParentIDs   []string      `json:"superParentIds"`
	SharedIDs        []string      `json:"sharedIds"`
	ResponsibleIDs   []string      `json:"responsibleIds"`
	Status           string        `json:"status"`
	Importance       string        `json:"importance"`
	CreatedDate      Time          `json:"createdDate"`
	UpdatedDate      Time          `json:"updatedDate"`
	Dates            Dates         `json:"dates"`
	Scope            string        `json:"scope"`
	AuthorIds        []string      `json:"authorIds"`
	CustomStatusID   string        `json:"customStatusId"`
	HasAttachments   bool          `json:"hasAttachments"`
	Permalink        string        `json:"permalink"`
	Priority         string        `json:"priority"`
	FollowedByMe     bool          `json:"followedByMe"`
	FollowerIDs      []string      `json:"followerIds"`
	SuperTaskIDs     []string      `json:"superTaskIds"`
	SubTaskIDs       []string      `json:"subTaskIds"`
	DependencyIDs    []string      `json:"dependencyIds"`
	Metadata         []Metadata    `json:"metadata"`
	CustomFields     []CustomField `json:"customFields"`
}

// GetTaskOptions represents options for /tasks/{id} endpoint
type GetTaskOptions struct {
}

// GetTask get task by id
func (t *TaskService) GetTask(id string, options GetTaskOptions) (*TaskResponse, *Response, error) {
	path := fmt.Sprintf("tasks/%s", id)
	req, err := t.client.NewRequest("GET", path, options)
	if err != nil {
		return nil, nil, err
	}

	taskResponse := new(TaskResponse)
	response, err := t.client.Do(req, taskResponse)
	if err != nil {
		return nil, response, err
	}

	return taskResponse, response, err
}
