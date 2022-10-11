package middleware

import "github.com/gin-gonic/gin"

func BasihAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"pragmatic": "reviews",
	})
}
