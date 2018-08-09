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
    "os"
)

var external = External{}
type External struct {
  open func(string) (io.Reader,error)
}

func osopen_wrapper(s string) (io.Reader,error) {
  return os.Open(s)
}

func inject() {
    external.open = osopen_wrapper
}

func main() {
    inject()
    log.Printf("hello main go %s\n")
    fmt.Printf("hello main go %s\n")
    var str,err = ReadTodo("/tmp/blah.txt")
    if err != nil {
      fmt.Printf("Error:  ReadTodo: %v\n",err)
      return
    }
    fmt.Printf("Todo:  %s\n",str)
    //myshtodo.Run()
}
