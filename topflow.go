package main

import (
    "fmt"
    "time"
)

func topflow() {
    //read config to string
    //parse config string to struct

    var todoContents string = ReadTodo(pathtotodo)

    TouchLastTimeFile()
    var tbefore time.Time = ReadLastPrintedTodoTime(pathtolasttime)

    var tnow time.Time = time.Now().UTC()

    if ShouldPrint(tnow,tbefore,30) {
      fmt.Print(todoContents)
      UpdateLastTimeFile(tnow)
    }
}
