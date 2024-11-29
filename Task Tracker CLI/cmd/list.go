package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"task-tracker/task"
)

var ListTasks Command = Command{
	Name:        "list",
	Description: "Lists all created tasks. list [status] to list all tasks with given status",
	Function: func(status []string) (string, error) {
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
		result := "ID\tDescription\tStatus\tCreated at\tUpdated at\n"
		if len(status) == 0 {
			for _, t := range tasks {
				result += fmt.Sprintln(t)
			}
		} else if len(status) == 1 {
			filter := task.ParseStatus(status[0])
			if filter == -1 {
				return fmt.Sprintf("Invalid input. Unknown task status \"%v\"", status[0]), nil
			} else {
				for _, t := range tasks {
					if t.Status == filter {
						result += fmt.Sprintln(t)
					}
				}
			}
		} else {
			return "Invalid input. Correct command usage: task-tracker list [status]", nil
		}
		return result, nil
	},
}
