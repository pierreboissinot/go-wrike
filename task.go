package wrike

type TaskService struct {
	client *Client
}

type TasksResponse struct {
	Kind string `json:"kind"`
	Data []Task `json:"data,omitempty"`
}

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
	CustomStatusId string `json:"customStatusId"`
	Permalink      string `json:"permalink"`
	Priority       string `json:"priority"`
}

type Dates struct {
	Type     string `json:"type"`
	Duration int    `json:"duration"`
	Start    string `json:"start"`
	Due      string `json:"due"`
}

type QueryTasksOptions struct {
	Permalink string `url:"permalink"`
}

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
