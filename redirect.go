package requestlogger

import (
	"github.com/sirupsen/logrus"
	"log"
)

func RedirectStdlogOutput(logger *logrus.Logger) {
	log.SetOutput(&redirectedWriter{Logger: logger})
	log.SetFlags(0)
}

type redirectedWriter struct {
	Logger *logrus.Logger
}

func (l *redirectedWriter) Write(p []byte) (n int, err error) {
	if len(p) > 0 {
		l.Logger.Infof("%s", p[:len(p)-1])
	}

	return len(p), nil
}
