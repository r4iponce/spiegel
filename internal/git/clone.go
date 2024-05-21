package git

import (
	"errors"
	"io"
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

func (c RepoConfig) fullClone() {
	logrus.Debug("Clone ", c.Name, "...")
	logger := logrus.New()
	w := logger.Writer()
	defer func(w *io.PipeWriter) {
		err := w.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(w)

	repoConfig := &goGit.CloneOptions{
		URL:      c.URL,
		Progress: w,
		Mirror:   true,
	}

	_, err := goGit.PlainClone(c.FullPath, true, repoConfig)
	if err != nil {
		logrus.Error(err)
	}
}

func (c RepoConfig) Update() {
	repo, err := goGit.PlainOpen(c.FullPath)
	if err != nil {
		logrus.Error(err)

		return
	}

	logrus.Debug("Clone ", c.Name, "...")
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
			logrus.Info(c.Name, " is already up-to-date")
		} else {
			logrus.Error(err)
		}
	}
}
