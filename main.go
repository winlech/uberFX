package main

import (
	"aprendendoUberfx/httphandler"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

type ApplicationConfig struct {
	Address string `yaml:"address"`
}

type Config struct {
	ApplicationConfig `yaml:"application"`
}

func main() {
	conf := &Config{}
	data, err := ioutil.ReadFile("config/base.yaml")

	_ = yaml.Unmarshal([]byte(data), &conf)

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	slogger := logger.Sugar()

	mux := http.NewServeMux()
	httphandler.New(mux, slogger)

	err = http.ListenAndServe(conf.ApplicationConfig.Address, mux)
	if err != nil {
		return
	}
}
