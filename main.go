package main

import (
    "io"
    "io/ioutil"
    "os"
)

var external = External{}
type External struct {
  open func(string) (io.Reader,error)
  readfile func(string) ([]byte,error)
}

func osopen_wrapper(s string) (io.Reader,error) {
  return os.Open(s)
}

func inject() {
    external.open = osopen_wrapper
    external.readfile = ioutil.ReadFile
}

func main() {

    inject()
    topflow()
}
