package requestlogger

import (
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
)

func NewRequestLogger(logger *logrus.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&requestLogger{logger})
}

// PrintPanics is a development middleware that preempts the request logger
// and prints a panic message and stack trace to stdout.

func PrintPanics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				fmt.Printf("\nPANIC: %+v\n", rec)
				fmt.Printf("%s", debug.Stack())
				fmt.Printf("\nPANIC: %+v\n", rec)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
