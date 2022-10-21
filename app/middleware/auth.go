package middleware

import (
	"final_assignment/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := utils.TokenValid(c)
		if userId == 0 {
			c.String(http.StatusUnauthorized, "You Not Authorized")
			c.Abort()
			return
		}
		c.Set("UserID", userId)
		c.Next()
	}
}
