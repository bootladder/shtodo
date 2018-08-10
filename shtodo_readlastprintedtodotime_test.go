package main

import (
  "testing"
  "fmt"
//  "io"
  "errors"
  "bytes"
//  "time"
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

  _,err := ReadLastPrintedTodoTime("dummyfilenotexists.txt")
  if err == nil {
    t.Fatalf("Expected Error, got nil")
  }
}

func Test_ReadLastPrintedTodoTime_InvalidFormat_ReturnsError(t *testing.T) {
  external.open = MockOpen
  mockopen_error = nil
  mockopen_ioreader = bytes.NewBufferString("300/28-20160 z:31:46 PM")

  _,err := ReadLastPrintedTodoTime("dummyfilenotexists.txt")
  if err == nil {
    t.Fatalf("Expected Error, got nil")
  }
}

func Test_ReadLastPrintedTodoTime_ValidFormat_ReturnsNilError(t *testing.T) {
  external.open = MockOpen
  mockopen_error = nil
  mockopen_ioreader = bytes.NewBufferString("02/28/2016 9:31:46 PM")

  _,err := ReadLastPrintedTodoTime("dummyfilenotexists.txt")
  if err != nil {
    t.Fatalf("Expected Nil Error, got some error: %v\n",err)
  }
}

