package utils

import (
	"fmt"
	"net/http"
	"time"
)

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
	// bytesReceived int
	// bytesSent     int
}

func newLogResponseWriter(w http.ResponseWriter) *logResponseWriter {
	return &logResponseWriter{ResponseWriter: w}
}

func (w *logResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

//////////////////////////////////
//////////////////////////////////

type loggingMiddleware struct {
	handler http.Handler
}

func (s *loggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	logRespWriter := newLogResponseWriter(w)

	s.handler.ServeHTTP(logRespWriter, r)

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf(
		"duration=%v status=%d path=%s\n",
		elapsed,
		logRespWriter.statusCode,
		r.URL.Path)
}

func LogHttp(handlerToWrap http.Handler) *loggingMiddleware {
	return &loggingMiddleware{
		handler: handlerToWrap,
	}
}
