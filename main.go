package main

import (
    "io/ioutil"
)

var external = External{}
type External struct {
  readfile func(string) ([]byte,error)
}

func inject() {
    external.readfile = ioutil.ReadFile
}

func main() {

    inject()
    topflow()
}
