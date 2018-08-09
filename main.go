package main

// Read Config File , parse out fields
// Read Todo File
// Get current time
// Read file containing last printed time 
// Populate a shtodo struct
// Call myshtodo.Run()

import (
    "log"
    "fmt"
    "io"
    "io/ioutil"
    //"errors"
)

type External struct {
  open func(string) (io.Reader,error)
}

var external = External{}

func OpenTodo(filename string) (io.Reader,error) {
  b,err := external.open(filename)
  return b,err
}
func ReadTodo(filename string) (string,error) {

  reader,err := OpenTodo(filename)
  b,_ := ioutil.ReadAll(reader)
  str := string(b)
  if str == "" {
    str = "Nothing To Do!"
  }
  return str,err
}
func main() {
    log.Printf("hello main go %s\n")
    fmt.Printf("hello main go %s\n")
    //myshtodo.Run()
}
