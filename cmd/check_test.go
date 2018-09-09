package cmd

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	td "github.com/x-color/tdl/todo"
)

func TestCmdCheckToCheckOutputMessage(t *testing.T) {
	_, teardown := setupTodosFile(t)
	defer teardown()

	testCases := []struct {
		num    string
		output string
	}{
		{"1", "Check: 'message1'"},
		{"4", "Uncheck: 'message4'"},
	}

	for _, tc := range testCases {
		expected := tc.output
		output := executeCmd([]string{"check", tc.num})
		actual := strings.Split(output, "\n")[0]
		if actual != expected {
			msg := "Output message was not a text of checked or unchecked todo."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
		}
	}
}

func TestCmdCheckToCheckOutputMessageMultipleArguments(t *testing.T) {
	_, teardown := setupTodosFile(t)
	defer teardown()

	testCases := []struct {
		num    string
		output string
	}{
		{"1", "Check: 'message1'"},
		{"4", "Uncheck: 'message4'"},
	}

	output := executeCmd([]string{"check", testCases[0].num, testCases[1].num})
	for i, actual := range strings.Split(output, "\n")[:2] {
		expected := testCases[i].output
		if actual != expected {
			msg := "Output message was not a text of checked or unchecked todo."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
		}
	}
}

func TestCmdCheckToCheckTodos(t *testing.T) {
	_, teardown := setupTodosFile(t)
	defer teardown()

	testCases := []struct {
		num      string
		complete bool
	}{
		{"1", true},  // Check
		{"4", false}, // Uncheck
	}

	for _, tc := range testCases {
		executeCmd([]string{"check", tc.num})
	}

	b, err := ioutil.ReadFile(todosFile)
	if err != nil {
		t.Fatalf("Error occurred before checking result.\n%s\n", err)
	}
	todos := td.Todos{}
	err = json.Unmarshal(b, &todos)
	if err != nil {
		t.Fatalf("Error occurred before checking result.\n%s\n", err)
	}

	for _, tc := range testCases {
		expected := tc.complete
		i, _ := strconv.Atoi(tc.num)
		actual := todos[i-1].IsComplete
		if actual != expected {
			msg := "Didn't change a complete flag of todo in file."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
		}
	}
}

func TestCmdCheckToCheckOutOfRangeError(t *testing.T) {
	testOutOfRangeError(t, []string{"check", "0"})
	testOutOfRangeError(t, []string{"check", "10"})
}

func TestCmdCheckToCheckNoArgumentsError(t *testing.T) {
	testNoArgumentsError(t, []string{"check"})
}

func TestCmdCheckToCheckNoNumberError(t *testing.T) {
	testNoNumberError(t, []string{"check", "number"})
}

func TestCmdCheckToCheckNoFileError(t *testing.T) {
	testNoFileError(t, []string{"check", "1"})
}

func TestCmdCheckToCheckFileTextError(t *testing.T) {
	testFileTextError(t, []string{"check", "1"})
}
