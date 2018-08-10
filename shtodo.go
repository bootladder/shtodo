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
    "time"
//    "fmt"
//    "strconv"
)

func Open(filename string) (io.Reader,error) {
  if external.open == nil {
    return nil,errors.New("Open must have external.open() defined")
  }
  b,err := external.open(filename)
  return b,err
}
func ReadTodo(filename string) (string,error) {

  reader,err := Open(filename)
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

func ReadLastPrintedTodoTime(filename string) (time.Time,error) {

  reader,err := Open(filename)
  if err != nil {
    return time.Time{}, err
  }
  b,_ := ioutil.ReadAll(reader)
  str := string(b)
  if str == "" {
  }

  //Convert the string to a time.Time
  tnow, err := time.Parse(layout, str)
  return tnow,err
}
