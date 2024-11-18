package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
	"task-tracker/task"
	"time"
)

var AddTask Command = Command{
	Name:        "add",
	Description: "Adds new task with the given description",
	Function: func(description []string) (string, error) {
		var tasks []task.Task

		file, err := os.Open("tasks.json")
		if err != nil {
			file, err = os.Create("tasks.json")
			if err != nil {
				return "Error when creating file:\n", err
			}
			_, err = file.WriteString("[]")
			if err != nil {
				return "Error when creating file:\n", err
			}
		}
		defer file.Close()

		byteValue, _ := io.ReadAll(file)
		if len(byteValue) > 2 {
			err = json.Unmarshal(byteValue, &tasks)
			if err != nil {
				return "Error when unmarshalling:\n", err
			}
		}

		ids := make([]int, len(tasks))
		for _, t := range tasks {
			ids = append(ids, t.ID)
		}

		for i := 1; ; i++ {
			if !slices.Contains(ids, i) {
				newTask := task.Make(i, description[0], task.Todo, time.Now(), time.Now())
				tasks = append(tasks, newTask)

				file, err := json.MarshalIndent(tasks, "", "  ")
				if err != nil {
					return "Error when marshalling:\n", err
				}

				err = os.WriteFile("tasks.json", file, 0644)
				if err != nil {
					return "Error when writing into file:\n", err
				}
				return fmt.Sprintf("Task added successfully (ID: %v)", i), nil
			}
		}
	},
}
