package infra

import (
	"net"
	"net/http"

	uc "github.com/RodrigoMath/rateLimiter/internal/usecase"
)

// Agora o middleware é uma função que "fabrica" o middleware real,
// injetando o UseCase nele.
func NewRateLimiterMiddleware(uc *uc.RateLimiterUseCase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token := r.Header.Get("API_KEY")
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)

			identifier := ip
			isToken := false
			if token != "" {
				identifier = token
				isToken = true
			}

			// Chamando o UseCase!
			allowed, err := uc.Execute(r.Context(), identifier, isToken)
			if err != nil || !allowed {
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("you have reached the maximum number of requests or actions allowed within a certain time frame"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
