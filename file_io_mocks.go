package main

import (
  "io"
)

var mockopen_ioreader io.Reader = nil
var mockopen_error error = nil

var mockreadfile_bytes []byte = nil
var mockreadfile_error error = nil

func MockOpen(filename string) (io.Reader,error) {
  return mockopen_ioreader,mockopen_error
}

func MockReadFile(filename string) ([]byte,error) {
  return mockreadfile_bytes,mockreadfile_error
}
