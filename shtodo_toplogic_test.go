package main

import (
  "testing"
  "fmt"
  "io"
  "errors"
  "bytes"
)

/////////////////////////////////////
func init() {
  fmt.Println("Test Suite Init")
}

var mockopen_ioreader io.Reader = nil
var mockopen_error error = nil

func MockOpen(filename string) (io.Reader,error) {
  return mockopen_ioreader,mockopen_error
}

func Test_OpenTodo_NoSuchFile_ReturnsError(t *testing.T) {

  external.open = MockOpen
  mockopen_error = errors.New("NoSuchFile")
  _,err := OpenTodo("nosuchfilename.txt")
  if err == nil {
    t.Fatalf("Expected Error, got nil error")
  }
}

func Test_OpenTodo_FileOpenedOK_ReturnsOK(t *testing.T) {

  //Should return a byte slice and a nil error
  external.open = MockOpen
  mockopen_error = nil
  mockopen_ioreader = bytes.NewBuffer([]byte{1,2})
  b,err := OpenTodo("dummyOKfilename.txt")
  if err != nil {
    t.Fatalf("Expected No Error, got some error")
  }
  if b == nil {
    t.Fatalf("Expected Byte Slice, got nil")
  }
}

func Test_ReadTodo_ErrorOpening_ReturnsError(t *testing.T) {
  external.open = MockOpen
  mockopen_error = errors.New("error opening file")
  mockopen_ioreader = bytes.NewBuffer([]byte{1,2})

  _,err := ReadTodo("dummyfilenotexists.txt")
  if err == nil {
    t.Fatalf("Expected Error, got nil")
  }
}

func Test_ReadTodo_EmptyFile_ReturnsSpecialString_NoError(t *testing.T) {
  external.open = MockOpen
  mockopen_error = nil
  mockopen_ioreader = bytes.NewBufferString("")

  str,err := ReadTodo("dummyOKfile.txt")
  if str != "Nothing To Do!" {
    t.Fatalf("Expected Special String, got %s",str)
  }
  if err != nil{
    t.Fatalf("Expected Nil Error, got some error")
  }
}

