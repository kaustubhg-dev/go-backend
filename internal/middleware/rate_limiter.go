package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"go-backend/internal/utils"
)

type clientLimiter struct {
	limiter *rate.Limiter
}

var (
	clients = make(map[string]*clientLimiter)
	mu      sync.Mutex
)

func getClientLimiter(
	ip string,
	rps float64,
	burst int,
) *rate.Limiter {

	mu.Lock()
	defer mu.Unlock()

	if cl, ok := clients[ip]; ok {
		return cl.limiter
	}

	lim := rate.NewLimiter(
		rate.Limit(rps),
		burst,
	)

	clients[ip] = &clientLimiter{
		limiter: lim,
	}

	return lim
}

func RateLimiter(
	rps float64,
	burst int,
) gin.HandlerFunc {

	return func(c *gin.Context) {
		lim := getClientLimiter(
			c.ClientIP(),
			rps,
			burst,
		)

		if !lim.Allow() {
			utils.ErrorResponse(
				c,
				http.StatusTooManyRequests,
				"Rate limit exceeded",
			)
			c.Abort()
			return
		}

		c.Next()
	}
}