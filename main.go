package main

// Read Config File , parse out fields
// Read Todo File
// Get current time
// Read file containing LastPrintedTime.txt
// If ShouldPrint, then print
// Write current time to LastPrintedTime.txt

import (
    "fmt"
    "io"
    "os"
    "time"
    "bufio"
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

    //read config to string
    //parse config string to struct

    //touch the file, ie. create it if it doesn't exist
    os.OpenFile("/tmp/lasttime.txt", os.O_RDONLY|os.O_CREATE, 0666)

    var str,err = ReadTodo("/tmp/blah.txt")
    if err != nil {
      fmt.Printf("Error:  ReadTodo: %v\n",err)
      return
    }

    var tnow time.Time = time.Now().UTC()
    tbefore, err := ReadLastPrintedTodoTime("/tmp/lasttime.txt")
    if err != nil {
      fmt.Printf("Error:  ReadLastPrintedTodoTime: %v\n",err)
      return
    }

    if ShouldPrint(tnow,tbefore,30) {
      fmt.Printf("%s",str)

      //write current time string to file
      str := tnow.Format(layout)
      fileHandle, _ := os.Create("/tmp/lasttime.txt")
      writer := bufio.NewWriter(fileHandle)
      defer fileHandle.Close()

      fmt.Fprint(writer, str)
      writer.Flush()
    }
}
