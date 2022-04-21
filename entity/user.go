package entity

import "cooking-backend-go/common"

type User struct {
	Id         string `gorm:"column:id;primaryKey;type:varchar(32)"`
	Nickname   string `gorm:"column:name;type:varchar(32)"`
	Openid     string `gorm:"column:openid;type:varchar(64)"`
	Avatar     string `gorm:"column:avatar;type:varchar(32)"`
	Gender     int    `gorm:"column:gender"`
	Birthday   int64  `gorm:"birthday"`
	CreateTime int64  `gorm:"column:create_time"`
}

func (*User) TableName() string {
	return common.TableUser
}
