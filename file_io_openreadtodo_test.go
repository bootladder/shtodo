package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "io"
  "errors"
  "bytes"
)

var mockopen_ioreader io.Reader = nil
var mockopen_error error = nil

func MockOpen(filename string) (io.Reader,error) {
  return mockopen_ioreader,mockopen_error
}

func Test_Open_NoExternalOpenDefined_ReturnsError(t *testing.T) {

  external.open = nil
  _,err := Open("doesntmatter.txt")
  if err == nil {
    t.Fatalf("Must Fail with no dependency defined.  Did not fail")
  }
}

func Test_Open_NoSuchFile_ReturnsError(t *testing.T) {

  external.open = MockOpen
  mockopen_error = errors.New("NoSuchFile")
  _,err := Open("nosuchfilename.txt")
  if err == nil {
    t.Fatalf("Expected Error, got nil error")
  }
}

func Test_Open_FileOpenedOK_ReturnsOK(t *testing.T) {

  //Should return a byte slice and a nil error
  external.open = MockOpen
  mockopen_error = nil
  mockopen_ioreader = bytes.NewBuffer([]byte{1,2})
  b,err := Open("dummyOKfilename.txt")
  if err != nil {
    t.Fatalf("Expected No Error, got some error")
  }
  if b == nil {
    t.Fatalf("Expected Byte Slice, got nil")
  }
}

func Test_ReadTodo_NoExternalOpenDefined_ReturnsError(t *testing.T) {
  external.open = nil
  mockopen_error = nil
  mockopen_ioreader = nil

  assert.Panics(t, func(){ ReadTodo("dummy.txt") })
}

func Test_ReadTodo_ErrorOpening_ReturnsError(t *testing.T) {
  external.open = MockOpen
  mockopen_error = errors.New("error opening file")
  mockopen_ioreader = bytes.NewBuffer([]byte{1,2})

  assert.Panics(t, func(){ ReadTodo("dummy.txt") },
      "MockOpen set to error, should panic but did not")
}

func Test_ReadTodo_EmptyFile_ReturnsSpecialString_NoPanic(t *testing.T) {
  external.open = MockOpen
  mockopen_error = nil
  mockopen_ioreader = bytes.NewBufferString("")

  var str string
  assert.NotPanics(t, func(){
      str = ReadTodo("dummy.txt")
  }, "MockOpen set to success, ReadTodo should not panic, but it did")

  if str != "Nothing To Do!" {
    t.Fatalf("Expected Special String, got %s",str)
  }
}

