package main

import (
	"encoding/json"
	"net/url"
)

// TimeEntrie struct
type TimeEntrie struct {
	Date   string  `json:"date"`
	Value  float32 `json:"value"`
	TaskID int     `json:"taskId"`
}

// GetTimeEntries fetches all the time entries available to the authenticated user
// between the dates denoted by fromDateInclusive to toDateInclusive formated as "YYYY-MM-DD"
func (c *AlvtimeClient) GetTimeEntries(
	fromDateInclusive, toDateInclusive string,
) (timeEntries []TimeEntrie, err error) {
	baseURL, err := url.Parse(c.domain)
	if err != nil {
		return nil, err
	}

	baseURL.Path += "/api/user/TimeEntries"

	params := url.Values{}
	params.Add("fromDateInclusive", fromDateInclusive)
	params.Add("toDateInclusive", toDateInclusive)
	baseURL.RawQuery = params.Encode()

	req, err := c.newRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	byteArr, err := c.do(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteArr, &timeEntries)
	if err != nil {
		return nil, err
	}

	return timeEntries, nil
}

// EditTimeEntries is used to edit favorite tasks by setting the Favorite
// attribute on and passing it in a slice. The edited tasks are returned as stored to confirm.
func (c *AlvtimeClient) EditTimeEntries(timeEntriesToEdit []TimeEntrie) (editedTimeEntreis []TimeEntrie, err error) {
	baseURL, err := url.Parse(c.domain)
	if err != nil {
		return nil, err
	}
	baseURL.Path += "/api/user/TimeEntries"

	req, err := c.newRequest("POST", baseURL, timeEntriesToEdit)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	byteArr, err := c.do(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteArr, &editedTimeEntreis)
	if err != nil {
		return nil, err
	}

	return editedTimeEntreis, nil
}
