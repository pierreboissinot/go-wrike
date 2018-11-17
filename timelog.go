package wrike

import (
	"fmt"
)

type TimelogService struct {
	client *Client
}

type Timelog struct {
	ID          string  `json:"id"`
	TaskID      string  `json:"taskId"`
	UserID      string  `json:"userId"`
	CategoryID  string  `json:"categoryId"`
	Hours       float64 `json:"hours"`
	CreatedDate string  `json:"createdDate"`
	UpdatedDate string  `json:"updatedDate"`
	TrackedDate string  `json:"trackedDate"`
	Comment     string  `json:"comment"`
}

type Timelogs struct {
	Kind string    `json:"kind"`
	Data []Timelog `json:"data,omitempty"`
}

func (s *TimelogService) GetTimelogs(id string) (*Timelogs, *Response, error) {
	u := fmt.Sprintf("folders/%s/timelogs", id)
	fmt.Println(u)
	req, err := s.client.NewRequest("GET", u)
	if err != nil {
		return nil, nil, err
	}

	timelogs := new(Timelogs)
	resp, err := s.client.Do(req, timelogs)

	if err != nil {
		return nil, resp, err
	}
	return timelogs, resp, err
}
