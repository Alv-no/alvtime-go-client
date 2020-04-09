package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	alvtimeClient := AlvtimeClient{
		domain:     "http://dotnet-backend",
		httpClient: &http.Client{},
	}

	tasks, err := alvtimeClient.GetTasks()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(tasks)
}
