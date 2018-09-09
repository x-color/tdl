package cmd

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"

	td "github.com/x-color/tdl/todo"
)

func TestCmdAddToCheckOutputMessage(t *testing.T) {
	_, teardown := setupTodosFile(t)
	defer teardown()

	output := executeCmd([]string{"add", "todo title"})

	expected := "Add: 'todo title'"
	actual := strings.Split(output, "\n")[0]
	if actual != expected {
		msg := "Output message was not a text of added new todo."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestCmdAddToCheckOutputMessageMultipleArguments(t *testing.T) {
	_, teardown := setupTodosFile(t)
	defer teardown()

	testCases := []struct {
		text   string
		output string
	}{
		{"message5", "Add: 'message5'"},
		{"message6", "Add: 'message6'"},
	}

	output := executeCmd([]string{"add", testCases[0].text, testCases[1].text})
	for i, actual := range strings.Split(output, "\n")[:2] {
		expected := testCases[i].output
		if actual != expected {
			msg := "Output message was not a text of added todo."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
		}
	}
}

func TestCmdAddToCheckTodos(t *testing.T) {
	n, teardown := setupTodosFile(t)
	defer teardown()

	executeCmd([]string{"add", "todo title"})

	b, err := ioutil.ReadFile(todosFile)
	if err != nil {
		t.Fatalf("Error occurred before checking result.\n%s\n", err)
	}
	todos := td.Todos{}
	err = json.Unmarshal(b, &todos)
	if err != nil {
		t.Fatalf("Error occurred before checking result.\n%s\n", err)
	}
	expected := n + 1
	actual := len(todos)
	if actual != expected {
		msg := "Didn't save a new todo to file."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestCmdAddToCheckNoArgumentsError(t *testing.T) {
	testNoArgumentsError(t, []string{"add"})
}

func TestCmdAddToCheckNoFileError(t *testing.T) {
	testNoFileError(t, []string{"add", "todo text"})
}

func TestCmdAddToCheckFileTextError(t *testing.T) {
	testFileTextError(t, []string{"add", "todo text"})
}
