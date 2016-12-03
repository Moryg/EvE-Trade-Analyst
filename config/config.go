package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type DBConfig struct {
	Address  string
	Database string
	Password string
	Port     int
	Username string
}

type ApiConfig struct {
	BasicAuth string
	Parallel  int
	RPS       int
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
	apiURL := "https://crest-tq.eveonline.com"
	apiAuth := "https://login.eveonline.com"
	for _, arg := range os.Args {
		if arg == "--dev" {

			apiURL = "http://localhost:8888"
			apiAuth = "http://localhost:8888"
		}
	}
	os.Setenv("API", apiURL)
	os.Setenv("APIAuth", apiAuth)

	Config = new(Cfg)
	raw, err := ioutil.ReadFile("./config.json")

	if err != nil {
		log.Fatal("config" + err.Error())
	}

	json.Unmarshal(raw, &Config)
}
