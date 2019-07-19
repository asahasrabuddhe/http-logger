package requestlogger

import (
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type requestLogger struct {
	Logger *logrus.Logger
}

func (h *requestLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &requestLoggerEntry{Logger: logrus.NewEntry(h.Logger)}
	logFields := logrus.Fields{}

	logFields["ts"] = time.Now().UTC().Format(time.RFC1123)

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["request_id"] = reqID
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	host := r.Host

	logFields["scheme"] = scheme
	logFields["protocol"] = r.Proto
	logFields["method"] = r.Method

	logFields["remote_address"] = r.RemoteAddr
	logFields["user_agent"] = r.UserAgent()

	if val := r.Header.Get("X-Forwarded-For"); val != "" {
		logFields["X-Forwarded-For"] = val
	}
	if val := r.Header.Get("X-Forwarded-Host"); val != "" {
		logFields["X-Forwarded-Host"] = val
		host = val
	}
	if val := r.Header.Get("X-Forwarded-Scheme"); val != "" {
		logFields["X-Forwarded-Scheme"] = val
		scheme = val
	}

	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, host, r.RequestURI)

	entry.Logger = entry.Logger.WithFields(logFields)

	entry.Logger.Infoln("request started")

	return entry
}

