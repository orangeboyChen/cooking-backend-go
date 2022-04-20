package controller

import (
	"cooking-backend-go/dto"
	"cooking-backend-go/response"
	"cooking-backend-go/service"
	"cooking-backend-go/utils/jwtutils"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	request := ctx.Request

	bodyStream := request.Body
	var bodyByte = make([]byte, 1024)
	_, _ = bodyStream.Read(bodyByte)

	var loginDto dto.UserLoginDto
	if err := json.Unmarshal(bodyByte, &loginDto); err != nil {
		response.Error(ctx, response.ResultPatternError)
		return
	}

	//开始保存user信息
	userId, err := service.UserService.Login(loginDto)
	if err != nil {
		response.Error(ctx, response.ResultPatternError)
		return
	}

	//通过userId颁发jwt
	jwt, _ := jwtutils.CreateJwtToken(userId)

	//返回成功
	response.SuccessData(ctx, &gin.H{
		"token": "Bearer " + jwt,
	})
}
