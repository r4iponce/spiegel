package config

import "errors"

var (
	errLogLevel                 = errors.New("log level is invalid, valid list is : \"DEBUG\", \"INFO\", \"WARN\", \"ERROR\", \"FATAL\"")
	errCloneDirectoryUnwritable = errors.New("clone directory is not writable")
	errConfigFileNotReadable    = errors.New("config file is not loadable")
)
