package utils

import (
	"regexp"

	"github.com/sirupsen/logrus"
)

func IsHTTPRepo(url string) bool {
	regex := "^http.?://.*"
	result, err := regexp.MatchString(regex, url)
	if err != nil {
		logrus.Fatal(err)
	}
	if result {
		return true
	}

	return false
}
