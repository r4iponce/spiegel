package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.gnous.eu/ada/spiegel/internal/config"
	"git.gnous.eu/ada/spiegel/internal/cron"
	"git.gnous.eu/ada/spiegel/internal/git"
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

	cron.Launch(time.Minute, initConfig.RepoList)

	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal
	logrus.Info("Bye!")
}
