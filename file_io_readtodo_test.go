package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadTodo_NoExternalReadFileDefined_Panics(t *testing.T) {
	external.readfile = nil

	assert.Panics(t, func() { readTodo("dummy.txt") })
}

func Test_ReadTodo_ErrorOpening_ReturnsError(t *testing.T) {

	usingMockReadFileFail(errors.New("error opening file"))

	assert.Panics(t, func() { readTodo("dummy.txt") },
		"MockOpen set to error, should panic but did not")
}

func Test_ReadTodo_EmptyFile_ReturnsSpecialString_NoPanic(t *testing.T) {
	usingMockReadFileSuccess([]byte(""))

	var str string
	assert.NotPanics(t, func() {
		str = readTodo("dummy.txt")
	}, "MockOpen set to success, ReadTodo should not panic, but it did")

	if str != "Nothing To Do!" {
		t.Fatalf("Expected Special String, got %s", str)
	}
}
