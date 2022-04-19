package middleware

import (
	"cooking-backend-go/response"
	"cooking-backend-go/utils/jwtutils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware /**
/**
 * @Author orangeboyChen
 * @Description
 * @Date 2022/4/19 16:17
 **/
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.Request

		//登录放行
		if request.URL.Path == "/login" {
			ctx.Next()
			return
		}

		token := request.Header.Get("Authorization")

		if token == "" {
			response.Error(ctx, response.ResultPermissionDenied)
			ctx.Abort()
			return
		}

		claims, err := jwtutils.DecodeJwtToken(token)
		if err != nil {
			response.Error(ctx, response.ResultPermissionDenied)
			ctx.Abort()
			return
		}

		request.Header.Set("userId", claims["userId"].(string))
		ctx.Next()
	}
}
