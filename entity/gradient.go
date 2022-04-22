package entity

import "cooking-backend-go/common"

type Ingredient struct {
	Id           string `gorm:"primaryKey;column:id;type:varchar(32)"`
	Name         string `gorm:"column:name;type:varchar(32)"`
	Image        string `gorm:"column:image;type:varchar(32)"`
	Introduction string `gorm:"column:introduction;type:varchar(32)"`
}

type IngredientCourse struct {
	Id           string `gorm:"primaryKey;column:id;type:varchar(32)"`
	IngredientId string `gorm:"column:ingredient_id;type:varchar(32)"`
	CourseId     string `gorm:"column:course_id;type:varchar(32)"`
	Quantity     int    `gorm:"column:quantity;type:int"`
	Unit         string `gorm:"column:unit;type:varchar(32)"`
}

func (*Ingredient) TableName() string {
	return common.TableIngredient
}

func (*IngredientCourse) TableName() string {
	return common.TableIngredientCourse
}
