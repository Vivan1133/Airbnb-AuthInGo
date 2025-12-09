package middlewares

import (
	"AuthInGo/utils"
	"net"
	"net/http"
)

// var limiter = rate.NewLimiter(5, 5)	// 5 req per second
// 5 token per second refill in the bucket with a burst of 10 

func RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

		ip, _ , err := net.SplitHostPort(r.RemoteAddr)

		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, "error while getting the ip", err)
			return 
		}

		limiter := utils.GetVisitor(ip)

		if !limiter.Allow() {
			http.Error(w, "too many req please try again later", http.StatusTooManyRequests)
			return 
		}
		
		next.ServeHTTP(w, r)
	})
}