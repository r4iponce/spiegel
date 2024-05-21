package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
	"golang.org/x/sys/unix"
)

func LoadToml(file string) (Config, error) {
	var c Config

	source, err := os.ReadFile(file)
	if err != nil {
		return c, errConfigFileNotReadable
	}

	err = toml.Unmarshal(source, &c)
	if err != nil {
		panic(err)
	}

	fillFullPath(&c)
	fillFullPath(&c)

	return c, nil
}

func fillFullPath(c *Config) {
	for i, content := range c.RepoList {
		c.RepoList[i].FullPath = c.CloneDirectory + "/" + content.Name
	}
}

func (c Config) Verify() error {
	allowedValue := []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	found := false
	for _, v := range allowedValue {
		if v == c.Log.Level {
			found = true
		}
	}

	if !found {
		return errLogLevel
	}

	if unix.Access(c.CloneDirectory, unix.W_OK) != nil {
		return errCloneDirectoryUnwritable
	}

	return nil
}
