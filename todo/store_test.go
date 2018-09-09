package todo

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestReadTodosFromJSONNotToFindFile(t *testing.T) {
	file := ""
	_, actual := readTodosFromJSON(file)
	if actual == nil {
		expected := "open : no such file or directory"
		msg := "Didn't return file open error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestReadTodosFromJSONIncorrectJSONData(t *testing.T) {
	file, teardown := setupEmptyTodosFile(t)
	defer teardown()

	_, actual := readTodosFromJSON(file)
	if actual == nil {
		expected := "invalid character '[' looking for beginning of object key string"
		msg := "Didn't return parsing json data error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestReadTodosFromJSONCorrectJSONData(t *testing.T) {
	file, teardown := setupTodosFile(t)
	defer teardown()

	todos := Todos{
		{
			Text:      "message",
			Timestamp: time.Time{},
			IsStarred: true,
		},
	}
	readTodos, err := readTodosFromJSON(file)
	if err != nil {
		msg := "Didn't parse json data."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, nil, err)
	}
	for i, actual := range readTodos {
		expected := todos[i]
		if actual != expected {
			msg := "Didn't parse json data correctory."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
		}
	}
}

func TestWriteTodosToJSON(t *testing.T) {
	file, teardown := setupEmptyTodosFile(t)
	defer teardown()

	todos := Todos{
		{
			Text:      "message",
			Timestamp: time.Time{},
			IsStarred: true,
		},
	}

	if err := writeTodosToJSON(file, todos); err != nil {
		msg := "Didn't write to file. Caught a error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, nil, err)
	}

	expected := `[{"text":"message","timestamp":"0001-01-01T00:00:00Z","is_starred":true,"is_complete":false}]`
	actual, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatalf("Error occurred before checking result.\n%s\n", err)
	}
	if string(actual) != expected {
		msg := "Didn't write todo list to file correctory."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, nil, err)
	}
}

func TestLoadTodos(t *testing.T) {
	file, teardown := setupTodosFile(t)
	defer teardown()

	todos, err := LoadTodos(file)
	if err != nil {
		msg := "Didn't load todo list. Caught a error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, nil, err)
	}
	expected := Todo{
		Text:      "message",
		Timestamp: time.Time{},
		IsStarred: true,
	}
	actual := todos[0]
	if actual != expected {
		msg := "Didn't load todo list correctory."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestSaveTodos(t *testing.T) {
	file, teardown := setupEmptyTodosFile(t)
	defer teardown()

	todos := Todos{
		{
			Text:      "message",
			Timestamp: time.Time{},
			IsStarred: true,
		},
	}

	if err := SaveTodos(file, todos); err != nil {
		msg := "Didn't save todo list."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, nil, err)
	}
}
