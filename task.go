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
	Dates          []Date `json:"dates"`
	Scope          string `json:"scope"`
	CustomStatusId string `json:"customStatusId"`
	Permalink      string `json:"permalink"`
	Priority       string `json:"priority"`
}

type Date struct {
	Type     string `json:"type"`
	Duration int    `json:"duration"`
	Start    Time   `json:"start"`
	Due      Time   `json:"due"`
}

func (t *TaskService) QueryTasks(permalink string) (*TasksResponse, *Response, error) {
	path := "tasks"
	req, err := t.client.NewRequest("GET", path)
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
