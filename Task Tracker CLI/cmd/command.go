package cmd

type Command struct {
	Name        string
	Description string
	Function    func(args []string) (string, error)
}
