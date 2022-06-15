package impl

import (
	"cooking-backend-go/common"
	"cooking-backend-go/entity"
	"cooking-backend-go/utils"
)

type IngredientDaoImpl struct{}

func (*IngredientDaoImpl) InsertIngredient(ingredient *entity.Ingredient) error {
	return common.DB.Model(&entity.Ingredient{}).Create(ingredient).Error
}

func (*IngredientDaoImpl) FindIngredientByIdList(idList []string) ([]*entity.Ingredient, error) {
	var result []entity.Ingredient
	err := common.DB.Model(&entity.Ingredient{}).Where("id in (?)", idList).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return utils.ToPointerList(result), nil
}
