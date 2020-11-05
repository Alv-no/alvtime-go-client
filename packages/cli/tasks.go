package alvtimeclient

import (
	"encoding/json"
	"net/url"
)

// Customer struct
type Customer struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	InvoiceAddress string `json:"invoiceAddress"`
	ContactPerson  string `json:"contactPerson"`
	ContactEmail   string `json:"contactEmail"`
	ContactPhone   string `json:"contactPhone"`
}

// Project struct
type Project struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Customer Customer `json:"customer"`
}

// Task struct
type Task struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Favorite         bool    `json:"favorite"`
	Locked           bool    `json:"locked"`
	CompensationRate float32 `json:"compensationRate"`
	Project          Project `json:"project"`
}

// GetTasks fetches all the tasks available to the authenticated user
func (c *AlvtimeClient) GetTasks() (tasks []Task, err error) {
    baseURL, err := url.Parse(c.domain)
    if err != nil {
        return nil, err
    }
	baseURL.Path += "/api/user/tasks"

	req, err := c.newRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	byteArr, err := c.do(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteArr, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// EditFavoriteTasks is used to edit favorite tasks by setting the Favorite
// attribute on and passing it in a slice. The edited tasks are returned as stored to confirm.
func (c *AlvtimeClient) EditFavoriteTasks(tasksToEdit []Task) (editedTasks []Task, err error) {
    baseURL, err := url.Parse(c.domain)
    if err != nil {
        return nil, err
    }
    baseURL.Path += "/api/user/tasks"

	req, err := c.newRequest("POST", baseURL, tasksToEdit)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	byteArr, err := c.do(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteArr, &editedTasks)
	if err != nil {
		return nil, err
	}

	return editedTasks, nil
}
