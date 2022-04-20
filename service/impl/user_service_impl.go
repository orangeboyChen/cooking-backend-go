package impl

import (
	"cooking-backend-go/dao"
	"cooking-backend-go/dto"
	"cooking-backend-go/entity"
	"cooking-backend-go/response"
	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserServiceImpl struct {
}

func (*UserServiceImpl) Login(dto dto.UserLoginDto) (string, error) {
	//1. 向苹果要签名
	jwks, err := keyfunc.Get("https://appleid.apple.com/auth/keys", keyfunc.Options{})
	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(dto.IdentityToken, jwks.Keyfunc)
	if err != nil {
		return "", &response.AppException{Code: response.ResultPermissionDenied}
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims["iss"].(string) != "https://appleid.apple.com" || claims["exp"].(int64) > time.Now().Unix() {
		return "", &response.AppException{Code: response.ResultPermissionDenied}
	}

	openid := claims["sub"].(string)
	user, err := dao.UserDao.FindUserByOpenid(openid)
	if err != nil {
		return "", err
	}

	if user == nil {
		dao.UserDao.InsertUser(&entity.User{
			Nickname: claims["email"].(string),
			Openid:   openid,
			Avatar:   "",
		})
	} else {
		//更改用户，但目前还没有处理逻辑
	}

	return user.Id, nil
}
