package impl

import (
	"cooking-backend-go/common"
	"cooking-backend-go/entity"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type UserDaoImpl struct {
}

func (*UserDaoImpl) InsertUser(user *entity.User) error {
	user.Id = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	user.CreateTime = time.Now().UnixMilli()
	return common.DB.Create(user).Error
}

func (s *UserDaoImpl) UpdateUser(user *entity.User) error {
	return common.DB.Select("id", user.Id).Updates(user).Error
}

func (*UserDaoImpl) FindUserById(id string) (*entity.User, error) {
	var user entity.User
	if err := common.DB.Find(&user, id).Error; err != nil {
		return nil, err
	}

	if user.Id == "" {
		return nil, nil
	}

	return &user, nil
}

func (*UserDaoImpl) FindUserByUserIdList(idList []string) ([]*entity.User, error) {
	var userList []entity.User
	if err := common.DB.Where("id in (?)", idList).Pluck("avatar", &userList).Error; err != nil {
		return nil, err
	}

	var result = make([]*entity.User, len(userList))
	for i := range userList {
		result[i] = &userList[i]
	}
	return result, nil
}

func (*UserDaoImpl) FindUserByOpenid(openid string) (*entity.User, error) {
	var user entity.User
	if err := common.DB.Where("openid = ?", openid).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &user, nil
}
