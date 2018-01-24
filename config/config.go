package config

import (
	"sync"
	"io/ioutil"
	"github.com/labstack/gommon/log"
	"encoding/json"
)

var config *Config
var once sync.Once

type Config struct {
	Port string `json:"port"`
}

func Get() Config{
	once.Do(loadConfig)
	return *config
}

func loadConfig(){
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Cannot load config file: %s", err)
	}
	config = &Config{}
	if err = json.Unmarshal(data, config);err!= nil{
		log.Fatalf("Corrupted data in config file: %s", err)
	}
}