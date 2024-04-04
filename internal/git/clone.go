package git

import (
	"errors"
	"io"
	"log"
	"os"

	goGit "github.com/go-git/go-git/v5"
	"github.com/sirupsen/logrus"
)

func StartClone(repoList []RepoConfig) {
	logrus.Debug("Start first repository clone")
	for _, content := range repoList {
		_, err := os.Stat(content.FullPath)
		if os.IsNotExist(err) {
			content.fullClone()
		}
	}
}

func (config RepoConfig) fullClone() {
	logrus.Debug("Clone ", config.Name, "...")
	logger := logrus.New()
	w := logger.Writer()
	defer func(w *io.PipeWriter) {
		err := w.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(w)

	_, err := goGit.PlainClone(config.FullPath, true, &goGit.CloneOptions{
		URL:      config.URL,
		Progress: w,
		Mirror:   true,
	})
	if err != nil {
		log.Panic(err)
	}

}

func (config RepoConfig) Update() {
	repo, err := goGit.PlainOpen(config.FullPath)
	if err != nil {
		logrus.Error(err)

		return
	}

	logrus.Debug("Clone ", config.Name, "...")
	logger := logrus.New()
	w := logger.Writer()
	defer func(w *io.PipeWriter) {
		err := w.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(w)

	err = repo.Fetch(&goGit.FetchOptions{
		Progress: w,
	})
	if err != nil {
		if errors.Is(err, goGit.NoErrAlreadyUpToDate) {
			logrus.Info(config.Name, " is already up-to-date")
		} else {
			logrus.Error(err)
		}
	}
}
