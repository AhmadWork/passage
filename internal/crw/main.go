package crw
import (
    "net/http"
    "bytes"
)
type LoggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Body       bytes.Buffer
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	// Default the status code to 200 in case WriteHeader is not called
	return &LoggingResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *LoggingResponseWriter) Write(b []byte) (int, error) {
	lrw.Body.Write(b)
	return lrw.ResponseWriter.Write(b)
}

