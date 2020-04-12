package main

import (
	"testing"
)

func TestGetTimeEntries(t *testing.T) {
	c := createTestAlvtimeClient()

	timeEntries, err := c.GetTimeEntries("2019-01-01", "2020-01-01")
	if err != nil {
		t.Error(err)
	}

	length := len(timeEntries)
	if length <= 0 {
		t.Errorf("Length of the timeEntries array is %v, not above 0", length)
	}
}

func TestEditTimeEntries(t *testing.T) {
	c := createTestAlvtimeClient()

	timeEntriesToEdit := []TimeEntrie{
		{"2019-01-27", 7, 2},
		{"2019-01-26", 7.5, 2},
	}

	editedTimeEntries, err := c.EditTimeEntries(timeEntriesToEdit)
	if err != nil {
		t.Error(err)
	}

	length := len(editedTimeEntries)
	expectedLength := 2
	if length != expectedLength {
		t.Errorf("Length of timeEntries array is %v, not the expected %v", length, expectedLength)
	}

	for _, editedTimeEntrie := range editedTimeEntries {
		if editedTimeEntrie.Date == timeEntriesToEdit[0].Date && editedTimeEntrie.Value != timeEntriesToEdit[0].Value {
			t.Errorf("TimeEntrie on 2019-01-27 is %v, not the expected %v", editedTimeEntrie.Value, timeEntriesToEdit[0].Value)
		}
		if editedTimeEntrie.Date == timeEntriesToEdit[1].Date && editedTimeEntrie.Value != timeEntriesToEdit[1].Value {
			t.Errorf("TimeEntrie on 2019-01-27 is %v, not the expected %v", editedTimeEntrie.Value, timeEntriesToEdit[1].Value)
		}
	}
}
