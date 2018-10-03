package main

import (
    "fmt"
    "time"
)

func topflow() {

    var myConfig = &Config{}
    var err = myConfig.ParseConfigFile("/etc/shtodo.conf")
    Fatal(err)

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
