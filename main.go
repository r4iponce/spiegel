package main

import (
	"git.gnous.eu/ada/git-mirror/internal/config"
	"git.gnous.eu/ada/git-mirror/internal/git"
	"github.com/sirupsen/logrus"
)

func main() {
	initConfig, err := config.LoadToml("config.example.toml")
	if err != nil {
		logrus.Fatal(err)
	}

	err = initConfig.Verify()
	if err != nil {
		logrus.Fatal(err)
	}

	initConfig.Log.Init()
	logrus.Info("Config loaded")
	logrus.Debug("Config: ", initConfig)

	git.StartClone(initConfig.RepoList)
}
