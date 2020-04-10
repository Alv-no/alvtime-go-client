package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type TimeEntrie struct {
	Id     int     `json:"id"`
	Date   string  `json:"date"`
	Value  float32 `json:"value"`
	TaskId int     `json:"taskId"`
}

type timeEntries []TimeEntrie

func (p timeEntries) bytesBuffer() *bytes.Buffer {
	b, _ := json.Marshal(p)
	return bytes.NewBuffer(b)
}

// GetTimeEntries fetches all the time entries available to the authenticated user
// between the dates denoted by fromDateInclusive to toDateInclusive formated as "YYYY-MM-DD"
func (alvtimeClient *AlvtimeClient) GetTimeEntries(
	fromDateInclusive, toDateInclusive string,
) ([]TimeEntrie, error) {
	payload := payload{}
	req, err := http.NewRequest(
		"GET",
		alvtimeClient.domain+"/api/user/timeentries"+
			"?fromDateInclusive="+fromDateInclusive+
			"&toDateInclusive="+toDateInclusive,
		payload.bytesBuffer(),
	)
	if err != nil {
		return nil, err
	}

	byteArr, err := alvtimeClient.do(req)
	if err != nil {
		return nil, err
	}

	var timeEntries []TimeEntrie
	json.Unmarshal(byteArr, &timeEntries)

	return timeEntries, nil
}
