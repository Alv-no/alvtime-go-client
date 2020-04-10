package main

import (
	"testing"
)

func TestGetTimeEntries(t *testing.T) {
	alvtimeClient := createTestAlvtimeClient()

	timeEntries, err := alvtimeClient.GetTimeEntries("2019-01-01", "2020-01-01")
	if err != nil {
		t.Error(err)
	}

	length := len(timeEntries)
	if length <= 0 {
		t.Errorf("Length of the timeEntries array is %v, not above 0", length)
	}
}
