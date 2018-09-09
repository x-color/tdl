package cmd

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	td "github.com/x-color/tdl/todo"
)

func TestCmdStarToCheckOutputMessage(t *testing.T) {
	_, teardown := setupTodosFile(t)
	defer teardown()

	testCases := []struct {
		num    string
		output string
	}{
		{"1", "Star: 'message1'"},
		{"4", "Unstar: 'message4'"},
	}

	for _, tc := range testCases {
		expected := tc.output
		output := executeCmd([]string{"star", tc.num})
		actual := strings.Split(output, "\n")[0]
		if actual != expected {
			msg := "Output message was not a text of starred or unstarred todo."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
		}
	}
}

func TestCmdStarToCheckOutputMessageMultipleArguments(t *testing.T) {
	_, teardown := setupTodosFile(t)
	defer teardown()

	testCases := []struct {
		num    string
		output string
	}{
		{"1", "Star: 'message1'"},
		{"4", "Unstar: 'message4'"},
	}

	output := executeCmd([]string{"star", testCases[0].num, testCases[1].num})
	for i, actual := range strings.Split(output, "\n")[:2] {
		expected := testCases[i].output
		if actual != expected {
			msg := "Output message was not a text of starred or unstarred todo."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
		}
	}
}

func TestCmdStarToCheckTodos(t *testing.T) {
	_, teardown := setupTodosFile(t)
	defer teardown()

	testCases := []struct {
		num     string
		starred bool
	}{
		{"1", true},  // Star
		{"4", false}, // Unstar
	}

	for _, tc := range testCases {
		executeCmd([]string{"star", tc.num})
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
		expected := tc.starred
		i, _ := strconv.Atoi(tc.num)
		actual := todos[i-1].IsStarred
		if actual != expected {
			msg := "Didn't change a starred flag of todo in file."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
		}
	}
}

func TestCmdStarToCheckOutOfRangeError(t *testing.T) {
	testOutOfRangeError(t, []string{"star", "0"})
	testOutOfRangeError(t, []string{"star", "10"})
}

func TestCmdStarToCheckNoArgumentsError(t *testing.T) {
	testNoArgumentsError(t, []string{"star"})
}

func TestCmdStarToCheckNoNumberError(t *testing.T) {
	testNoNumberError(t, []string{"star", "number"})
}

func TestCmdStarToCheckNoFileError(t *testing.T) {
	testNoFileError(t, []string{"star", "1"})
}

func TestCmdStarToCheckFileTextError(t *testing.T) {
	testFileTextError(t, []string{"star", "1"})
}
