package todo

import (
	"io/ioutil"
	"os"
	"testing"
)

func setupEmptyTodosFile(t *testing.T) (string, func()) {
	t.Helper()
	tmpFile, err := ioutil.TempFile("", "tmptest")
	if err != nil {
		t.Fatalf("Error occurred in setup process before testing.\n%s\n", err)
	}
	return tmpFile.Name(), func() { os.Remove(tmpFile.Name()) }
}

func setupTodosFile(t *testing.T) (string, func()) {
	t.Helper()
	file, teardown := setupEmptyTodosFile(t)
	jsonData := []byte(`[{"text":"message","timestamp":"0001-01-01T00:00:00Z","is_starred":true,"is_complete":false}]`)
	if err := ioutil.WriteFile(file, jsonData, 0644); err != nil {
		t.Fatalf("Error occurred in setup process before testing.\n%s\n", err)
	}
	return file, teardown
}
