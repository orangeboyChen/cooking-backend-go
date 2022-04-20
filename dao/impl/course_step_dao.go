package impl

import (
	"cooking-backend-go/common"
	"cooking-backend-go/entity"
	"cooking-backend-go/utils"
)

type CourseStepDaoImpl struct{}

func (*CourseStepDaoImpl) FindCourseStepByCourseId(courseId string) ([]*entity.CourseStep, error) {
	var courseStepList []entity.CourseStep
	if err := common.DB.Where("course_id = ?", courseId).Find(&courseStepList).Error; err != nil {
		return nil, err
	}

	var courseStepPointerList = make([]*entity.CourseStep, len(courseStepList))
	for i := range courseStepList {
		courseStepPointerList[i] = &courseStepList[i]
	}

	return courseStepPointerList, nil
}

func (*CourseStepDaoImpl) InsertList(list []*entity.CourseStep) error {
	structList := utils.ToStructList(list)
	return common.DB.Create(&structList).Error
}

func (*CourseStepDaoImpl) DeleteByCourseId(courseId string) error {
	return common.DB.Delete("course_id = ?", courseId).Error
}
