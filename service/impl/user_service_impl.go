package impl

import (
	"cooking-backend-go/dao"
	"cooking-backend-go/dto"
	"cooking-backend-go/entity"
	"cooking-backend-go/response"
	"cooking-backend-go/vo"
	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserServiceImpl struct {
}

func (*UserServiceImpl) FindUserById(userId string) (*vo.UserInfoVO, error) {
	user, err := dao.UserDao.FindUserById(userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return &vo.UserInfoVO{
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Birthday: user.Birthday,
	}, nil
}

func (*UserServiceImpl) UpdateUserInfo(userInfoDto dto.UserInfoDto, userId string) error {
	user, err := dao.UserDao.FindUserById(userId)
	if err != nil {
		return err
	}

	if user == nil {
		return &response.AppException{Code: response.ResultNoSuchUser}
	}

	user.Nickname = userInfoDto.NickName
	user.Birthday = userInfoDto.Birthday
	user.Gender = userInfoDto.Gender

	return dao.UserDao.UpdateUser(user)
}

func (*UserServiceImpl) GetAvatar(userId string) (string, error) {
	user, err := dao.UserDao.FindUserById(userId)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", &response.AppException{Code: response.ResultNoSuchUser}
	}

	return user.Avatar, nil
}

func (*UserServiceImpl) SetAvatar(userId string, avatarFilePath string) error {
	user, err := dao.UserDao.FindUserById(userId)
	if err != nil {
		return err
	}

	if user == nil {
		return &response.AppException{Code: response.ResultNoSuchUser}
	}

	user.Avatar = avatarFilePath
	dao.UserDao.UpdateUser(user)
	return nil
}

func (*UserServiceImpl) Login(dto dto.UserLoginDto) (string, error) {
	//1. 向苹果要签名
	jwks, err := keyfunc.Get("https://appleid.apple.com/auth/keys", keyfunc.Options{})
	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(dto.IdentityToken, jwks.Keyfunc)
	if err != nil {
		return "", &response.AppException{Code: response.ResultLoginError}
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims["iss"].(string) != "https://appleid.apple.com" || int64(claims["exp"].(float64)) < time.Now().Unix() {
		return "", &response.AppException{Code: response.ResultLoginError}
	}

	openid := claims["sub"].(string)
	user, err := dao.UserDao.FindUserByOpenid(openid)
	if err != nil {
		return "", err
	}

	nickname, ok := claims["email"].(string)
	if !ok {
		nickname = openid
	}

	if user == nil {
		user = &entity.User{
			Nickname: nickname,
			Openid:   openid,
			Avatar:   "",
		}
		dao.UserDao.InsertUser(user)
	} else {
		//更改用户，但目前还没有处理逻辑
	}

	return user.Id, nil
}
