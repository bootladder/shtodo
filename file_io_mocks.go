package main

var mockreadfile_bytes []byte = nil
var mockreadfile_error error = nil

func MockReadFile(filename string) ([]byte,error) {
  return mockreadfile_bytes,mockreadfile_error
}

func usingMockReadFile_Success(myBytes []byte) {

  external.readfile = MockReadFile
  mockreadfile_error = nil
  mockreadfile_bytes = myBytes
}

func usingMockReadFile_Fail(err error) {

  external.readfile = MockReadFile
  mockreadfile_error = err
  mockreadfile_bytes = []byte("doesn't matter")
}
