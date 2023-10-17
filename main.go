package main

import (
	"io"
	"net/http"
	"os"
)

type flushingWriter struct {
	io.Writer
}

func (w *flushingWriter) Write(p []byte) (int, error) {
	n, err := w.Writer.Write(p)
	w.Writer.(http.Flusher).Flush()
	return n, err
}

func handler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		w.Header().Set("Content-Type", "application/octet-stream")
	} else {
		w.Header().Set("Content-Type", contentType)
	}
	w.WriteHeader(http.StatusOK)
	io.Copy(&flushingWriter{Writer: w}, r.Body)
}

func main() {
	http.HandleFunc("/", handler)
	listenAddr := ":8080"
	if len(os.Args) > 1 {
		listenAddr = os.Args[1]
	}
	http.ListenAndServe(listenAddr, nil)
}
