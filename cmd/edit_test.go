package cmd

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"

	td "github.com/x-color/tdl/todo"
)

func TestCmdEditToCheckOutputMessage(t *testing.T) {
	_, teardown := setupTodosFile(t)
	defer teardown()

	output := executeCmd([]string{"edit", "1", "edited text"})

	expected := "Edit: 'message1' => 'edited text'"
	actual := strings.Split(output, "\n")[0]
	if actual != expected {
		msg := "Output message was not a correctory text."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestCmdEditToCheckTodos(t *testing.T) {
	n, teardown := setupTodosFile(t)
	defer teardown()

	executeCmd([]string{"edit", "1", "edited text"})

	b, err := ioutil.ReadFile(todosFile)
	if err != nil {
		t.Fatalf("Error occurred before checking result.\n%s\n", err)
	}
	todos := td.Todos{}
	err = json.Unmarshal(b, &todos)
	if err != nil {
		t.Fatalf("Error occurred before checking result.\n%s\n", err)
	}
	expected := n
	actual := len(todos)
	if actual != expected {
		msg := "Didn't edit a todo in file."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestCmdEditToCheckOutOfRangeError(t *testing.T) {
	testOutOfRangeError(t, []string{"edit", "10", "edited text"})
}

func TestCmdEditToCheckNoArgumentsError(t *testing.T) {
	_, teardown := setupTodosFile(t)
	defer teardown()

	output := executeCmd([]string{"edit"})
	actual := strings.Split(output, "\n")[0]
	expected := "Error: accepts 2 arg(s), received 0"
	if actual != expected {
		msg := "Output message was not a message of no arguments error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestCmdEditToCheckManyArguments(t *testing.T) {
	_, teardown := setupTodosFile(t)
	defer teardown()

	output := executeCmd([]string{"edit", "1", "edited text", "2", "edited text 2"})
	actual := strings.Split(output, "\n")[0]
	expected := "Error: accepts 2 arg(s), received 4"
	if actual != expected {
		msg := "Output message was not a message of no arguments error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestCmdEditToCheckNoNumberError(t *testing.T) {
	_, teardown := setupTodosFile(t)
	defer teardown()

	output := executeCmd([]string{"edit", "number", "edited text"})

	expected := "Error: requires number and text, first argument is not number"
	actual := strings.Split(output, "\n")[0]
	if actual != expected {
		msg := "Output message was not a message of received no number error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestCmdEditToCheckNoFileError(t *testing.T) {
	testNoFileError(t, []string{"edit", "1", "edited text"})
}

func TestCmdEditToCheckFileTextError(t *testing.T) {
	testFileTextError(t, []string{"edit", "1", "edited text"})
}
