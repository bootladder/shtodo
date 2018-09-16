package main

import (
  "github.com/pkg/errors"
)

//Fatal panics on error
//First parameter of msgs is used each following variadic arg is dropped
func Fatal(err error, msgs ...string) {
  if err != nil {
    var str string
    for _, msg := range msgs {
      str = msg
      break
    }
    panic(errors.Wrap(err, str))
  }
}
