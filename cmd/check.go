package cmd

import (
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	td "github.com/x-color/tdl/todo"
)

func runCmdCheck(cmd *cobra.Command, args []string) error {
	todos, err := td.LoadTodos(todosFile)
	if err != nil {
		return err
	}
	sort.Sort(sort.StringSlice(args))
	for _, sNum := range args {
		num, _ := strconv.Atoi(sNum) // Aready validation all arguments. So, don't catch error.
		todo, err := todos.Check(num)
		if err != nil {
			return err
		}
		if todo.IsComplete {
			cmd.Printf("Check: '%s'\n", todo.Text)
		} else {
			cmd.Printf("Uncheck: '%s'\n", todo.Text)
		}
	}
	return td.SaveTodos(todosFile, todos)
}

func newCmdCheck() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check or uncheck one or more todos",
		Args:  validateNumberArgs,
		RunE:  runCmdCheck,
	}
	return cmd
}
