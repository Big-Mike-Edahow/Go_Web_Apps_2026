// logger.go

package main

import (
	"log"
	"net/http"
)

type LoggerResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *LoggerResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func logRequest(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		lrw := &LoggerResponseWriter{w, http.StatusOK}
		wrappedHandler.ServeHTTP(lrw, req)
		statusCode := lrw.statusCode

		log.Printf("--> %s \"%s\" %d - %s", req.Method, req.URL.Path, statusCode, http.StatusText(statusCode))
	})
}

