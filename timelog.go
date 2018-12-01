package wrike

import (
	"fmt"
)

// TimelogService is Timelog endpoint, see https://developers.wrike.com/documentation/api/methods/query-timelogs
type TimelogService struct {
	client *Client
}

// Timelog represents a timelog
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

// Timelogs represents /timelogs reponse
type Timelogs struct {
	Kind string    `json:"kind"`
	Data []Timelog `json:"data,omitempty"`
}

// TimelogsParams represents params passed to query timelogs
type TimelogsParams struct {
}

// GetTimelogs get folder timelogs
func (s *TimelogService) GetTimelogs(id string) (*Timelogs, *Response, error) {
	u := fmt.Sprintf("folders/%s/timelogs", id)
	fmt.Println(u)
	req, err := s.client.NewRequest("GET", u, TimelogsParams{})
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
