package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"task-tracker/task"
)

var DeleteTask Command = Command{
	Name:        "delete",
	Description: "Deletes a task",
	Function: func(args []string) (string, error) {
		if len(args) != 1 {
			return "Invalid input. Correct command usage: delete [description]", nil
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Sprintf("Invalid input. Can't convert \"%v\" to an integer", args[0]), nil
		}

		var tasks []task.Task

		file, err := os.Open("tasks.json")
		if err != nil {
			return "No tasks found. Use \"add [description]\" to add tasks", nil
		}
		defer file.Close()

		byteValue, _ := io.ReadAll(file)
		if len(byteValue) > 2 {
			err = json.Unmarshal(byteValue, &tasks)
			if err != nil {
				return "Error when unmarshalling:\n", err
			}
		}
		for _, t := range tasks {
			if t.ID == id {
				tasks = append(tasks[0:id-1], tasks[id:]...)
				file, err := json.MarshalIndent(tasks, "", "  ")
				if err != nil {
					return "Error when marshalling:\n", err
				}

				err = os.WriteFile("tasks.json", file, 0644)
				if err != nil {
					return "Error when writing into file:\n", err
				}
				return fmt.Sprintf("Task deleted successfully (ID: %v)", id), nil
			}
		}
		return fmt.Sprintf("Couldn't find task with ID %v", id), nil
	},
}
