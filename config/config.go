package config

import (
  "io/ioutil"
  "encoding/json"
  "fmt"
  "os"
)

type DBConfig struct {
  Username string
  Password string
}

type Cfg struct {
  HttpPort int
  MySQL    DBConfig
}

var Config *Cfg

func init() {
  Config = new(Cfg)
  raw, err := ioutil.ReadFile("./config.json")

  if err != nil {
    fmt.Println("config" + err.Error())
    os.Exit(1)
  }

  json.Unmarshal(raw, &Config)
}