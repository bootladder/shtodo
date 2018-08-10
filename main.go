package main

// Read Config File , parse out fields
// Read Todo File
// Get current time
// Read file containing last printed time 
// Populate a shtodo struct
// Call myshtodo.Run()

import (
    "fmt"
    "io"
    "os"
    "time"
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
    var str,err = ReadTodo("/tmp/blah.txt")
    if err != nil {
      fmt.Printf("Error:  ReadTodo: %v\n",err)
      return
    }

    var tnow time.Time = time.Now()
    tbefore, err := ReadLastPrintedTodoTime("/tmp/lasttime.txt")
    if err != nil {
      fmt.Printf("Error:  ReadLastPrintedTodoTime: %v\n",err)
      return
    }

    if ShouldPrint(tnow,tbefore,30) {
      fmt.Printf("%s",str)
    }
}
