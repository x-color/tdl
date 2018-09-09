package cmd

import (
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	td "github.com/x-color/tdl/todo"
)

func runCmdDelete(cmd *cobra.Command, args []string) error {
	todos, err := td.LoadTodos(todosFile)
	if err != nil {
		return err
	}
	sort.Sort(sort.StringSlice(args))
	for i, sNum := range args {
		num, _ := strconv.Atoi(sNum)       // Aready validation all arguments. So, don't catch error.
		todo, err := todos.Delete(num - i) // If give multiple arguments to this command, todos shifts by one.
		if err != nil {
			return err
		}
		cmd.Printf("Delete: '%s'\n", todo.Text)
	}
	return td.SaveTodos(todosFile, todos)
}

func newCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete one or more todos",
		Args:  validateNumberArgs,
		RunE:  runCmdDelete,
	}
	return cmd
}
