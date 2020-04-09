package main

import (
	"encoding/json"
	"net/http"
)

type Customer struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	InvoiceAddress string `json:"invoiceAddress"`
	ContactPerson  string `json:"contactPerson"`
	ContactEmail   string `json:"contactEmail"`
	ContactPhone   string `json:"contactPhone"`
}

type Project struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Customer Customer `json:"customer"`
}

type Task struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	Favorite         bool    `json:"favorite"`
	Locked           bool    `json:"locked"`
	CompensationRate int     `json:"compensationRate"`
	Project          Project `json:"project"`
}

// GetTasks fetches all the tasks available to the authenticated user
func (alvtimeClient *AlvtimeClient) GetTasks() ([]Task, error) {
	payload := payload{}
	req, err := http.NewRequest("GET", alvtimeClient.domain+"/api/user/tasks", payload.bytesBuffer())
	if err != nil {
		return nil, err
	}

	byteArr, err := alvtimeClient.do(req)

	var tasks []Task
	json.Unmarshal(byteArr, &tasks)

	return tasks, nil
}
