package main

import (
  "testing"
//  "github.com/stretchr/testify/assert"
  "errors"
)

func Test_ParseConfigFile_NoReadFileDefined_ReturnsError(t *testing.T) {

  external.readfile = nil
  _,err := ReadFile("doesntmatter.txt")
  if err == nil {
    t.Fatalf("Must Fail with no dependency defined.  Did not fail")
  }
}

func Test_ParseConfigFile_FailToRead_Panics(t *testing.T) {

    external.readfile = MockReadFile
    mockreadfile_error = errors.New("FailToRead")
    var myConfig = &Config{}
    var err = myConfig.ParseConfigFile("dummyfile.txt")
    if err == nil {
      t.Fatalf("Expected error: Fail to Read Config File")
    }
}

func Test_ParseConfigFile_ConfigFileOK_DoesNotPanic(t *testing.T) {

    var myInput string =
`
# A sample YAML config file.
todopath: /tmp/hello
`
    external.readfile = MockReadFile
    mockreadfile_error = nil
    mockreadfile_bytes = []byte(myInput)

    var myConfig = &Config{}
    var err = myConfig.ParseConfigFile("dummyfile.txt")
    if err != nil {
      t.Fatalf("Expected nil error on OK config file")
    }
}

func Test_ParseConfigString_Empty_Panics(t *testing.T) {
    var myConfig = &Config{}
    var err = myConfig.ParseString("")
    if err == nil {
      t.Fatalf("Expected error on empty input")
    }
}

func Test_ParseConfigFile_InvalidConfig_Panics(t *testing.T) {

    var myInput string =
`
z A invalid config file
\development]
port = # 8080
`
    external.readfile = MockReadFile
    mockreadfile_error = nil
    mockreadfile_bytes = []byte(myInput)

    var myConfig = &Config{}
    var err = myConfig.ParseConfigFile("blah")
    if err == nil {
      t.Fatalf("Expected error on Invalid Config")
    }
}

func Test_GetPathToTodo_ValidConfig_ReturnsCorrectValue(t *testing.T) {

    var myConfig = &Config{}
    var myInput string =
`
# A sample YAML config file.
todopath: "/home/hello/todo.txt"
todointerval: 30
pushinterval: 30
pullinterval: 30
`
    var str string
    var i int

    var err = myConfig.ParseString(myInput)
    if err != nil {
      t.Fatalf("Expected No Error on Valid Config")
    }

    str = myConfig.GetPathToTodo()
    if str != "/home/hello/todo.txt" {
        t.Fatalf("Did not get the correct path from the config file, got %s",str)
    }

    i = myConfig.GetTodoInterval()
    if i != 30 {
        t.Fatalf("Did not get the correct path from the config file, got %s",str)
    }

    i = myConfig.GetPushInterval()
    if i != 30 {
        t.Fatalf("Did not get the correct path from the config file, got %s",str)
    }

    i = myConfig.GetPullInterval()
    if i != 30 {
        t.Fatalf("Did not get the correct path from the config file, got %s",str)
    }
}
