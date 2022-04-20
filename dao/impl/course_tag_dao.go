package impl

import (
	"cooking-backend-go/common"
	"cooking-backend-go/entity"
	uuid "github.com/satori/go.uuid"
	"strings"
)

type CourseTagDaoImpl struct {
}

func (*CourseTagDaoImpl) InsertCourseTagList(list []*entity.CourseTag) error {
	structList := make([]entity.CourseTag, len(list))
	for i, item := range list {
		item.Id = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
		structList[i] = *item
	}

	return common.DB.Create(&structList).Error
}

func (*CourseTagDaoImpl) DeleteCourseTagByCourseId(courseId string) error {
	return common.DB.Delete("course_id = ?", courseId).Error
}
