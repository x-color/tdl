package cmd

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/spf13/cobra"
	td "github.com/x-color/tdl/todo"
)

var showAllTemplate = `All Todos
{{- range $i, $todo := .todos}}
  {{printf "%02d." (add $i 1)}} [{{if $todo.IsComplete }}{{"\x1b[32m✔\x1b[0m"}}{{else}} {{end}}] {{if $todo.IsComplete }}{{"\x1b[9m"}}{{textOverflow $todo.Text}}{{"\x1b[0m"}}{{else}}{{textOverflow $todo.Text}}{{end}} {{if $todo.IsStarred }}{{"\x1b[33m★\x1b[0m"}}{{end}}
{{- end}}

-- {{.done}} done  {{.active}} active --`

var showActiveTemplate = `Active Todos
{{- range $i, $todo := .todos}}
{{- if not $todo.IsComplete }}
  {{printf "%02d." (add $i 1)}} [{{if $todo.IsComplete }}{{"\x1b[32m✔\x1b[0m"}}{{else}} {{end}}] {{textOverflow $todo.Text}} {{if $todo.IsStarred }}{{"\x1b[33m★\x1b[0m"}}{{end}}
{{- end}}
{{- end}}

-- {{.active}} active --`

var showDoneTemplate = `Done Todos
{{- range $i, $todo := .todos}}
{{- if $todo.IsComplete }}
  {{printf "%02d." (add $i 1)}} [{{if $todo.IsComplete }}{{"\x1b[32m✔\x1b[0m"}}{{else}} {{end}}] {{textOverflow $todo.Text}} {{if $todo.IsStarred }}{{"\x1b[33m★\x1b[0m"}}{{end}}
{{- end}}
{{- end}}

-- {{.done}} done --`

type optionCmdShow struct {
	active bool
	done   bool
	length int
}

var optCmdShow = optionCmdShow{}

func add(a, b int) int {
	return a + b
}

func textOverflow(rowText string) string {
	if optCmdShow.length > 0 {
		if len(rowText) > optCmdShow.length {
			return rowText[:optCmdShow.length] + "..."
		}
		return rowText + strings.Repeat(" ", optCmdShow.length-len(rowText)) + "   "
	}
	return rowText
}

func mappingData(todos td.Todos) map[string]interface{} {
	active := 0
	done := 0
	for _, todo := range todos {
		if todo.IsComplete {
			done++
		} else {
			active++
		}
	}
	return map[string]interface{}{
		"todos":  todos,
		"active": active,
		"done":   done,
	}
}

func runCmdShow(cmd *cobra.Command, args []string) error {
	todos, err := td.LoadTodos(todosFile)
	if err != nil {
		return err
	}

	funcMap := template.FuncMap{
		"add":          add,
		"textOverflow": textOverflow,
	}
	showTemp := ""
	switch {
	case optCmdShow.active && !optCmdShow.done:
		showTemp = showActiveTemplate
	case !optCmdShow.active && optCmdShow.done:
		showTemp = showDoneTemplate
	default:
		showTemp = showAllTemplate
	}
	t, err := template.New("showTodos").Funcs(funcMap).Parse(showTemp)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, mappingData(todos))
	if err != nil {
		return err
	}
	cmd.Println(buf)
	return nil
}

func newCmdShow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show todos",
		Args:  cobra.NoArgs,
		RunE:  runCmdShow,
	}

	cmd.Flags().BoolVarP(&optCmdShow.active, "active", "a", false, "show only incomplete todos")
	cmd.Flags().BoolVarP(&optCmdShow.done, "done", "d", false, "show only complete todos")
	cmd.Flags().IntVarP(&optCmdShow.length, "length", "l", -50, "show only `NUM` charactors of todo")

	return cmd
}
