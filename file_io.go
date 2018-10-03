package main

import (
    "os"
    "errors"
    "time"
    "fmt"
    "bufio"
)

var pathtolasttime string = "/tmp/lasttime.txt"

//format used in storing last printed time to a file
var layout = "01/02/2006 3:04:05 PM"

func ReadFile(filename string) ([]byte,error) {
  if external.readfile == nil {
    return nil,errors.New("ReadFile must have external.readfile() defined")
  }
  b,err := external.readfile(filename)
  return b,err
}

func ReadTodo(filename string) (string) {

  reader,err := ReadFile(filename)
  Fatal(err, "ReadTodo: ReadFile")
  str := string(reader)
  if str == "" {
    str = "Nothing To Do!"
  }
  return str
}

func ReadLastPrintedTodoTime(filename string) (time.Time) {

  reader,err := ReadFile(filename)
  Fatal(err, "ReadLastPrintedTodoTime: ReadFile")

  str := string(reader)
  if str == "" {
    return time.Time{}
  }

  //Convert the string to a time.Time
  tnow, err := time.Parse(layout, str)
  Fatal(err, "ReadLastPrintedTodoTime: time.Parse")
  return tnow
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
