package main

import (
    "time"
)


func ShouldPrint(tnow,tbefore time.Time, duration int) bool {

  var delta float64 = tnow.Sub(tbefore).Seconds()
  if int(delta) >= duration {
    return true
  }
  return false
}
