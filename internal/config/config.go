package config

import (
	"git.gnous.eu/ada/git-mirror/internal/git"
	"git.gnous.eu/ada/git-mirror/internal/log"
)

type Config struct {
	CloneDirectory string // Repository where gir-mirror keep repository
	Log            log.Config
	Interval       int // Update interval in minute
	RepoList       []git.RepoConfig
}
