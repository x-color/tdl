package cmd

import (
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	td "github.com/x-color/tdl/todo"
)

func runCmdStar(cmd *cobra.Command, args []string) error {
	todos, err := td.LoadTodos(todosFile)
	if err != nil {
		return err
	}
	sort.Sort(sort.StringSlice(args))
	for _, sNum := range args {
		num, _ := strconv.Atoi(sNum) // Aready validation all arguments. So, don't catch error.
		todo, err := todos.Star(num)
		if err != nil {
			return err
		}
		if todo.IsStarred {
			cmd.Printf("Star: '%s'\n", todo.Text)
		} else {
			cmd.Printf("Unstar: '%s'\n", todo.Text)
		}
	}
	return td.SaveTodos(todosFile, todos)
}

func newCmdStar() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "star",
		Short: "Star or unstar one or more todos",
		Args:  validateNumberArgs,
		RunE:  runCmdStar,
	}
	return cmd
}
