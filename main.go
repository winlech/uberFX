package main

import (
	"aprendendoUberfx/httphandler"
	"context"
	"go.uber.org/fx"
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

func ProvideConfig() *Config {
	conf := Config{}
	data, _ := ioutil.ReadFile("config/base.yaml")

	_ = yaml.Unmarshal([]byte(data), &conf)

	return &conf
}

func ProvideLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()

	return slogger
}

func main() {
	fx.New(
		fx.Provide(ProvideConfig),
		fx.Provide(ProvideLogger),
		fx.Provide(http.NewServeMux),
		fx.Invoke(httphandler.New),
		fx.Invoke(registerHooks),
	).Run()

}

func registerHooks(lifecycle fx.Lifecycle, logger *zap.SugaredLogger, cfg *Config, mux *http.ServeMux) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					err := http.ListenAndServe(cfg.ApplicationConfig.Address, mux)
					if err != nil {
						logger.Error("Something went wrong while connecting the server")
					}
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		})
}
