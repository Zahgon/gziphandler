//go:build go1.8
// +build go1.8

package gziphandler

import "net/http"

// Push initiates an HTTP/2 server push.
// Push returns ErrNotSupported if the client has disabled push or if push
// is not supported on the underlying connection.
func (w *GzipResponseWriter) Push(target string, opts *http.PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// setAcceptEncodingForPushOptions sets "Accept-Encoding" : "gzip" for PushOptions without overriding existing headers.
func setAcceptEncodingForPushOptions(opts *http.PushOptions) *http.PushOptions {
	_ = "STUB: not implemented"
	return nil
}
