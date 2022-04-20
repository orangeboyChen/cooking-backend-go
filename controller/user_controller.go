package controller

import (
	"cooking-backend-go/dto"
	"cooking-backend-go/response"
	"cooking-backend-go/service"
	"cooking-backend-go/utils/jwtutils"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type UserController struct {
}

var UserControllerImpl = UserController{}

// Login 登录
// @Tags     用户
// @Summary  登录
// @Param    dto  body      dto.UserLoginDto  true  "dto"
// @Success  200  {object}  response.Result
// @Router   /login [POST]
func (*UserController) Login(ctx *gin.Context) {
	request := ctx.Request

	bodyByte, err := ioutil.ReadAll(request.Body)

	var loginDto dto.UserLoginDto
	if err := json.Unmarshal(bodyByte, &loginDto); err != nil {
		response.Error(ctx, response.ResultPatternError)
		return
	}

	//开始保存user信息
	userId, err := service.UserService.Login(loginDto)
	if err != nil {
		var exception *response.AppException
		if errors.As(err, &exception) {
			response.ErrorException(ctx, *exception)
		} else {
			response.Error(ctx, response.ResultPatternError)
		}
		return
	}

	//通过userId颁发jwt
	jwt, _ := jwtutils.CreateJwtToken(userId)

	//返回成功
	response.SuccessData(ctx, &gin.H{
		"token": "Bearer " + jwt,
	})
}
