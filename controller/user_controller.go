package controller

import (
	"cooking-backend-go/dto"
	"cooking-backend-go/response"
	"cooking-backend-go/service"
	"cooking-backend-go/utils/jwtutils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	uuid "github.com/satori/go.uuid"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type UserController struct {
}

var UserControllerInstance = UserController{}

// Login 登录
// @Tags     鉴权
// @Summary  登录
// @Param    body  body      dto.UserLoginDto  true  "登录数据"
// @Success  200   {object}  response.Result
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
		response.ErrorHandler(ctx, err)
		return
	}

	//通过userId颁发jwt
	jwt, _ := jwtutils.CreateJwtToken(userId)

	//返回成功
	response.SuccessData(ctx, &gin.H{
		"token": "Bearer " + jwt,
	})
}

// UpdateUserInfo 更新用户信息
// @Tags      用户
// @Summary   更新用户信息
// @Security  ApiAuthToken
// @Param     body  body  dto.UserInfoDto  true  "body"
// @Router    /user/info [PUT]
func (*UserController) UpdateUserInfo(ctx *gin.Context) {
	request := ctx.Request

	var userInfoDto dto.UserInfoDto
	body, err := ioutil.ReadAll(request.Body)
	err = json.Unmarshal(body, &userInfoDto)
	if err != nil {
		response.Error(ctx, response.ResultPatternError)
		return
	}

	userId := request.Header.Get("userId")
	err = service.UserService.UpdateUserInfo(userInfoDto, userId)
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}

	response.Success(ctx)
}

// GetAvatar 获取用户头像
// @Tags      用户
// @Summary   获取用户头像
// @Security  ApiAuthToken
// @Param     avatarFileName  path  string  true  "avatarFileName"
// @Router    /user/avatar/{avatarFileName} [GET]
func (*UserController) GetAvatar(ctx *gin.Context) {
	avatarFileName := ctx.Param("avatarFileName") + ".jpg"

	open, err := os.Open("./data/avatar/" + avatarFileName)
	defer open.Close()

	if err != nil {
		response.Error(ctx, response.ResultNoSuchFile)
		return
	}
	ctx.File("./data/avatar/" + avatarFileName)
}

// UploadAvatar 上传头像
// @Tags      用户
// @Summary   上传头像
// @Param     avatar  formData  file  true  "头像"
// @Security  ApiAuthToken
// @Router    /user/avatar [PUT]
func (*UserController) UploadAvatar(ctx *gin.Context) {
	request := ctx.Request
	userId := request.Header.Get("userId")
	avatarFile, err := ctx.FormFile("avatar")

	//判断文件名
	fileName := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	ext := strings.ToLower(path.Ext(avatarFile.Filename))

	//准备路径
	os.MkdirAll("./data/temp/", 0777)
	os.MkdirAll("./data/avatar/", 0777)

	//缓存原始图片
	err = ctx.SaveUploadedFile(avatarFile, "./data/temp/"+fileName+ext)
	defer func() {
		os.Remove("./data/temp/" + fileName + ext)
	}()
	if err != nil {
		response.Error(ctx, response.ResultInternalServerError)
		panic(err)
	}
	file, err := os.Open("./data/temp/" + fileName + ext)

	//压缩图片
	var img image.Image
	if ext == ".png" {
		img, err = png.Decode(file)
	} else if ext == ".jpg" || ext == ".jpeg" {
		img, err = jpeg.Decode(file)
	}

	if err != nil {
		response.Error(ctx, response.ResultInternalServerError)
		panic(err)
	}

	compressedImage := resize.Resize(512, 512, img, resize.NearestNeighbor)
	out, err := os.Create("./data/avatar/" + fileName + ".jpg")
	defer out.Close()

	jpeg.Encode(out, compressedImage, &jpeg.Options{Quality: jpeg.DefaultQuality})

	//保存进数据库
	err = service.UserService.SetAvatar(userId, fileName+".jpg")
	if err != nil {
		response.ErrorHandler(ctx, err)
		os.Remove("./data/avatar/" + fileName + ".jpg")
		return
	}
	response.Success(ctx)
}

// GetUserInfo 获取用户信息
// @Tags      用户
// @Summary   获取用户信息
// @Security  ApiAuthToken
// @Router    /user [GET]
func (*UserController) GetUserInfo(ctx *gin.Context) {
	request := ctx.Request
	userId := request.Header.Get("userId")

	userVO, err := service.UserService.FindUserById(userId)
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}

	if userVO == nil {
		response.Error(ctx, response.ResultNoSuchUser)
		return
	}

	response.SuccessData(ctx, userVO)
}
