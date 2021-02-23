package router

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/greeting-svc/log"
)

func LogRequest(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Logger().Info(fmt.Sprintf("method: %v, uri: %v, name: %v, response_time: %v",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start)))
	})
}
