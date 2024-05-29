package Log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func Init() {
	file, err := os.OpenFile("Log/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint: true,
	})
}
