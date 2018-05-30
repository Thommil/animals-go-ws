package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/thommil/animals-go-common/model"

	"github.com/gin-gonic/gin"
)

// AuthenticationSettings defines configuration of Authentication
type AuthenticationSettings struct {
	URL string
}

// Authenticated middleware check authentication
func Authenticated(settings *AuthenticationSettings) gin.HandlerFunc {
	var httpClient = &http.Client{Timeout: 10 * time.Second}

	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		if strings.Contains(authorizationHeader, "Bearer") {
			tokenString := strings.TrimSpace(strings.Replace(authorizationHeader, "Bearer", "", -1))
			if response, err := httpClient.Get(strings.Replace(settings.URL, ":tokenString", tokenString, 1)); err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"message": err.Error(),
				})
			} else {
				defer response.Body.Close()
				user := &model.User{}
				json.NewDecoder(response.Body).Decode(user)
				c.Set("user", user)
				c.Next()
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Missing Bearer in headers",
			})
		}
	}
}
