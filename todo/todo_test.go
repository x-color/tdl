package todo

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func makeTodo(text string) Todo {
	return Todo{
		Text:      text,
		Timestamp: time.Now(),
	}
}

// ***************************************
// Test Todo receivers
// ***************************************

func TestTodoCheck(t *testing.T) {
	todo := makeTodo("message")

	todo.Check()

	expected := true
	actual := todo.IsComplete
	if actual != expected {
		msg := "Didn't make a todo complete."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestTodoUncheck(t *testing.T) {
	todo := makeTodo("message")
	todo.IsComplete = true

	todo.Uncheck()

	expected := false
	actual := todo.IsComplete
	if actual != expected {
		msg := "Din't make a todo incomplete."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestTodoStar(t *testing.T) {
	todo := makeTodo("message")

	todo.Star()

	expected := true
	actual := todo.IsStarred
	if actual != expected {
		msg := "Din't star a todo."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestTodoUnstar(t *testing.T) {
	todo := makeTodo("message")
	todo.IsStarred = true

	todo.Unstar()

	expected := false
	actual := todo.IsStarred
	if actual != expected {
		msg := "Didn't unstar a todo."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestTodoEdit(t *testing.T) {
	todo := makeTodo("message")
	text := "test message"

	todo.Edit(text)

	expected := makeTodo(text)
	actual := todo
	if actual.Text != expected.Text {
		msg := "Didn't edit a text of todo."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
	if !(actual.Timestamp.After(expected.Timestamp.Add(-time.Second)) && actual.Timestamp.Before(expected.Timestamp.Add(time.Second))) {
		msg := "Didn't update a timestamp of todo."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

// ***************************************
// Test functions
// ***************************************

func TestNewTodo(t *testing.T) {
	todo := makeTodo("message")
	expected := todo
	actual := NewTodo(todo.Text)
	if actual.Text != expected.Text {
		msg := "Din't make correctory a todo variable. A text of todo is not correctory."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
	if !(actual.Timestamp.After(expected.Timestamp.Add(-time.Second)) && actual.Timestamp.Before(expected.Timestamp.Add(time.Second))) {
		msg := "Din't make correctory a todo variable. A timestamp of todo is not correctory."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

// ***************************************
// Test Todos receivers
// ***************************************

// ***************************************
// Test Todos.Get
// ***************************************

func TestTodosGetToCheckGot(t *testing.T) {
	todos := Todos{
		makeTodo("hahaha"),
	}
	todo, err := todos.Get(1)
	if err != nil {
		msg := "Caught the unexpected out of range error. Didn't get a todo from todo list."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, nil, err)
	}

	actual := todos[0]
	expected := todo
	if actual != expected {
		msg := "Didn't get a todo from todo list."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestTodosGetToCheckOutOfRange(t *testing.T) {
	todos := Todos{
		makeTodo("message"),
	}

	expected := fmt.Sprint(errors.New("out of range of todo list"))
	for _, n := range []int{0, 2} {
		_, err := todos.Get(n)
		actual := fmt.Sprint(err)
		if actual != expected {
			msg := "Didn't catch the out of range error."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
		}
	}
}

// ***************************************
// Test Todos.Add
// ***************************************

func TestTodosAddToCheckAdded(t *testing.T) {
	todos := Todos{
		makeTodo("hahaha"),
	}
	todo := makeTodo("message")
	todos.Add(todo)
	actual := todos[0]
	expected := todo
	if actual != expected {
		msg := "Didn't add a new todo to head of todo list."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestTodosAddToCheckReturn(t *testing.T) {
	todos := Todos{
		makeTodo("hahaha"),
	}
	todo := makeTodo("message")
	actual := todos.Add(todo)
	expected := todo
	if actual != expected {
		msg := "Didn't return an added todo to todo list."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

// ***************************************
// Test Todos.Delete
// ***************************************

func TestTodosDeleteToCheckDeleted(t *testing.T) {
	newTodo := makeTodo("test delete todo")
	todos := Todos{
		newTodo,
		makeTodo("message 2"),
	}
	_, err := todos.Delete(1)
	if err != nil {
		msg := "Caught the unexpected out of range error. Didn't delete a todo from todo list."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, nil, err)
	}
	for _, todo := range todos {
		if newTodo == todo {
			msg := "Didn't delete a todo from todo list."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, "deleted", "exist")
		}
	}
}

func TestTodosDeleteToCheckOutOfRange(t *testing.T) {
	newTodo := makeTodo("test delete todo")
	todos := Todos{
		newTodo,
		makeTodo("message 2"),
	}

	expected := fmt.Sprint(errors.New("out of range of todo list"))
	for _, n := range []int{0, 3} {
		_, err := todos.Delete(n)
		actual := fmt.Sprint(err)
		if actual != expected {
			msg := "Didn't catch the out of range error."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
		}
	}
}

func TestTodosDeleteToCheckReturn(t *testing.T) {
	newTodo := makeTodo("test delete todo")
	todos := Todos{
		newTodo,
		makeTodo("message 2"),
	}
	expected := newTodo
	actual, err := todos.Delete(1)
	if err != nil {
		msg := "Caught the unexpected out of range error. Didn't delete a todo from todo list."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, nil, err)
	}
	if actual != expected {
		msg := "Didn't return a deleted todo."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

// ***************************************
// Test Todos.Check
// ***************************************

func TestTodosCheckToCheckComplete(t *testing.T) {
	todos := Todos{
		{Text: "text", IsComplete: false},
		{Text: "message", IsComplete: true},
	}
	testCases := []struct {
		num      int
		expected bool
		msg      string
	}{
		{1, true, "Didn't check a todo in todo list."},
		{2, false, "Didn't uncheck a todo in todo list."},
	}

	for i, tc := range testCases {
		_, err := todos.Check(tc.num)
		if err != nil {
			msg := "Caught the unexpected out of range error. Didn't check or uncheck a todo in todo list."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, nil, err)
		}
		actual := todos[i].IsComplete
		expected := tc.expected
		if actual != expected {
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", tc.msg, expected, actual)
		}
	}
}

func TestTodosCheckToCheckOutOfRangeError(t *testing.T) {
	todos := Todos{
		{Text: "text", IsComplete: false},
	}
	expected := fmt.Sprint(errors.New("out of range of todo list"))
	_, err := todos.Check(3)
	actual := fmt.Sprint(err)
	if actual != expected {
		msg := "Didn't return out of range error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestTodosCheckToCheckReturn(t *testing.T) {
	todo := makeTodo("message")
	todos := Todos{
		todo,
	}

	actual, err := todos.Check(1)
	if err != nil {
		msg := "Caught the unexpected out of range error. Didn't check or uncheck a todo in todo list."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, nil, err)
	}
	todo.IsComplete = true
	expected := todo
	if actual != expected {
		msg := "Didn't return a checked todo."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

// ***************************************
// Test Todos.Star
// ***************************************

func TestTodosStarToCheckStarred(t *testing.T) {
	todos := Todos{
		{Text: "text", IsStarred: false},
		{Text: "message", IsStarred: true},
	}
	testCases := []struct {
		num      int
		expected bool
		msg      string
	}{
		{1, true, "Didn't star a todo in todo list."},
		{2, false, "Didn't unstar a todo in todo list."},
	}

	for i, tc := range testCases {
		_, err := todos.Star(tc.num)
		if err != nil {
			msg := "Caught the unexpected out of range error. Didn't star or unstar a todo in todo list."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, nil, err)
		}
		actual := todos[i].IsStarred
		expected := tc.expected
		if actual != expected {
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", tc.msg, expected, actual)
		}
	}
}

func TestTodosStarToCheckOutOfRangeError(t *testing.T) {
	todos := Todos{
		makeTodo("text"),
	}
	expected := fmt.Sprint(errors.New("out of range of todo list"))
	_, err := todos.Star(3)
	actual := fmt.Sprint(err)
	if actual != expected {
		msg := "Didn't return out of range error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestTodosStarToCheckReturn(t *testing.T) {
	todo := makeTodo("message")
	todos := Todos{
		todo,
	}

	actual, err := todos.Star(1)
	if err != nil {
		msg := "Caught the unexpected out of range error. Didn't star or unstar a todo in todo list."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, nil, err)
	}
	todo.IsStarred = true
	expected := todo
	if actual != expected {
		msg := "Didn't return a starred todo."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

// ***************************************
// Test Todos.Edit
// ***************************************

func TestTodosEditToCheckEdited(t *testing.T) {
	todos := Todos{
		makeTodo("message"),
	}
	text := "edited text"
	_, err := todos.Edit(1, text)
	if err != nil {
		msg := "Caught the unexpected out of range error. Didn't edit a todo in todo list."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, nil, err)
	}
	actual := todos[0].Text
	expected := text
	if actual != expected {
		msg := "Didn't edit a text of todo."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestTodosEditToCheckOutOfRangeError(t *testing.T) {
	todos := Todos{
		makeTodo("message"),
	}
	expected := fmt.Sprint(errors.New("out of range of todo list"))
	_, err := todos.Edit(3, "edited text")
	actual := fmt.Sprint(err)
	if actual != expected {
		msg := "Didn't return out of range error."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}

func TestTodosEditToCheckReturn(t *testing.T) {
	todo := makeTodo("message")
	todos := Todos{
		todo,
	}

	text := "edited text"
	actual, err := todos.Edit(1, text)
	if err != nil {
		msg := "Caught the unexpected out of range error. Didn't star or unstar a todo in todo list."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, nil, err)
	}
	todo.Text = text
	todo.Timestamp = time.Now()
	expected := todo
	if !(actual.Timestamp.After(expected.Timestamp.Add(-time.Second)) && actual.Timestamp.Before(expected.Timestamp.Add(time.Second))) || actual.Text != expected.Text {
		msg := "Didn't return an edited todo."
		t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
	}
}
