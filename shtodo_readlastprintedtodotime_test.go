package main

import (
  "testing"
  "fmt"
//  "io"
  "errors"
  "bytes"
  "time"
  "encoding/binary"
)

var layout = "01/02/2006 3:04:05 PM"

/////////////////////////////////////
func init() {
  fmt.Println("Test Suite Init")
}

func Test_ReadLastPrintedTodoTime_ErrorOpening_ReturnsError(t *testing.T) {
  external.open = MockOpen
  mockopen_error = errors.New("error opening file")
  mockopen_ioreader = bytes.NewBuffer([]byte{1,2})

  _,err := ReadLastPrintedTodoTime("dummyfilenotexists.txt")
  if err == nil {
    t.Fatalf("Expected Error, got nil")
  }
}

func Test_ReadLastPrintedTodoTime_InvalidFormat_ReturnsError(t *testing.T) {
  external.open = MockOpen
  mockopen_error = errors.New("error opening file")
  t1, _ := time.Parse(layout, "02/28/2016 9:31:46 PM")
  t_unix := t1.Unix()
  t_bytes := make([]byte,8)
  binary.PutVarint(t_bytes,t_unix)

  mockopen_ioreader = bytes.NewBuffer(t_bytes)

  _,err := ReadLastPrintedTodoTime("dummyfilenotexists.txt")
  if err == nil {
    t.Fatalf("Expected Error, got nil")
  }
}

func Test_ReadLastPrintedTodoTime_ValidFormat_ReturnsNilError(t *testing.T) {
  external.open = MockOpen
  mockopen_error = nil
  t1, _ := time.Parse(layout, "02/28/2016 9:31:46 PM")
  t_unix := t1.Unix()
  t_bytes := make([]byte,8)
  binary.PutVarint(t_bytes,t_unix)

  mockopen_ioreader = bytes.NewBuffer(t_bytes)

  _,err := ReadLastPrintedTodoTime("dummyfilenotexists.txt")
  if err != nil {
    t.Fatalf("Expected Nil Error, got some error: %v\n",err)
  }
}

