package main

import (
	"fmt"
	"os"
	"task-tracker/cmd"
	"task-tracker/handler"
)

func main() {
	h := handler.Handler{Commands: make(map[string]cmd.Command)}
	h.AddCommand(cmd.AddTask)
	h.AddCommand(cmd.ListTasks)
	h.AddCommand(cmd.MarkTask)
	h.AddCommand(cmd.UpdateTask)
	h.AddCommand(cmd.DeleteTask)

	args := os.Args[1:]
	if len(args) == 0 {
		h.PrintHelp()
		return
	}

	command, ok := h.Commands[args[0]]
	if !ok {
		fmt.Printf("Unknown command: %s\nRun task-tracker with no arguments to see the list of commands", args[0])
		return
	}

	result, err := command.Function(args[1:])
	if err != nil {
		fmt.Println(result, err)
	} else {
		fmt.Println(result)
	}
}
