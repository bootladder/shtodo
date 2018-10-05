package main

import (
    "io/ioutil"
    "fmt"
    "time"
    "os"
    "os/exec"
    "path/filepath"
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
    todo_dir := filepath.Dir(myConfig.GetPathToTodo())
    command := "cd " + todo_dir + ";git commit -am \"hello $(date)\"; " + "git push"
    fmt.Print(command)
    bash_command(command)
}

func pulltodo() {
    todo_dir := filepath.Dir(myConfig.GetPathToTodo())
    command := "cd " + todo_dir + ";git pull"
    bash_command(command)
}
func edittodo() {
    command := "vi " + myConfig.GetPathToTodo()
    bash_command(command)
}

func bash_command(command string) {
    cmd := exec.Command("bash", "-c", command)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
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

