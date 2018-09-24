package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "errors"
)

func Test_ParseConfigFile_FailToRead_Panics(t *testing.T) {

    external.readfile = MockReadFile
    mockreadfile_error = errors.New("FailToRead")
    var myConfig = &Config{}
    assert.Panics(t, func() {
        myConfig.ParseConfigFile("dummyfile.txt")
    }, "Should panic on Fail To Read Config File but did not panic")
}

func Test_ParseConfigFile_ConfigFileOK_DoesNotPanic(t *testing.T) {

    var myInput string =
`
# A sample TOML config file.
[development]
enabled = true
`
    external.readfile = MockReadFile
    mockreadfile_error = nil
    mockreadfile_bytes = []byte(myInput)

    var myConfig = &Config{}
    assert.NotPanics(t, func() {
        myConfig.ParseConfigFile("dummyfile.txt")
    }, "Should not panic on OK config file, but did panic")
}

func Test_ParseConfigString_Empty_Panics(t *testing.T) {
    var myConfig = &Config{}
    assert.Panics(t, func() {
        myConfig.ParseString("")
    }, "Should panic on empty input but did not panic")
}

func Test_ParseConfigFile_InvalidTOML_Panics(t *testing.T) {

    var myInput string =
`
z A sample TOML config file.
\development]
port = # 8080
`
    external.readfile = MockReadFile
    mockreadfile_error = nil
    mockreadfile_bytes = []byte(myInput)

    var myConfig = &Config{}
    assert.Panics(t, func() {
        myConfig.ParseConfigFile("blah")
    }, "Should panic on Invalid TOML but did not panic")
}

func Test_GetPathToTodo_ValidConfig_ReturnsCorrectValue(t *testing.T) {

    var myConfig = &Config{}
    var myInput string =
`
# A sample TOML config file.
[todo]
path = "/home/hello/todo.txt"
`

    myConfig.ParseString(myInput)

    var str string = myConfig.GetPathToTodo()
    if str != "/home/hello/todo.txt" {
        t.Fatalf("Did not get the correct path from the config file, got %s",str)
    }
}
