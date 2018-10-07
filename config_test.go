package main

import (
	"errors"
	"testing"
)

func Test_ParseConfigFile_NoReadFileDefined_ReturnsError(t *testing.T) {

	external.readfile = nil
	_, err := readFile("doesntmatter.txt")
	if err == nil {
		t.Fatalf("Must Fail with no dependency defined.  Did not fail")
	}
}

func Test_ParseConfigFile_FailToRead_ReturnsError(t *testing.T) {

	usingMockReadFileFail(errors.New("FailToRead"))

	var myConfig = &config{}
	var err = myConfig.parseConfigFile("dummyfile.txt")
	if err == nil {
		t.Fatalf("Expected error: Fail to Read Config File")
	}
}

func Test_ParseConfigFile_ConfigFileOK_NilError(t *testing.T) {

	var myInput = `
# A sample YAML config file.
todopath: /tmp/hello
`
	usingMockReadFileSuccess([]byte(myInput))

	var myConfig = &config{}
	var err = myConfig.parseConfigFile("dummyfile.txt")
	if err != nil {
		t.Fatalf("Expected nil error on OK config file")
	}
}

func Test_ParseConfigString_Empty_ReturnsError(t *testing.T) {
	var myConfig = &config{}
	var err = myConfig.parseString("")
	if err == nil {
		t.Fatalf("Expected error on empty input")
	}
}

func Test_ParseConfigFile_InvalidConfig_ReturnsError(t *testing.T) {

	var myInput = `
z A invalid config file
\development]
port = # 8080
`
	usingMockReadFileSuccess([]byte(myInput))

	var myConfig = &config{}
	var err = myConfig.parseConfigFile("blah")
	if err == nil {
		t.Fatalf("Expected error on Invalid Config")
	}
}

func Test_ValidConfig_ValuesCanBeQueried(t *testing.T) {

	var myConfig = &config{}
	var myInput = `
# A sample YAML config file.
todopath: "/home/hello/todo.txt"
todointerval: 30
pushinterval: 30
pullinterval: 30
`
	var str string
	var i int

	var err = myConfig.parseString(myInput)
	if err != nil {
		t.Fatalf("Expected No Error on Valid Config")
	}

	str = myConfig.getPathToTodo()
	if str != "/home/hello/todo.txt" {
		t.Fatalf("Did not get the correct path from the config file, got %s", str)
	}

	i = myConfig.getPrintInterval()
	if i != 30 {
		t.Fatalf("Did not get the correct path from the config file, got %s", str)
	}

	i = myConfig.getPushInterval()
	if i != 30 {
		t.Fatalf("Did not get the correct path from the config file, got %s", str)
	}

	i = myConfig.getPullInterval()
	if i != 30 {
		t.Fatalf("Did not get the correct path from the config file, got %s", str)
	}
}
