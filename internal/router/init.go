package router

import (
	"net/http"

	"git.gnous.eu/ada/spiegel/internal/constant"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Listen  string
	Archive string
}

func (c Config) Router() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(c.Archive))
	mux.Handle("/git/", http.StripPrefix("/git/", fs))

	server := &http.Server{
		Addr:              c.Listen,
		Handler:           mux,
		ReadHeaderTimeout: constant.HTTPTimeout,
	}

	logrus.Info("HTTP listen on :5000")

	logrus.Fatal(server.ListenAndServe())
}
