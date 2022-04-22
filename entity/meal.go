package entity

import "cooking-backend-go/common"

type Meal struct {
	Id     string `gorm:"primaryKey;column:id;type:varchar(32)"`
	UserId string `gorm:"column:user_id;type:varchar(32)"`
	Date   int64  `gorm:"column:date;type:bigint"`
	Type   int    `gorm:"column:type;type:int"`
}

type MealCourse struct {
	Id       string `gorm:"primaryKey;column:id;type:varchar(32)"`
	MealId   string `gorm:"column:meal_id;type:varchar(32)"`
	CourseId string `gorm:"column:course_id;type:varchar(32)"`
}

func (*Meal) TableName() string {
	return common.TableMeal
}

func (*MealCourse) TableName() string {
	return common.TableMealCourse
}
