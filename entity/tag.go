package entity

type Tag struct {
	Id        string `gorm:"column:id;primaryKey"`
	Name      string `gorm:"column:name"`
	TagTypeId string `gorm:"column:tag_type_id"`
}

type SearchTag struct {
	Id        string
	Name      string
	TagTypeId string
}

type TagType struct {
	Id   string `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}
