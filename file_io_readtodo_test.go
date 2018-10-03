package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "errors"
)

func Test_ReadTodo_NoExternalReadFileDefined_Panics(t *testing.T) {
  external.readfile = nil

  assert.Panics(t, func(){ ReadTodo("dummy.txt") })
}

func Test_ReadTodo_ErrorOpening_ReturnsError(t *testing.T) {

  usingMockReadFile_Fail(errors.New("error opening file"))

  assert.Panics(t, func(){ ReadTodo("dummy.txt") },
      "MockOpen set to error, should panic but did not")
}

func Test_ReadTodo_EmptyFile_ReturnsSpecialString_NoPanic(t *testing.T) {
  usingMockReadFile_Success([]byte(""))

  var str string
  assert.NotPanics(t, func(){
      str = ReadTodo("dummy.txt")
  }, "MockOpen set to success, ReadTodo should not panic, but it did")

  if str != "Nothing To Do!" {
    t.Fatalf("Expected Special String, got %s",str)
  }
}

