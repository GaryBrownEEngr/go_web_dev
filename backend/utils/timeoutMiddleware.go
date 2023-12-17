package utils

import (
	"net/http"
	"time"
)

// type timeoutMiddleware struct {
// 	handler http.Handler
// }

// func (s *timeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	s.handler.ServeHTTP(w, r)
// }

// func TimeoutMiddleware(timeout time.Duration, next http.Handler) *timeoutMiddleware {
// 	next = http.TimeoutHandler(next, timeout, "HTTP Request timed out")
// 	return &timeoutMiddleware{
// 		handler: next,
// 	}
// }

func TimeoutMiddleware(timeout time.Duration, next http.Handler) http.Handler {
	next = http.TimeoutHandler(next, timeout, "HTTP Request timed out")
	return next
}
