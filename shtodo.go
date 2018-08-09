package main

// Read Config File , parse out fields
// Read Todo File
// Get current time
// Read file containing last printed time 
// Populate a shtodo struct
// Call myshtodo.Run()

import (
    "io"
    "io/ioutil"
    "errors"
)

func OpenTodo(filename string) (io.Reader,error) {
  if external.open == nil {
    return nil,errors.New("OpenTodo must have external.open() defined")
  }
  b,err := external.open(filename)
  return b,err
}
func ReadTodo(filename string) (string,error) {

  reader,err := OpenTodo(filename)
  if err != nil {
    return "fail", err
  }
  b,_ := ioutil.ReadAll(reader)
  str := string(b)
  if str == "" {
    str = "Nothing To Do!"
  }
  return str,err
}
