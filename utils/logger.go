package utils

import (
	"log"
	"os"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/logWriter"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type LogService struct {
}

var (
	logger *log.Logger
)

var Logger LogService

func (lo *LogService) InitLogger(app *newrelic.Application) {
	writer := logWriter.New(os.Stdout, app)
	logger = log.New(&writer, "", log.Default().Flags())
}

func (lo *LogService) Info(msg string) {
	logger.Println("[INFO]" + " " + msg)
}

func (lo *LogService) Debug(msg string) {
	logger.Println("[DEBUG]" + " " + msg)
}

func (lo *LogService) Error(msg string, err error) {
	logger.Println("[ERROR]" + " " + msg)
}
