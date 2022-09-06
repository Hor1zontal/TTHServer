package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var Server  = &BaseConfig{}

type BaseConfig struct {
	Database DBConfig `yaml:"database"`

	HTTPAddress string `yaml:"http"`

	IsCors bool `yaml:"isCors"`
	//DefaultPWD string
}

type DBConfig struct {
	DB string `yaml:"db"`
	DBHost string `yaml:"DBHost"`
	DBUser string `yaml:"DBUser"`
	DBPassword string `yaml:"DBPassword"`
	DBName string `yaml:"DBName"`
	DBPort string `yaml:"DBPort"`
}

var (

	//tag   = ""
	configPath = ""
)

func init() {
	flag.StringVar(&configPath, "config", "./config.yml", "configuration path")
	flag.Parse()

}

func Init() {
	LoadConfigData(configPath, &Server)
}

func LoadConfigData(path string, config interface{}) {
	LoadConfigDataEx(path, config, true)
}

func LoadConfigDataEx(path string, config interface{}, fatal bool) {
	if config == nil {
		return
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		if fatal {
			log.Fatal("config file %v  is not found", path)
		}
		return
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		if fatal {
			log.Fatal("load config %v err %v", path, err)
		}
	}
}