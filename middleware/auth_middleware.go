package middleware

import (
	"cooking-backend-go/response"
	"cooking-backend-go/utils/jwtutils"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := ctx.Request

		//登录放行
		if request.URL.Path == "/login" {
			ctx.Next()
			return
		}

		token := request.Header.Get("Authorization")
		token = strings.ReplaceAll(token, "Bearer ", "")

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

		if int64(claims["timestamp"].(float64)) > time.Now().UnixMilli()+1000*60*60*24 {
			response.Error(ctx, response.ResultTokenExpired)
			ctx.Abort()
			return
		}

		request.Header.Set("userId", claims["userId"].(string))
	}
}
