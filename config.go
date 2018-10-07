package main

import (
	"bytes"
	"errors"

	"github.com/spf13/viper"
)

type config struct {
}

func (c *config) parseConfigFile(filename string) error {

	var b, err = external.readfile(filename)
	if err != nil {
		return err
	}

	err = c.parseString(string(b))
	return err
}

func (c *config) parseString(input string) error {

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

func (c *config) getPathToTodo() string {
	return viper.GetString("todopath")
}

func (c *config) getTodoInterval() int {
	return viper.GetInt("todointerval")
}

func (c *config) getPullInterval() int {
	return viper.GetInt("pullinterval")
}

func (c *config) getPushInterval() int {
	return viper.GetInt("pushinterval")
}
