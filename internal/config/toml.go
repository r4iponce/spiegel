package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
	"golang.org/x/sys/unix"
)

func LoadToml(file string) (Config, error) {
	var config Config

	source, err := os.ReadFile(file)
	if err != nil {
		return config, errConfigFileNotReadable
	}

	err = toml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}

	fillFullPath(&config)

	return config, nil
}

func fillFullPath(config *Config) {
	for i, content := range config.RepoList {
		config.RepoList[i].FullPath = config.CloneDirectory + "/" + content.Name
	}
}

func (config Config) Verify() error {
	allowedValue := []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	found := false
	for _, v := range allowedValue {
		if v == config.Log.Level {
			found = true
		}
	}

	if !found {
		return errLogLevel
	}

	if unix.Access(config.CloneDirectory, unix.W_OK) != nil {
		return errCloneDirectoryUnwritable
	}

	return nil
}
