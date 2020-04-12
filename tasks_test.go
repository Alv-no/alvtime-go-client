package main

import (
	"testing"
)

func createTestAlvtimeClient() *AlvtimeClient {
	c, _ := New("http://alvtime-web-api-no-auth")
	return c
}

func TestGetTasks(t *testing.T) {
	c := createTestAlvtimeClient()

	tasks, err := c.GetTasks()
	if err != nil {
		t.Error(err)
	}

	length := len(tasks)
	if length <= 0 {
		t.Errorf("Length of tasks array is %v, not above 0", length)
	}
}

func TestEditFavoriteTasks(t *testing.T) {
	c := createTestAlvtimeClient()

	tasks, err := c.GetTasks()
	if err != nil {
		t.Error(err)
	}

	var tasksToEdit = []Task{tasks[0], tasks[1]}
	tasksToEditCopy := make([]Task, len(tasksToEdit))
	copy(tasksToEditCopy, tasksToEdit)
	tasksToEdit[0].Favorite = !tasksToEdit[0].Favorite
	tasksToEdit[1].Favorite = !tasksToEdit[1].Favorite

	editedTasks, err := c.EditFavoriteTasks(tasksToEdit)
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
			if taskToEditCopy.ID == editedTask.ID {
				if taskToEditCopy.Favorite == editedTask.Favorite {
					t.Errorf("hei")
				}
			}
		}
	}
}
