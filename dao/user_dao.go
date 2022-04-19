package dao

import (
	"cooking-backend-go/common"
	"cooking-backend-go/entity"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

type UserDao struct {
}

var UserDaoInstance = UserDao{}

func (*UserDao) InsertUser(user *entity.User) {
	user.Id = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	user.CreateTime = time.Now().UnixMilli()
	common.DB.Table(common.TableUser).Create(user)
}

func (s *UserDao) UpdateUser(user *entity.User) {
	common.DB.Table(common.TableUser).Select("id", user.Id).Updates(user)
}

func (*UserDao) FindUserById(id string) (*entity.User, error) {
	var user entity.User
	if err := common.DB.Table(common.TableUser).Find(&user, id).Error; err != nil {
		return nil, err
	}

	if user.Id == "" {
		return nil, nil
	}

	return &user, nil
}

func (*UserDao) FindUserByUserIdList(idList []string) ([]*entity.User, error) {
	var userList []entity.User
	if err := common.DB.Table(common.TableUser).Select("id in (?)", idList).Pluck("avatar", &userList).Error; err != nil {
		return nil, err
	}

	var result = make([]*entity.User, len(userList))
	for i := range userList {
		result[i] = &userList[i]
	}
	return result, nil
}

func (*UserDao) FindUserByOpenid(openid string) (*entity.User, error) {
	var user entity.User
	if err := common.DB.Table(common.TableUser).Select("openid = ?", openid).Limit(1).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
