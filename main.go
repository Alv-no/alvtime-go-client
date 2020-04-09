package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	alvtimeClient := AlvtimeClient{
		domain:     "http://alvtime-web-api-no-auth",
		httpClient: &http.Client{},
	}

	tasks, err := alvtimeClient.GetTasks()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(tasks)
}
