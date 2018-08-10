package main

import (
  "testing"
  "time"
)

func Test_ShouldPrint_TooLongPassed_ReturnsTrue(t *testing.T) {

  tnow, _    := time.Parse(layout, "02/28/2016 9:31:46 PM")
  tbefore, _ := time.Parse(layout, "02/25/2013 9:31:46 PM")
  var z = ShouldPrint(tnow,tbefore,30)
  if z == false {
    t.Fatalf("Expected True, got false")
  }
}

func Test_ShouldPrint_TooShortPassed_ReturnsFalse(t *testing.T) {

  tnow, _    := time.Parse(layout, "02/28/2016 9:31:46 PM")
  tbefore, _ := time.Parse(layout, "02/28/2016 9:31:45 PM")
  var z = ShouldPrint(tnow,tbefore,30)
  if z == true {
    t.Fatalf("Expected false, got true")
  }
}
