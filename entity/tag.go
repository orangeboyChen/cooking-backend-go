package entity

import "cooking-backend-go/common"

type Tag struct {
	Id        string `gorm:"column:id;primaryKey;type:varchar(32)"`
	Name      string `gorm:"column:name;type:varchar(32)"`
	TagTypeId string `gorm:"column:tag_type_id;type:varchar(32)"`
}

type SearchTag struct {
	Id        string
	Name      string
	TagTypeId string
}

type TagType struct {
	Id   string `gorm:"column:id;primaryKey;type:varchar(32)"`
	Name string `gorm:"column:name;type:varchar(32)"`
}

func (*Tag) TableName() string {
	return common.TableTag
}

func (*TagType) TableName() string {
	return common.TableTagType
}
