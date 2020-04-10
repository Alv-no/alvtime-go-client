package main

import (
	"net/http"
	"testing"
)

func createTestAlvtimeClient() AlvtimeClient {
	return AlvtimeClient{
		domain:     "http://alvtime-web-api-no-auth",
		httpClient: &http.Client{},
	}
}

func TestGetTasks(t *testing.T) {
	alvtimeClient := createTestAlvtimeClient()

	tasks, err := alvtimeClient.GetTasks()
	if err != nil {
		t.Error(err)
	}

	length := len(tasks)
	if length < 0 {
		t.Errorf("Length of tasks array is %v, not above 0", length)
	}
}

func TestEditFavoriteTasks(t *testing.T) {
	alvtimeClient := createTestAlvtimeClient()

	tasks, err := alvtimeClient.GetTasks()
	if err != nil {
		t.Error(err)
	}

	var tasksToEdit = []Task{tasks[0], tasks[1]}
	tasksToEditCopy := make([]Task, len(tasksToEdit))
	copy(tasksToEditCopy, tasksToEdit)
	tasksToEdit[0].Favorite = !tasksToEdit[0].Favorite
	tasksToEdit[1].Favorite = !tasksToEdit[1].Favorite

	editedTasks, err := alvtimeClient.EditFavoriteTasks(tasksToEdit)
	if err != nil {
		t.Error(err)
	}

	length := len(editedTasks)
	expectedLength := 2
	if length != expectedLength {
		t.Errorf("Length of tasks array is %v, not the expected %v", length, expectedLength)
	}

	for _, taskToEditCopy := range tasksToEditCopy {
		for _, editedTask := range editedTasks {
			if taskToEditCopy.Id == editedTask.Id {
				if taskToEditCopy.Favorite == editedTask.Favorite {
					t.Errorf("hei")
				}
			}
		}
	}
}
