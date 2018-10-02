package main

import (
    "fmt"
    "time"
)

// Read Config File , parse out fields
// Read Todo File
// Get current time
// Read file containing last printed time 
// Populate a shtodo struct
// Call myshtodo.Run()

func topflow() {
    //read config to string

    var myConfig = &Config{}
    var err = myConfig.ParseConfigFile("/etc/shtodo.conf")
    if err != nil {
        Fatal(err)
    }
    var path = myConfig.GetPathToTodo()
    //parse config string to struct

    var todoContents string = ReadTodo(path)

    TouchLastTimeFile()
    var tbefore time.Time = ReadLastPrintedTodoTime(pathtolasttime)

    var tnow time.Time = time.Now().UTC()

    if ShouldPrint(tnow,tbefore,30) {
      fmt.Print(todoContents)
      UpdateLastTimeFile(tnow)
    }
}
