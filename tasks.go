package main

import (
	"bytes"
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
	CompensationRate float32 `json:"compensationRate"`
	Project          Project `json:"project"`
}

type tasks []Task

func (p tasks) bytesBuffer() *bytes.Buffer {
	b, _ := json.Marshal(p)
	return bytes.NewBuffer(b)
}

// GetTasks fetches all the tasks available to the authenticated user
func (alvtimeClient *AlvtimeClient) GetTasks() ([]Task, error) {
	payload := payload{}
	req, err := http.NewRequest(
		"GET",
		alvtimeClient.domain+"/api/user/tasks",
		payload.bytesBuffer(),
	)
	if err != nil {
		return nil, err
	}

	byteArr, err := alvtimeClient.do(req)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	json.Unmarshal(byteArr, &tasks)

	return tasks, nil
}

// Edit favorite tasks by setting the Favorite attribute on and passing it in a slice.
// The edited tasks are returned as stored to confirm.
func (alvtimeClient *AlvtimeClient) EditFavoriteTasks(tasksToEdit []Task) ([]Task, error) {
	payload := tasks(tasksToEdit)

	req, err := http.NewRequest(
		"POST",
		alvtimeClient.domain+"/api/user/tasks",
		payload.bytesBuffer(),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	byteArr, err := alvtimeClient.do(req)
	if err != nil {
		return nil, err
	}

	var editedTasks []Task
	err = json.Unmarshal(byteArr, &editedTasks)
	if err != nil {
		return nil, err
	}

	return editedTasks, nil
}
