package models

import (
    "fmt"
    "github.com/BurntSushi/toml"
)

type Config struct {
    Data struct{
        Organization    string
        User            string
        Ticket          string
    }
}

func (obj *Config) ReadConfig(pathFile string) {
    if _, err := toml.DecodeFile(pathFile, obj); err != nil {
        fmt.Println(err)
    }
}
