package service

import (
	"cooking-backend-go/dao"
	"cooking-backend-go/entity"
	"cooking-backend-go/vo"
)

type CourseServiceInterface interface {
	SearchCourse(keyword string, pageNum int, pageSize int)
	GetCourseByTag(tagId string, pageNum int, pageSize int)
}

type CourseService struct {
}

var CourseServiceInstance = CourseService{}

func (*CourseService) SearchCourse(keyword string, pageNum int, pageSize int) (*vo.PageVO[vo.SearchCourseVO], error) {
	page, err := dao.CourseDaoInstance.SearchCourse(keyword, pageNum, pageSize)
	if err != nil {
		return nil, err
	}

	var pageVO *vo.PageVO[vo.SearchCourseVO]
	pageVO = vo.ConveyPageToPageVO(page, searchCourseModelToVo)

	return pageVO, nil
}

func (*CourseService) GetCourseByTag(tagId string, pageNum int, pageSize int) (*vo.PageVO[vo.CourseVO], error) {
	page, err := dao.CourseDaoInstance.FindCourseByTagId(tagId, pageNum, pageSize)
	if err != nil {
		return nil, err
	}

	size := len(page.Data)
	var userIdAvatarMap = make(map[string]string, size)
	var userIdList = make([]string, size)
	for i := range page.Data {
		userIdList[i] = page.Data[i].UserId
	}

	userList, err := dao.UserDaoInstance.FindUserByUserIdList(userIdList)
	if err != nil {
		return nil, err
	}

	for i := range userList {
		userIdAvatarMap[userList[i].Id] = userList[i].Avatar
	}

	result := vo.ConveyPageToPageVO(page, func(t *entity.Course) *vo.CourseVO {
		return courseModelToVo(t, userIdAvatarMap[t.UserId])
	})

	return result, nil
}

func GetCourseDetail(courseId string) (*vo.CourseDetailVO, error) {
	course, err := dao.CourseDaoInstance.FindCourseById(courseId)
	if err != nil {
		return nil, err
	}

	if course == nil {
		return nil, nil
	}

	courseStepList, err := dao.CourseDaoInstance.FindCourseStepByCourseId(courseId)
	if err != nil {
		return nil, err
	}

	user, err := dao.UserDaoInstance.FindUserById(course.UserId)
	if err != nil {
		return nil, err
	}

	var courseStepVOList = make([]vo.CourseStepVO, len(courseStepList))

	for i := range courseStepList {
		courseStep := courseStepList[i]
		courseStepVOList[i] = vo.CourseStepVO{
			Order:   courseStep.Order,
			Content: courseStep.Content,
			Second:  courseStep.Second,
		}
	}

	return &vo.CourseDetailVO{
		Id:         course.Id,
		Name:       course.Name,
		Detail:     course.Detail,
		Image:      course.Image,
		UserId:     course.UserId,
		UserAvatar: user.Avatar,
		Step:       courseStepVOList,
		CreateTime: course.CreateTime,
	}, nil
}

func searchCourseModelToVo(entity *entity.SearchCourse) *vo.SearchCourseVO {
	return &vo.SearchCourseVO{
		Id:          entity.Id,
		Name:        entity.Name,
		Detail:      entity.Detail,
		Image:       entity.Image,
		UserId:      entity.UserId,
		UserAvatar:  entity.UserAvatar,
		NameWithHit: entity.NameWithHit,
		CreateTime:  entity.CreateTime,
	}
}

func courseModelToVo(entity *entity.Course, userAvatar string) *vo.CourseVO {
	return &vo.CourseVO{
		Id:         entity.Id,
		Name:       entity.Name,
		Detail:     entity.Detail,
		Image:      entity.Image,
		UserId:     entity.UserId,
		UserAvatar: userAvatar,
		CreateTime: entity.CreateTime,
	}
}
