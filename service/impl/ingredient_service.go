package impl

import (
	"cooking-backend-go/dao"
	"cooking-backend-go/entity"
	"cooking-backend-go/vo"
)

type IngredientServiceImpl struct{}

func (*IngredientServiceImpl) ConveyModelToVo(model *entity.Ingredient) *vo.IngredientVO {
	return &vo.IngredientVO{
		Id:          model.Id,
		Name:        model.Name,
		Image:       model.Image,
		Description: model.Description,
	}
}

func (*IngredientServiceImpl) GetIngredientByCourseId(courseId string) ([]*entity.Ingredient, error) {
	ingredientResult, err := dao.IngredientCourseDao.FindIngredientCourseByCourseId(courseId)
	if err != nil {
		return nil, err
	}

	//转化为id列表
	var idList = make([]string, len(ingredientResult))
	for i, ingredientCourse := range ingredientResult {
		idList[i] = ingredientCourse.IngredientId
	}

	//通过id列表查询
	list, err := dao.IngredientDao.FindIngredientByIdList(idList)
	if err != nil {
		return nil, err
	}

	return list, nil
}
