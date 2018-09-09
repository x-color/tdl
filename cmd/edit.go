package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	td "github.com/x-color/tdl/todo"
)

func runCmdEdit(cmd *cobra.Command, args []string) error {
	todos, err := td.LoadTodos(todosFile)
	if err != nil {
		return err
	}
	// Aready validation first argument.
	num, _ := strconv.Atoi(args[0])
	oldTodo, err := todos.Get(num)
	if err != nil {
		return err
	}
	// Already check error the one above code.
	todo, _ := todos.Edit(num, args[1])
	cmd.Printf("Edit: '%s' => '%s'\n", oldTodo.Text, todo.Text)
	return td.SaveTodos(todosFile, todos)
}

func validateCmdEditArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.ExactArgs(2)(cmd, args); err != nil {
		return err
	}
	if _, err := strconv.Atoi(args[0]); err != nil {
		return fmt.Errorf("requires number and text, first argument is not number")
	}
	return nil
}

func newCmdEdit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit one todo",
		Args:  validateCmdEditArgs,
		RunE:  runCmdEdit,
	}

	return cmd
}
