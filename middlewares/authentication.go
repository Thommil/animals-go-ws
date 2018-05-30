package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	cache "github.com/patrickmn/go-cache"
	"github.com/thommil/animals-go-common/model"

	"github.com/gin-gonic/gin"
)

// AuthenticationSettings defines configuration of Authentication
type AuthenticationSettings struct {
	// URL for authentication check when not found in cache
	URL string
	// Expired defines the delay in seconds before cache expiration
	Expired time.Duration
}

// Authenticated middleware check authentication
func Authenticated(settings *AuthenticationSettings) gin.HandlerFunc {
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	var userCache = cache.New(settings.Expired*time.Second, settings.Expired*time.Second)

	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		// First check presence of Authorization header
		if strings.Contains(authorizationHeader, "Bearer") {
			tokenString := strings.TrimSpace(strings.Replace(authorizationHeader, "Bearer", "", -1))
			if user, ok := userCache.Get(tokenString); ok {
				// Try to find user in local cache
				c.Set("user", user)
				c.Next()
			} else {
				// Try to get user from authentication service
				if response, err := httpClient.Get(strings.Replace(settings.URL, ":tokenString", tokenString, 1)); err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"code":    http.StatusUnauthorized,
						"message": err.Error(),
					})
				} else {
					defer response.Body.Close()
					user = &model.User{}
					json.NewDecoder(response.Body).Decode(user)
					c.Set("user", user)
					userCache.Set(tokenString, user, cache.DefaultExpiration)
					c.Next()
				}
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Missing Bearer in headers",
			})
		}
	}
}
