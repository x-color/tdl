package cmd

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"

	td "github.com/x-color/tdl/todo"
)

func TestCmdDeleteToCheckOutputMessage(t *testing.T) {
	_, teardown := setupTodosFile(t)
	defer teardown()

	output := executeCmd([]string{"delete", "1"})

	expected := "Delete: 'message1'"
	actual := strings.Split(output, "\n")[0]
	if actual != expected {
		msg := "Output message was not a text of deleted todo."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestCmdDeleteToCheckOutputMessageMultipleArguments(t *testing.T) {
	_, teardown := setupTodosFile(t)
	defer teardown()

	testCases := []struct {
		num    string
		output string
	}{
		{"1", "Delete: 'message1'"},
		{"4", "Delete: 'message4'"},
	}

	output := executeCmd([]string{"delete", testCases[0].num, testCases[1].num})
	for i, actual := range strings.Split(output, "\n")[:2] {
		expected := testCases[i].output
		if actual != expected {
			msg := "Output message was not a text of deleted todo."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
		}
	}
}

func TestCmdDeleteToCheckTodos(t *testing.T) {
	n, teardown := setupTodosFile(t)
	defer teardown()

	executeCmd([]string{"delete", "1"})

	b, err := ioutil.ReadFile(todosFile)
	if err != nil {
		t.Fatalf("Error occurred before checking result.\n%s\n", err)
	}
	todos := td.Todos{}
	err = json.Unmarshal(b, &todos)
	if err != nil {
		t.Fatalf("Error occurred before checking result.\n%s\n", err)
	}
	expected := n - 1
	actual := len(todos)
	if actual != expected {
		msg := "Didn't delete a todo in file."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestCmdDeleteToCheckOutOfRangeError(t *testing.T) {
	testOutOfRangeError(t, []string{"delete", "0"})
	testOutOfRangeError(t, []string{"delete", "10"})
}

func TestCmdDeleteToCheckNoArgumentsError(t *testing.T) {
	testNoArgumentsError(t, []string{"delete"})
}

func TestCmdDeleteToCheckNoNumberError(t *testing.T) {
	testNoNumberError(t, []string{"delete", "number"})
}

func TestCmdDeleteToCheckNoFileError(t *testing.T) {
	testNoFileError(t, []string{"delete", "1"})
}

func TestCmdDeleteToCheckFileTextError(t *testing.T) {
	testFileTextError(t, []string{"delete", "1"})
}
