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
	var configPath string

	switch len(os.Args) {
	case 2: //nolint:mnd
		configPath = os.Args[1]
	case 1:
		configPath = "config.toml"
	default:
		logrus.Fatal("Max 1 argument is valid.")
	}

	initConfig, err := config.LoadToml(configPath)
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

	duration := time.Duration(initConfig.Interval) * time.Minute
	cron.Launch(duration, initConfig.RepoList)

	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-quitSignal
	logrus.Info("Bye!")
}
