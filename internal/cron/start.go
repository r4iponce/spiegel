package cron

import (
	"time"

	"git.gnous.eu/ada/git-mirror/internal/git"
	"github.com/sirupsen/logrus"
)

// start a regular background tasks
func start(duration time.Duration, fn func(), name string) {
	for {
		time.Sleep(duration)
		logrus.Info("Begin update job for: ", name)
		fn()
		logrus.Debug("Finished update job for: ", name)
	}
}

// Launch all repo update background tasks
func Launch(duration time.Duration, config []git.RepoConfig) {
	var counter int
	for _, content := range config {
		counter++
		logrus.Debug("Launch background tasks for: ", content.Name)
		go start(duration, content.Update, content.Name)
	}
	logrus.Info("Started: ", counter, " background tasks")
}
