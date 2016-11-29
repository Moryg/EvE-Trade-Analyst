package config

import (
  "io/ioutil"
  "encoding/json"
  "fmt"
  "os"
)

type DBConfig struct {
  Password string
  Username string
}

type ApiConfig struct {
  RPS         int
  Parallel    int
}

type Cfg struct {
  EveAPI   ApiConfig
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