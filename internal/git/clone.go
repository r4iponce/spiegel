package git

import (
	"errors"
	"git.gnous.eu/ada/spiegel/internal/utils"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"io"
	"os"
	"strings"

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

			return
		}
	}(w)

	repoConfig := &goGit.CloneOptions{
		URL:      c.URL,
		Progress: w,
		Mirror:   true,
	}

	if !utils.IsHttpRepo(c.URL) {
		key, err := os.ReadFile(c.SSHKey)
		if err != nil {
			logrus.Error(err)

			return
		}

		user := strings.Split(c.URL, "@")[0]
		url := strings.Split(c.URL, "@")[1]

		sshAuth, err := ssh.NewPublicKeys(user, key, "")
		if err != nil {
			logrus.Error(err)

			return
		}
		repoConfig.Auth = sshAuth
		repoConfig.URL = url
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

	fetchConfig := &goGit.FetchOptions{
		Progress: w,
	}

	if !utils.IsHttpRepo(c.URL) {
		key, err := os.ReadFile(c.SSHKey)
		if err != nil {
			logrus.Error(err)

			return
		}

		user := strings.Split(c.URL, "@")[0]

		sshAuth, err := ssh.NewPublicKeys(user, key, "")
		if err != nil {
			logrus.Error(err)

			return
		}
		fetchConfig.Auth = sshAuth
	}

	err = repo.Fetch(fetchConfig)
	if err != nil {
		if errors.Is(err, goGit.NoErrAlreadyUpToDate) {
			logrus.Info(c.Name, " is already up-to-date")
		} else {
			logrus.Error(err)
		}
	}
}
