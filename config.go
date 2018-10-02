package main

import (
  "github.com/spf13/viper"
  "errors"
  "bytes"
)

type Config struct {

}

func (c * Config) ParseConfigFile(filename string) error {

    var b, err = external.readfile(filename)
    if err != nil {
        return err
    }

    err = c.ParseString(string(b))
    return err
}

func (c * Config) ParseString(input string) error {

    viper.SetConfigType("yaml")
    var err = viper.ReadConfig(bytes.NewBufferString(input))
    if err != nil {
      return err
    }

    if input == "" {
      return errors.New("hello")
    }
    return nil
}

func (c * Config) GetPathToTodo() string {
    return viper.GetString("todopath")
}
