package main

import (
	"git.gnous.eu/ada/git-mirror/internal/config"
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

	cloneConfig := initConfig.RepoList[0]
	cloneConfig.FullPath = initConfig.CloneDirectory + cloneConfig.Name
	cloneConfig.FullClone()
}
