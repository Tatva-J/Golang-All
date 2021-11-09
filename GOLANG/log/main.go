package main

import log "github.com/sirupsen/logrus"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"Day":  "Foo",
			"Time": "bar",
		},
	).Info("this is json message")
	log.Info("this is logging")
	log.Info("this is another logging")
	log.Error("this is error")
}
