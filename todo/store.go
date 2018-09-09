package todo

import (
	"encoding/json"
	"io/ioutil"
)

func readTodosFromJSON(file string) ([]Todo, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return []Todo{}, err
	}
	var todos []Todo
	err = json.Unmarshal(bytes, &todos)
	return todos, err
}

func writeTodosToJSON(file string, todos []Todo) error {
	bytes, err := json.Marshal(&todos)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, bytes, 0644)
}

// LoadTodos reads data from json file to Todos
func LoadTodos(file string) (Todos, error) {
	return readTodosFromJSON(file)
}

// SaveTodos writes Todos to json file
func SaveTodos(file string, todos Todos) error {
	return writeTodosToJSON(file, todos)
}
