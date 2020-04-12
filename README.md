# alvtime-go-client

Alvtime go client

## Install

```console
go get github.com/Alv-no/alvtime-go-client
```

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"

	alvtimeClient "github.com/Alv-no/alvtime-go-client"
)

func main() {
	c, err := alvtimeClient.New("http://localhost:8080")

	tasks, err := c.GetTasks()
	if err != nil {
		fmt.Println(err)
	}

	tasks[0].Favorite = true
	tasksToEdit := []alvtimeClient.Task{tasks[0]}
	editedTasks, err := c.EditFavoriteTasks(tasksToEdit)
	if err != nil {
		fmt.Println(err)
	}
	prettyPrint("editedTasks", editedTasks)

	timeEntries, err := c.GetTimeEntries("2019-01-01", "2020-01-01")
	if err != nil {
		fmt.Println(err)
	}
	prettyPrint("timeEntries", timeEntries)

	timeEntriesToEdit := []alvtimeClient.TimeEntrie{
		{Date: "2019-01-27", Value: 7, TaskID: 2},
		{Date: "2019-01-26", Value: 7.5, TaskID: 2},
	}
	editedTimeEntries, err := c.EditTimeEntries(timeEntriesToEdit)
	if err != nil {
		fmt.Println(err)
	}

	prettyPrint("editedTimeEntries", editedTimeEntries)
}

func prettyPrint(key string, i interface{}) {
	s, _ := json.MarshalIndent(i, "", "  ")
	fmt.Printf("\n%s = %v", key, string(s))
}
```
## Setting up the development container

Follow these steps to open this project in a development container:

1. If this is your first time using a development container, please follow the [getting started steps](https://aka.ms/vscode-remote/containers/getting-started).

2. In Visual Studio Code, press <kbd>F1</kbd> and select the **Remote-Containers: Open Folder in Container...** command. Select the cloned copy of this folder, wait for the container to start, and try things out!

## Swagger documentation

Starting the development environment exposes the [alvtime-web-api](https://github.com/Alv-no/alvtime-web-api) running in a container on `localhost:8080/api/`. This also exposes the swagger documentation on `localhost:8080/swagger`
