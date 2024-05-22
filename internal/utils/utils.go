package utils

import (
	"github.com/sirupsen/logrus"
	"regexp"
)

func IsHttpRepo(url string) bool {
	regex := "^http.?://.*"
	result, err := regexp.Match(regex, []byte(url))
	if err != nil {
		logrus.Fatal(err)
	}
	if result {
		return true
	}

	return false
}
