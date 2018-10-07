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

func Test_readTimeFromFile_ErrorOpening_ReturnsError(t *testing.T) {

	usingMockReadFileFail(errors.New("error opening file"))

	assert.Panics(t, func() {
		readTimeFromFile("dummyfilenotexists.txt")
	}, "MockReadFile set to fail, but readTimeFromFile did not panic")
}

func Test_readTimeFromFile_InvalidFormat_ReturnsError(t *testing.T) {

	usingMockReadFileSuccess([]byte("300/28-20160 z:31:46 PM"))

	assert.Panics(t, func() {
		readTimeFromFile("dummyfilenotexists.txt")
	}, "Invalid format in time string, but did not panic")
}

func Test_readTimeFromFile_ValidFormat_DoesNotPanic(t *testing.T) {

	usingMockReadFileSuccess([]byte("02/28/2016 9:31:46 PM"))

	assert.NotPanics(t, func() {
		readTimeFromFile("dummyfilenotexists.txt")
	}, "Should Not Panic on valid time but did panic")
}

func Test_readTimeFromFile_EmptyFile_ReturnsLongTimeAgo(t *testing.T) {

	usingMockReadFileSuccess([]byte(""))

	var t1 = readTimeFromFile("dummyfilenotexists.txt")
	ttest := time.Time{}

	if t1 != ttest {
		t.Fatalf("Expected Unix Time = 0, got: %v\n", t1.Unix())
	}
}
