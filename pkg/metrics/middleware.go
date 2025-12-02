package metrics

import (
	"net/http"
	"strconv"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode   int
	bytesWritten int64
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.bytesWritten += int64(n)
	return n, err
}

func HTTPMiddleware(base *BaseMetrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			base.HTTPActiveConnections.Inc()
			defer base.HTTPActiveConnections.Dec()

			rw := &responseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			if r.ContentLength > 0 {
				base.HTTPRequestSize.WithLabelValues(r.Method, r.URL.Path).Observe(float64(r.ContentLength))
			}

			start := time.Now()

			next.ServeHTTP(rw, r)
			duration := time.Since(start).Seconds()

			status := strconv.Itoa(rw.statusCode)
			base.HTTPRequestsTotal.WithLabelValues(r.Method, r.URL.Path, status).Inc()
			base.HTTPRequestDuration.WithLabelValues(r.Method, r.URL.Path, status).Observe(duration)
			base.HTTPResponseSize.WithLabelValues(r.Method, r.URL.Path).Observe(float64(rw.bytesWritten))
		})
	}
}
