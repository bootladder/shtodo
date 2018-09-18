package main

import (
  "github.com/spf13/viper"
  "errors"
  "bytes"
)

type Config struct {

}

func (c * Config) ParseString(input string) {

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
