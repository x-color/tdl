package cmd

import (
	"github.com/spf13/cobra"
	td "github.com/x-color/tdl/todo"
)

func runCmdAdd(cmd *cobra.Command, args []string) error {
	todos, err := td.LoadTodos(todosFile)
	if err != nil {
		return err
	}
	for _, text := range args {
		todo := td.NewTodo(text)
		todos.Add(todo)
		cmd.Printf("Add: '%s'\n", todo.Text)
	}
	return td.SaveTodos(todosFile, todos)
}

func newCmdAdd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add one or more todos",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runCmdAdd,
	}
	return cmd
}
