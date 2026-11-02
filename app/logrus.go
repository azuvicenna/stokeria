package app

import (
	"github.com/sirupsen/logrus"
)

func NewLogrus() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05", // Format waktu standar Asia/Jakarta
	})

	log.SetLevel(logrus.InfoLevel)

	return log
}