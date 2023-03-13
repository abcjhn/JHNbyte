package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func AuthMiddleWare() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		tokenstring := ctx.Query("token")
		token,claims,err := ParseToken(tokenstring);
		if err!=nil||!token.Valid{
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"msg":"权限不足",
			})
			ctx.Abort()
			return
		}
		ctx.Set("uid",claims.userid)
		ctx.Next()
		
	}
}