package main

import (
    "io/ioutil"
    "fmt"
    "time"
    "os"
    "os/exec"
)

var command string = "print"
var myConfig = &Config{}

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
        case "edit", "e":
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
    cmd := exec.Command("vim", myConfig.GetPathToTodo())
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    _ = cmd.Run()
}

////////////////////////////////////////////////////////

var external = External{}
type External struct {
    readfile func(string) ([]byte,error)
}

func inject() {
    external.readfile = ioutil.ReadFile
}

