package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/theabdullahishola/to-do/util"
)

func Authenticate(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing token"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token format"})
		return
	}
	userID, err:=util.VerifyAccessToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"unauthorized"})
		return 
	}
	c.Set("userID", userID)
	c.Next()
}