package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func setupEmptyTodosFile(t *testing.T) func() {
	t.Helper()
	tmpFile, err := ioutil.TempFile("", "tmptest")
	if err != nil {
		t.Fatalf("Error occurred in setup process before testing.\n%s\n", err)
	}
	todosFile = tmpFile.Name()
	return func() { os.Remove(tmpFile.Name()) }
}

func setupTodosFile(t *testing.T) (int, func()) {
	t.Helper()
	teardown := setupEmptyTodosFile(t)
	jsonData := []byte(`[
        {"text":"message1","timestamp":"0001-01-01T00:00:00Z","is_starred":false,"is_complete":false},
        {"text":"message2","timestamp":"0001-01-01T00:00:00Z","is_starred":true,"is_complete":false},
        {"text":"message3","timestamp":"0001-01-01T00:00:00Z","is_starred":false,"is_complete":true},
        {"text":"message4","timestamp":"0001-01-01T00:00:00Z","is_starred":true,"is_complete":true}
        ]`)
	if err := ioutil.WriteFile(todosFile, jsonData, 0644); err != nil {
		t.Fatalf("Error occurred in setup process before testing.\n%s\n", err)
	}
	return 4, teardown // 4 is numbers of json data
}

func executeCmd(args []string) string {
	buf := new(bytes.Buffer)
	cmdRoot := newCmdRoot()
	cmdRoot.SetOutput(buf)
	cmdRoot.SetArgs(args)
	cmdRoot.Execute()
	return buf.String()
}

func testOutOfRangeError(t *testing.T, cmd []string) {
	t.Helper()
	_, teardown := setupTodosFile(t)
	defer teardown()

	output := executeCmd(cmd)

	expected := "Error: out of range of todo list"
	actual := strings.Split(output, "\n")[0]
	if actual != expected {
		msg := "Output message was not a message of out of range error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func testNoArgumentsError(t *testing.T, cmd []string) {
	t.Helper()
	_, teardown := setupTodosFile(t)
	defer teardown()

	output := executeCmd(cmd)

	expected := "Error: requires at least 1 arg(s), only received 0"
	actual := strings.Split(output, "\n")[0]
	if actual != expected {
		msg := "Output message was not a message of no arguments error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func testNoNumberError(t *testing.T, cmd []string) {
	t.Helper()
	_, teardown := setupTodosFile(t)
	defer teardown()

	output := executeCmd(cmd)

	expected := "Error: requires one or more numbers, argument 1 is not number"
	actual := strings.Split(output, "\n")[0]
	if actual != expected {
		msg := "Output message was not a message of received no number error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func testNoFileError(t *testing.T, cmd []string) {
	t.Helper()
	todosFile = ""

	output := executeCmd(cmd)

	expected := "Error: open " + todosFile + ": no such file or directory"
	actual := strings.Split(output, "\n")[0]
	if actual != expected {
		msg := "Output message was not a message of no todos file error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func testFileTextError(t *testing.T, cmd []string) {
	t.Helper()
	teardown := setupEmptyTodosFile(t)
	defer teardown()

	output := executeCmd(cmd)

	expected := "Error: unexpected end of JSON input"
	actual := strings.Split(output, "\n")[0]
	if actual != expected {
		msg := "Output message was not a message of parsing json error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}
