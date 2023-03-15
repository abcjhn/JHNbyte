package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func AuthMiddleWare() gin.HandlerFunc  {
	return func(c *gin.Context) {
		tokenstring := c.Query("token")
		token,claims,err := ParseToken(tokenstring);
		if err!=nil||!token.Valid{
			c.JSON(http.StatusUnauthorized,gin.H{
				"status_code":1,
				"status_msg":"权限不足",
			})
			c.Abort()
			return
		}
		c.Set("uid",claims.userid)
		c.Next()
		
	}
}