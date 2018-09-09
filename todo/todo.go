package todo

import (
	"errors"
	"time"
)

// Todo contains info of todo
type Todo struct {
	Text       string    `json:"text"`
	Timestamp  time.Time `json:"timestamp"`
	IsStarred  bool      `json:"is_starred"`
	IsComplete bool      `json:"is_complete"`
}

// Check makes todo complete
func (todo *Todo) Check() {
	todo.IsComplete = true
}

// Uncheck makes todo incomplete
func (todo *Todo) Uncheck() {
	todo.IsComplete = false
}

// Star marks todo star
func (todo *Todo) Star() {
	todo.IsStarred = true
}

// Unstar unmarks todo star
func (todo *Todo) Unstar() {
	todo.IsStarred = false
}

// Edit changes text of todo and updates timestamp
func (todo *Todo) Edit(text string) {
	todo.Text = text
	todo.Timestamp = time.Now()
}

// NewTodo returns a new todo
func NewTodo(text string) Todo {
	todo := Todo{
		Text:      text,
		Timestamp: time.Now(),
	}
	return todo
}

// Todos is type of todo list
type Todos []Todo

// These receivers have argument n. n is 1~len(todos).

func (todos Todos) checkInRange(n int) error {
	if 1 <= n && n <= len(todos) {
		return nil
	}
	return errors.New("out of range of todo list")
}

// Get gets a todo from todo list
func (todos Todos) Get(n int) (Todo, error) {
	if err := todos.checkInRange(n); err != nil {
		return Todo{}, err
	}
	return todos[n-1], nil
}

// Add adds a todo to todo list
func (todos *Todos) Add(todo Todo) Todo {
	tmp := append(Todos{todo}, (*todos)[0:]...)
	*todos = append((*todos)[:0], tmp...)
	return todo
}

// Delete deletes a todo from todo list
func (todos *Todos) Delete(n int) (Todo, error) {
	if err := todos.checkInRange(n); err != nil {
		return Todo{}, err
	}
	n = n - 1
	todo := (*todos)[n]
	*todos = append((*todos)[:n], (*todos)[n+1:]...)
	return todo, nil
}

// Check checks or unchecks a todo in todo list
func (todos Todos) Check(n int) (Todo, error) {
	if err := todos.checkInRange(n); err != nil {
		return Todo{}, err
	}
	n = n - 1
	if todos[n].IsComplete {
		todos[n].Uncheck()
	} else {
		todos[n].Check()
	}
	return todos[n], nil
}

// Star stars or unstars a todo in todo list
func (todos Todos) Star(n int) (Todo, error) {
	if err := todos.checkInRange(n); err != nil {
		return Todo{}, err
	}
	n = n - 1
	if todos[n].IsStarred {
		todos[n].Unstar()
	} else {
		todos[n].Star()
	}
	return todos[n], nil
}

// Edit edits a text of todo in todo list
func (todos Todos) Edit(n int, text string) (Todo, error) {
	if err := todos.checkInRange(n); err != nil {
		return Todo{}, err
	}
	n = n - 1
	todos[n].Edit(text)
	return todos[n], nil
}
