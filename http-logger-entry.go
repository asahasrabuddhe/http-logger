package requestlogger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type requestLoggerEntry struct {
	Logger logrus.FieldLogger
	Level  *logrus.Level
}

func (l *requestLoggerEntry) Write(status, bytes int, header http.Header,  elapsed time.Duration, extra interface{}) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"status":     status,
		"length":     bytes,
		"elapsed_ms": elapsed.String(),
	})

	if l.Level == nil {
		l.Logger.Infoln("request complete")
	} else {
		switch *l.Level {
		case logrus.DebugLevel:
			l.Logger.Debugln("request complete")
		case logrus.InfoLevel:
			l.Logger.Infoln("request complete")
		case logrus.WarnLevel:
			l.Logger.Warnln("request complete")
		case logrus.ErrorLevel:
			l.Logger.Errorln("request complete")
		case logrus.FatalLevel:
			l.Logger.Fatalln("request complete")
		case logrus.PanicLevel:
			l.Logger.Errorln("request complete")
		}
	}
}

func (l *requestLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})

	panicLevel := logrus.PanicLevel
	l.Level = &panicLevel
}
