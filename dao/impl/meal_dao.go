package impl

import (
	"cooking-backend-go/common"
	"cooking-backend-go/entity"
)

type MealDaoImpl struct{}

func (*MealDaoImpl) InsertMeal(meal *entity.Meal) error {
	return common.DB.Model(&entity.Meal{}).Create(meal).Error
}
