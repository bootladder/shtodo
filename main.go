package main

import (
    "io/ioutil"
    "fmt"
    "time"
    "os"
)

var command string = "print"
var myConfig = &Config{}

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

    // Read Command Line Flags
    if len(os.Args[1:]) > 0 {
        command = os.Args[1]
    }

    var err = myConfig.ParseConfigFile("/etc/shtodo.conf")
    Fatal(err)

    switch(command) {
        case "print":
            printtodo()
        case "push":
            pushtodo()
        case "pull":
            pulltodo()
        case "edit":
            edittodo()
    }
}

func printtodo() {

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

func pushtodo() {
    fmt.Print("yay push")
}

func pulltodo() {
    fmt.Print("yay push")
}
func edittodo() {
    fmt.Print("yay edit")
}
