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

    inject()

    //read config to string
    //parse config string to struct

    var todoContents string = ReadTodo(pathtotodo)

    TouchLastTimeFile()
    var tbefore time.Time = ReadLastPrintedTodoTime(pathtolasttime)

    var tnow time.Time = time.Now().UTC()

    if ShouldPrint(tnow,tbefore,30) {
      fmt.Print(todoContents)
      UpdateLastTimeFile(tnow)
    }
}

func UpdateLastTimeFile(tnow time.Time) {
    currentTimeString := tnow.Format(layout)
    fileHandle, _ := os.Create(pathtolasttime)
    defer fileHandle.Close()
    writer := bufio.NewWriter(fileHandle)
    fmt.Fprint(writer, currentTimeString)
    writer.Flush()
}

func TouchLastTimeFile() {
    var f, err = os.OpenFile(pathtolasttime, os.O_RDONLY|os.O_CREATE, 0666)
    Fatal(err,"Open pathtolasttime")
    f.Close()
}
