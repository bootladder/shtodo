package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func Test_ParseConfigString_Empty_Panics(t *testing.T) {
    assert.Panics(t, func() {
        ParseConfigString("")
    }, "Should panic on empty input but did not panic")
}

func Test_ParseConfigString_ValidField_DoesNotPanic(t *testing.T) {

    var myInput string =
`
# A sample TOML config file.
[development]
enabled = true
`
    assert.NotPanics(t, func() {
        ParseConfigString(myInput)
    }, "Should not panic on Valid Field but did panic")
}

func Test_ParseConfigString_InvalidTOML_Panics(t *testing.T) {

    var myInput string =
`
z A sample TOML config file.
\development]
port = # 8080
`
    assert.Panics(t, func() {
        ParseConfigString(myInput)
    }, "Should panic on Invalid TOML but did not panic")
}

func Test_GetPathToTodo_ValidConfig_ReturnsCorrectValue(t *testing.T) {

    var myInput string =
`
# A sample TOML config file.
[todo]
path = "/home/hello/todo.txt"
`

    ParseConfigString(myInput)
    var str string = GetPathToTodo()
    if str != "/home/hello/todo.txt" {
        t.Fatalf("Did not get the correct path from the config file, got %s",str)
    }
}
