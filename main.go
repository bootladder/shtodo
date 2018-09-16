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

var pathtolasttime string = "/tmp/lasttime.txt"
var pathtotodo string = "/home/steve/Documents/todo.txt"

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
    var  err error
    inject()

    //read config to string
    //parse config string to struct

    //touch the file, ie. create it if it doesn't exist
    _,err = os.OpenFile(pathtolasttime, os.O_RDONLY|os.O_CREATE, 0666)
    Fatal(err,"Open pathtolasttime")

    var str string
    str,err = ReadTodo(pathtotodo)
    Fatal(err,"ReadTodo")

    var tnow time.Time = time.Now().UTC()
    tbefore, err := ReadLastPrintedTodoTime(pathtolasttime)
    Fatal(err,"ReadLastPrintedTodoTime")

    if ShouldPrint(tnow,tbefore,30) {
      fmt.Printf("%s",str)

      //write current time string to file
      str := tnow.Format(layout)
      fileHandle, _ := os.Create(pathtolasttime)
      writer := bufio.NewWriter(fileHandle)
      defer fileHandle.Close()

      fmt.Fprint(writer, str)
      writer.Flush()
    }
}
