package config_test

import (
	"testing"

	"git.gnous.eu/ada/git-mirror/internal/config"
)

func TestToml(t *testing.T) {
	t.Parallel()
	testLoadTomlValid(t)
}

func testLoadTomlValid(t *testing.T) {
	t.Helper()
	got, err := config.LoadToml("test_resources/valid.toml")
	if err != nil {
		t.Fatal("Cannot load config: ", err)
	}
	// {aa {WARN } [{aa yy} {bb yy}]}
	if got.CloneDirectory != "archive/" {
		t.Fatal("Invalid CloneDirectory: ", got.CloneDirectory)
	}

	if got.Log.Level != "WARN" {
		t.Fatal("Invalid log level: ", got.Log.Level)
	}

	if got.Log.File != "log.txt" {
		t.Fatal("Invalid log file: ", got.Log.File)
	}

	if got.RepoList[0].Name != "linux" {
		t.Fatal("Invalid first repo name: ", got.RepoList[0].Name)
	}

	if got.RepoList[0].URL != "https://github.com/torvalds/linux/" {
		t.Fatal("Invalid first repo url: ", got.RepoList[0].URL)
	}

	if got.RepoList[1].Name != "linuxtwo" {
		t.Fatal("Invalid second repo name: ", got.RepoList[1].Name)
	}

	if got.RepoList[1].URL != "https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git" {
		t.Fatal("Invalid second repo URL: ", got.RepoList[1].URL)
	}
}
