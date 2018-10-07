package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

var pathtolasttime = "/tmp/lasttime.txt"

//format used in storing last printed time to a file
var layout = "01/02/2006 3:04:05 PM"

func readFile(filename string) ([]byte, error) {
	if external.readfile == nil {
		return nil, errors.New("ReadFile must have external.readfile() defined")
	}
	b, err := external.readfile(filename)
	return b, err
}

func readTodo(filename string) string {

	reader, err := readFile(filename)
	Fatal(err, "ReadTodo: ReadFile")
	str := string(reader)
	if str == "" {
		str = "Nothing To Do!"
	}
	return str
}

func readLastPrintedTodoTime(filename string) time.Time {

	reader, err := readFile(filename)
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

/*
  Not Tested
*/
func updateLastTimeFile(tnow time.Time) {
	currentTimeString := tnow.Format(layout)
	fileHandle, _ := os.Create(pathtolasttime)
	defer fileHandle.Close()
	writer := bufio.NewWriter(fileHandle)
	fmt.Fprint(writer, currentTimeString)
	writer.Flush()
}

func touchLastTimeFile() {
	var f, err = os.OpenFile(pathtolasttime, os.O_RDONLY|os.O_CREATE, 0666)
	Fatal(err, "Open pathtolasttime")
	f.Close()
}
