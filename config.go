package main

import (
  "errors"
  "github.com/spf13/viper"
  "bytes"
)

func ParseConfigString(input string) {

    viper.SetConfigType("toml")
    var err = viper.ReadConfig(bytes.NewBufferString(input))
    Fatal(err)

    if input == "" {
      Fatal(errors.New("hello"),"hello")
    }
}

func GetPathToTodo() string {
    return viper.GetString("todo.path")
}
