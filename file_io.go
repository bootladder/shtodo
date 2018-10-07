package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

var pathToLastPrintTime = "/tmp/lasttime.txt"
var pathToLastPullTime = "/tmp/lastpulltime.txt"

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

func readTimeFromFile(filename string) time.Time {

	reader, err := readFile(filename)
	Fatal(err, "readTimeFromFile: "+filename)

	str := string(reader)
	if str == "" {
		return time.Time{}
	}

	//Convert the string to a time.Time
	tnow, err := time.Parse(layout, str)
	Fatal(err, "readTimeFromFile: time.Parse")
	return tnow
}

func readLastPullTime() time.Time {
	return time.Time{}
}

/*
  Not Tested
*/
func updateLastTimeFile(tnow time.Time) {
	currentTimeString := tnow.Format(layout)
	fileHandle, _ := os.Create(pathToLastPrintTime)
	defer fileHandle.Close()
	writer := bufio.NewWriter(fileHandle)
	fmt.Fprint(writer, currentTimeString)
	writer.Flush()
}

func touch(filename string) {
	var f, err = os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	Fatal(err, "Open "+filename)
	f.Close()
}
