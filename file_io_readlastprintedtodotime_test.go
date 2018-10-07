package main

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/////////////////////////////////////
func init() {
	fmt.Println("Test Suite Init")
}

func Test_ReadLastPrintedTodoTime_ErrorOpening_ReturnsError(t *testing.T) {

	usingMockReadFileFail(errors.New("error opening file"))

	assert.Panics(t, func() {
		readLastPrintedTodoTime("dummyfilenotexists.txt")
	}, "MockReadFile set to fail, but ReadLastPrintedTodoTime did not panic")
}

func Test_ReadLastPrintedTodoTime_InvalidFormat_ReturnsError(t *testing.T) {

	usingMockReadFileSuccess([]byte("300/28-20160 z:31:46 PM"))

	assert.Panics(t, func() {
		readLastPrintedTodoTime("dummyfilenotexists.txt")
	}, "Invalid format in time string, but did not panic")
}

func Test_ReadLastPrintedTodoTime_ValidFormat_DoesNotPanic(t *testing.T) {

	usingMockReadFileSuccess([]byte("02/28/2016 9:31:46 PM"))

	assert.NotPanics(t, func() {
		readLastPrintedTodoTime("dummyfilenotexists.txt")
	}, "Should Not Panic on valid time but did panic")
}

func Test_ReadLastPrintedTodoTime_EmptyFile_ReturnsLongTimeAgo(t *testing.T) {

	usingMockReadFileSuccess([]byte(""))

	var t1 = readLastPrintedTodoTime("dummyfilenotexists.txt")
	ttest := time.Time{}

	if t1 != ttest {
		t.Fatalf("Expected Unix Time = 0, got: %v\n", t1.Unix())
	}
}
