package handler

import (
	"fmt"
	"task-tracker/cmd"
)

type Handler struct {
	Commands map[string]cmd.Command
}

func (h Handler) PrintHelp() {
	for name, command := range h.Commands {
		fmt.Printf("%s\t%s\n", name, command.Description)
	}
}

func (h *Handler) AddCommand(command cmd.Command) {
	h.Commands[command.Name] = command
}
