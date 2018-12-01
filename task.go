package wrike

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
}

// QueryTasks query tasks
func (t *TaskService) QueryTasks(permalink string) (*TasksResponse, *Response, error) {
	path := "tasks"
	req, err := t.client.NewRequest("GET", path, QueryTasksOptions{Permalink: permalink})
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
