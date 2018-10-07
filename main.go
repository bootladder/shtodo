package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

var command string
var myConfig = &config{}

func main() {

	inject()
	topflow()
}

func topflow() {

	// Read Command Line Flags
	if len(os.Args[1:]) > 0 {
		command = os.Args[1]
	}

	var err = myConfig.parseConfigFile("/etc/shtodo.conf")
	Fatal(err)

	switch command {
	default: //no args.  print the todo
		if shouldPull(myConfig.getPullInterval()) {
			pulltodo()
			updateLastTimeFile(pathToLastPullTime)
		}
		if shouldPrint(myConfig.getPrintInterval()) {
			printtodo()
			updateLastTimeFile(pathToLastPrintTime)
		}
	case "push":
		pushtodo()
	case "pull":
		pulltodo()
	case "edit", "e":
		edittodo()
		pushtodo()
	}
}

func printtodo() {
	var todoContents = readTodo(myConfig.getPathToTodo())
	fmt.Print(todoContents)
}

func pushtodo() {
	todoDir := filepath.Dir(myConfig.getPathToTodo())
	command := "cd " + todoDir + ";git commit -am \"hello $(date)\"; " + "git push\n"
	fmt.Print(command)
	bashCommand(command)
}

func pulltodo() {
	todoDir := filepath.Dir(myConfig.getPathToTodo())
	command := "cd " + todoDir + ";git pull"
	bashCommand(command)
}
func edittodo() {
	command := "vi " + myConfig.getPathToTodo()
	bashCommand(command)
}

func bashCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

////////////////////////////////////////////////////////

var external = externalFuncs{}

type externalFuncs struct {
	readfile func(string) ([]byte, error)
}

func inject() {
	external.readfile = ioutil.ReadFile
}
