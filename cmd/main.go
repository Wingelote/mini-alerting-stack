package main

import (
	"os"

	"github.com/wingelote/aisprid-alerting/internal/config"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

func main() {
	var conf config.Config
	reader, _ := os.ReadFile("config.toml")

	if err := toml.Unmarshal(reader, &conf); err != nil {
		logrus.Error(err)
	}

	NewCLI(conf)
}
