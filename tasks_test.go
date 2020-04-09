package main

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

func TestGetTasks(t *testing.T) {
	alvtimeClient := AlvtimeClient{
		domain:     "http://dotnet-backend",
		httpClient: &http.Client{},
	}

	tasks, err := alvtimeClient.GetTasks()
	if err != nil {
		log.Fatalln(err)
	}

	length := len(tasks)
	expected := 17
	if length != expected {
		t.Error("Length of tasks array does not match expected")
		t.Errorf("Fetched tasks array length: %v", length)
		t.Errorf("Expected tasks array length: %v", expected)
	}

	task := tasks[5]
	expectedTask := Task{
		Id:               6,
		Name:             "Testradgiver",
		Description:      "",
		Favorite:         false,
		Locked:           false,
		CompensationRate: 0,
		Project: Project{
			Id:   3,
			Name: "Sklier",
			Customer: Customer{
				Id:             2,
				Name:           "Rutsjebaner AS",
				InvoiceAddress: "Alvvegen 21",
				ContactPerson:  "Willy",
				ContactEmail:   "willy@rutsjebaner.no",
				ContactPhone:   "53153162",
			},
		},
	}

	if task != expectedTask {
		indentedTask, _ := json.MarshalIndent(task, "", "    ")
		indentedExpectedTask, _ := json.MarshalIndent(expectedTask, "", "    ")
		t.Errorf("Resived: \n%v", string(indentedTask))
		t.Errorf("Expected: \n%v", string(indentedExpectedTask))
	}
}
