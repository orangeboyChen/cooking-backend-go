package impl

import (
	"cooking-backend-go/common"
	"cooking-backend-go/entity"
	"cooking-backend-go/utils"
)

type IngredientCourseDaoImpl struct{}

func (*IngredientCourseDaoImpl) FindIngredientCourseByCourseIdList(courseIdList []string) ([]*entity.IngredientCourse, error) {
	var result []entity.IngredientCourse
	err := common.DB.Model(&entity.IngredientCourse{}).Where("course_id in (?)", courseIdList).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return utils.ToPointerList(result), nil
}

func (*IngredientCourseDaoImpl) FindIngredientCourseByCourseId(courseId string) ([]*entity.IngredientCourse, error) {
	var result []entity.IngredientCourse
	err := common.DB.Model(&entity.IngredientCourse{}).Where(&entity.IngredientCourse{CourseId: courseId}).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return utils.ToPointerList(result), nil
}
