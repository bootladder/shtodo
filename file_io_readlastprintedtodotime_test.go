package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "fmt"
  "errors"
  "time"
)

/////////////////////////////////////
func init() {
  fmt.Println("Test Suite Init")
}

func Test_ReadLastPrintedTodoTime_ErrorOpening_ReturnsError(t *testing.T) {

  usingMockReadFile_Fail(errors.New("error opening file"))

  assert.Panics(t, func(){
      ReadLastPrintedTodoTime("dummyfilenotexists.txt")
  }, "MockReadFile set to fail, but ReadLastPrintedTodoTime did not panic")
}

func Test_ReadLastPrintedTodoTime_InvalidFormat_ReturnsError(t *testing.T) {

  usingMockReadFile_Success([]byte("300/28-20160 z:31:46 PM"))

  assert.Panics(t, func(){
      ReadLastPrintedTodoTime("dummyfilenotexists.txt")
  }, "Invalid format in time string, but did not panic")
}

func Test_ReadLastPrintedTodoTime_ValidFormat_DoesNotPanic(t *testing.T) {

  usingMockReadFile_Success([]byte("02/28/2016 9:31:46 PM"))

  assert.NotPanics(t, func(){
      ReadLastPrintedTodoTime("dummyfilenotexists.txt")
  }, "Should Not Panic on valid time but did panic")
}

func Test_ReadLastPrintedTodoTime_EmptyFile_ReturnsLongTimeAgo(t *testing.T) {

  usingMockReadFile_Success([]byte(""))

  var t1 = ReadLastPrintedTodoTime("dummyfilenotexists.txt")
  ttest := time.Time{}

  if t1 != ttest {
    t.Fatalf("Expected Unix Time = 0, got: %v\n",t1.Unix())
  }
}
