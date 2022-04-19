package entity

type User struct {
	Id         string `gorm:"column:id;primaryKey"`
	Nickname   string `gorm:"column:name"`
	Openid     string `gorm:"column:openid"`
	Avatar     string `gorm:"column:avatar"`
	CreateTime int64  `gorm:"column:create_time"`
}
