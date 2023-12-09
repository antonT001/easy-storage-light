package middlewares

import (
	"net/http"

	"github.com/antonT001/easy-storage-light/internal/lib/httplib"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(httplib.AccessControlAllowOrigin, "*")

		if r.Method == http.MethodOptions {
			w.Header().Set(httplib.AccessControlAllowHeaders, "*")
			w.Header().Set(httplib.AccessControlAllowMethod, http.MethodPost)
			return
		}

		next.ServeHTTP(w, r)
	})
}
