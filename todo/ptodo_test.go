package todo_test

import (
	"testing"

	"github.com/robsongomes/tri/todo"
)

func TestPrivateFunc(t *testing.T) {
	result := todo.PrivateFunc()
	if result != "This is a private func" {
		t.Errorf("Could not call private func")
	}
}
