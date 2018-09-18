package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "fmt"
//  "io"
  "errors"
  "bytes"
  "time"
//  "encoding/binary"
)

/////////////////////////////////////
func init() {
  fmt.Println("Test Suite Init")
}

func Test_ReadLastPrintedTodoTime_ErrorOpening_ReturnsError(t *testing.T) {
  external.open = MockOpen
  mockopen_error = errors.New("error opening file")
  mockopen_ioreader = bytes.NewBufferString("contents don't matter")

  assert.Panics(t, func(){
      ReadLastPrintedTodoTime("dummyfilenotexists.txt")
  }, "MockOpen set to fail, but ReadTodo did not panic")
}

func Test_ReadLastPrintedTodoTime_InvalidFormat_ReturnsError(t *testing.T) {
  external.open = MockOpen
  mockopen_error = nil
  mockopen_ioreader = bytes.NewBufferString("300/28-20160 z:31:46 PM")

  assert.Panics(t, func(){
      ReadLastPrintedTodoTime("dummyfilenotexists.txt")
  }, "Invalid format in time string, but did not panic")
}

func Test_ReadLastPrintedTodoTime_ValidFormat_DoesNotPanic(t *testing.T) {
  external.open = MockOpen
  mockopen_error = nil
  mockopen_ioreader = bytes.NewBufferString("02/28/2016 9:31:46 PM")

  assert.NotPanics(t, func(){
      ReadLastPrintedTodoTime("dummyfilenotexists.txt")
  }, "Invalid format in time string, but did not panic")
}

func Test_ReadLastPrintedTodoTime_EmptyFile_ReturnsLongTimeAgo(t *testing.T) {
  external.open = MockOpen
  mockopen_error = nil
  mockopen_ioreader = bytes.NewBufferString("")

  var t1 = ReadLastPrintedTodoTime("dummyfilenotexists.txt")
  ttest := time.Time{}

  if t1 != ttest {
    t.Fatalf("Expected Unix Time = 0, got: %v\n",t1.Unix())
  }
}

