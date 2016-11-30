package config

import (
  "io/ioutil"
  "encoding/json"
  "fmt"
  "log"
)

type DBConfig struct {
  Address  string
  Database string
  Password string
  Port     int
  Username string
}

type ApiConfig struct {
  BasicAuth   string
  Parallel    int
  RPS         int
}

type Cfg struct {
  EveAPI   ApiConfig
  HttpPort int
  MySQL    DBConfig
}

var Config *Cfg

func (c DBConfig) String() string {
  return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Username,
    c.Password, c.Address, c.Port, c.Database)
}

func init() {
  Config = new(Cfg)
  raw, err := ioutil.ReadFile("./config.json")

  if err != nil {
    log.Fatal("config" + err.Error())
  }

  json.Unmarshal(raw, &Config)
}