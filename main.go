package main

import (
    "io/ioutil"
    "fmt"
    "time"
)

var external = External{}
type External struct {
  readfile func(string) ([]byte,error)
}

func inject() {
    external.readfile = ioutil.ReadFile
}

func main() {

    inject()
    topflow()
}

func topflow() {

    var myConfig = &Config{}
    var err = myConfig.ParseConfigFile("/etc/shtodo.conf")
    Fatal(err)

    var path = myConfig.GetPathToTodo()

    var todoContents string = ReadTodo(path)

    TouchLastTimeFile()
    var tbefore time.Time = ReadLastPrintedTodoTime(pathtolasttime)

    var tnow time.Time = time.Now().UTC()

    if ShouldPrint(tnow,tbefore,30) {
      fmt.Print(todoContents)
      UpdateLastTimeFile(tnow)
    }
}
