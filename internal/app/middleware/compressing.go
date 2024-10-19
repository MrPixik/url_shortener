package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func CompressingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encodingHeader := r.Header.Get("Accept-Encoding")
		encodings := strings.Split(encodingHeader, ",")

		for _, encoding := range encodings {
			if strings.TrimSpace(encoding) == "gzip" {
				w.Header().Set("Content-Encoding", "gzip")
				gz, err := gzip.NewWriterLevel(w, gzip.BestCompression)
				if err != nil {
					panic(err)
				}
				defer gz.Close()
				next.ServeHTTP(gzipResponseWriter{ResponseWriter: w, Writer: gz}, r)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
