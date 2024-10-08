package util

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func LogFailed(msg string, method string, waktu time.Time, httpstatus int, err error) {
	log.Info().
		Int("httpStatusCode", httpstatus).
		Str("StatusDescription", http.StatusText(httpstatus)).
		TimeDiff("ProcessTime", time.Now(), waktu).
		Str("httpMethod", method).
		Str("Error: ", err.Error()).
		Msg(msg)
}

func LogSuccess(msg string, method string, waktu time.Time, httpstatus int) {
	log.Info().
		Int("httpStatusCode", httpstatus).
		Str("StatusDescription", http.StatusText(httpstatus)).
		TimeDiff("ProcessTime", time.Now(), waktu).
		Str("httpMethod", method).
		Msg(msg)
}
