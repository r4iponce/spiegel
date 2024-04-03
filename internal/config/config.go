package config

import "git.gnous.eu/ada/git-mirror/internal/log"

type Config struct {
	CloneDirectory string // Repository where gir-mirror keep repository
	Log            log.Config
	RepoList       []repoConfig
}

type repoConfig struct {
	URL  string // Source url
	Name string // Name of clone (directory name)
}
