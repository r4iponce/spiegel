package git

import (
	"errors"
	"io"
	"os"
	"strings"

	"git.gnous.eu/ada/spiegel/internal/utils"
	goGit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/serverinfo"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/sirupsen/logrus"
)

func StartClone(repoList []Config) {
	logrus.Debug("Start first repository clone")
	for _, content := range repoList {
		_, err := os.Stat(content.FullPath)
		if os.IsNotExist(err) {
			content.fullClone()
		}
	}
}

func (c Config) fullClone() {
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

	if !utils.IsHTTPRepo(c.URL) {
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

	err = c.UpdateInfo()
	if err != nil {
		logrus.Error(err)
	}
}

func (c Config) Update() {
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

	if !utils.IsHTTPRepo(c.URL) {
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

	err = c.UpdateInfo()
	if err != nil {
		logrus.Error(err)
	}
}

func (c Config) UpdateInfo() error {
	r, err := goGit.PlainOpen(c.FullPath)
	if err != nil {
		return err
	}

	err = serverinfo.UpdateServerInfo(r.Storer, r.Storer.(*filesystem.Storage).Filesystem()) //nolint: forcetypeassert
	if err != nil {
		return err
	}

	return nil
}
