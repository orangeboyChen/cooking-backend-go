package impl

import (
	"cooking-backend-go/dao"
	"cooking-backend-go/dto"
	"cooking-backend-go/response"
)

type UserServiceImpl struct {
}

func (*UserServiceImpl) Login(dto dto.UserLoginDto) (string, error) {
	if dto.Openid == "" {
		return "", &response.AppException{Code: response.ResultPatternError}
	}

	userDao := dao.UserDao
	user, err := userDao.FindUserByOpenid(dto.Openid)
	if err != nil {
		userDao.InsertUser(user)
	} else {
		userDao.UpdateUser(user)
	}

	return user.Id, nil
}
