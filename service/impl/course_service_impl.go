package impl

import (
	"cooking-backend-go/dao"
	"cooking-backend-go/dto"
	"cooking-backend-go/entity"
	"cooking-backend-go/response"
	"cooking-backend-go/vo"
)

type CourseServiceImpl struct{}

func (*CourseServiceImpl) SearchCourse(keyword string, pageNum int, pageSize int) (*vo.PageVO[vo.SearchCourseVO], error) {
	page, err := dao.CourseDao.SearchCourse(keyword, pageNum, pageSize)
	if err != nil {
		return nil, err
	}

	var pageVO *vo.PageVO[vo.SearchCourseVO]
	pageVO = vo.ConveyPageToPageVO(page, searchCourseModelToVo)

	return pageVO, nil
}

func (*CourseServiceImpl) GetCourseByTag(tagId string, pageNum int, pageSize int) (*vo.PageVO[vo.CourseVO], error) {
	page, err := dao.CourseDao.FindCourseByTagId(tagId, pageNum, pageSize)
	if err != nil {
		return nil, err
	}

	size := len(page.Data)
	var userIdAvatarMap = make(map[string]string, size)
	var userIdList = make([]string, size)
	for i := range page.Data {
		userIdList[i] = page.Data[i].UserId
	}

	userList, err := dao.UserDao.FindUserByUserIdList(userIdList)
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

func (*CourseServiceImpl) GetCourseDetail(courseId string) (*vo.CourseDetailVO, error) {
	course, err := dao.CourseDao.FindCourseById(courseId)
	if err != nil {
		return nil, err
	}

	if course == nil {
		return nil, &response.AppException{Code: response.ResultNotFound}
	}

	courseStepList, err := dao.CourseStepDao.FindCourseStepByCourseId(courseId)
	if err != nil {
		return nil, err
	}

	user, err := dao.UserDao.FindUserById(course.UserId)
	if err != nil {
		return nil, err
	}

	var courseStepVOList = make([]*vo.CourseStepVO, len(courseStepList))

	for i := range courseStepList {
		courseStep := courseStepList[i]
		courseStepVOList[i] = &vo.CourseStepVO{
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

func (*CourseServiceImpl) GetCourseRecommendation() ([]*vo.SearchCourseVO, error) {
	courseList, err := dao.CourseDao.GetRecommendationCourse()
	if err != nil {
		return nil, err
	}

	var courseVOList = make([]*vo.SearchCourseVO, len(courseList))
	for i := range courseList {
		vo := searchCourseModelToVo(courseList[i])
		courseVOList[i] = vo
	}

	return courseVOList, nil
}

func (*CourseServiceImpl) InsertCourse(courseDto dto.CourseDto, userId string) (string, error) {
	//查找用户信息
	user, err := dao.UserDao.FindUserById(userId)
	if err != nil {
		return "", err
	}

	//检验tag的合法性
	tagList, err := dao.TagDao.GetTagListByIdList(courseDto.Tags)
	if err != nil {
		return "", err
	}
	if len(tagList) != len(courseDto.Tags) {
		return "", &response.AppException{Code: response.ResultPatternError}
	}

	//组装Search Course
	var searchCourse = entity.SearchCourse{
		Name:       courseDto.Name,
		Detail:     courseDto.Detail,
		Image:      courseDto.Image,
		UserId:     userId,
		UserAvatar: user.Avatar,
	}

	//保存SearchCourse
	err = dao.CourseDao.InsertSearchCourse(&searchCourse)
	if err != nil {
		return "", nil
	}

	//组装Course
	var course = entity.Course{
		Id:     searchCourse.Id,
		Name:   courseDto.Name,
		Detail: courseDto.Detail,
		Image:  courseDto.Image,
		UserId: userId,
	}

	//保存Course
	dao.CourseDao.InsertCourse(&course)

	//组装CourseTag
	var courseTagList = make([]*entity.CourseTag, len(courseDto.Tags))
	for i, tagId := range courseDto.Tags {
		courseTagList[i] = &entity.CourseTag{
			CourseId: course.Id,
			TagId:    tagId,
		}
	}

	//保存CourseTag
	err = dao.CourseTagDao.InsertCourseTagList(courseTagList)
	if err != nil {
		return "", err
	}

	//组装Step
	var stepList = make([]*entity.CourseStep, len(courseDto.Step))
	for i, element := range courseDto.Step {
		step := &entity.CourseStep{
			CourseId: searchCourse.Id,
			Content:  element.Content,
			Order:    element.Order,
			Second:   element.Second,
		}

		stepList[i] = step
	}

	//保存Step
	err = dao.CourseStepDao.InsertList(stepList)
	if err != nil {
		return "", err
	}

	return userId, nil
}

func (*CourseServiceImpl) UpdateCourse(courseDto dto.CourseDto, courseId string, userId string) error {
	var err error

	//检测用户合法性
	course, err := dao.CourseDao.FindCourseById(courseId)
	if err != nil || course == nil || course.UserId != userId {
		return &response.AppException{Code: response.ResultPermissionDenied}
	}

	//检验tag的合法性
	tagList, err := dao.TagDao.GetTagListByIdList(courseDto.Tags)
	if err != nil {
		return err
	}
	if len(tagList) != len(courseDto.Tags) {
		return &response.AppException{Code: response.ResultPatternError}
	}

	//这里采用先删除再新增
	//集联删除
	err = dao.CourseStepDao.DeleteByCourseId(courseId)
	err = dao.CourseTagDao.DeleteCourseTagByCourseId(courseId)

	if err != nil {
		return err
	}

	//开始重新添加
	user, err := dao.UserDao.FindUserById(userId)
	if err != nil {
		return err
	}

	//组装Search Course
	var searchCourse = entity.SearchCourse{
		Id:         courseId,
		Name:       courseDto.Name,
		Detail:     courseDto.Detail,
		Image:      courseDto.Image,
		UserId:     userId,
		UserAvatar: user.Avatar,
	}

	//保存SearchCourse
	err = dao.CourseDao.InsertSearchCourse(&searchCourse)
	if err != nil {
		return nil
	}

	//组装Course
	course = &entity.Course{
		Id:     searchCourse.Id,
		Name:   courseDto.Name,
		Detail: courseDto.Detail,
		Image:  courseDto.Image,
		UserId: userId,
	}

	//保存Course
	dao.CourseDao.InsertCourse(course)

	//组装CourseTag
	var courseTagList = make([]*entity.CourseTag, len(courseDto.Tags))
	for i, tagId := range courseDto.Tags {
		courseTagList[i] = &entity.CourseTag{
			CourseId: course.Id,
			TagId:    tagId,
		}
	}

	//保存CourseTag
	err = dao.CourseTagDao.InsertCourseTagList(courseTagList)
	if err != nil {
		return err
	}

	//组装Step
	var stepList = make([]*entity.CourseStep, len(courseDto.Step))
	for i, element := range courseDto.Step {
		step := &entity.CourseStep{
			CourseId: searchCourse.Id,
			Content:  element.Content,
			Order:    element.Order,
			Second:   element.Second,
		}

		stepList[i] = step
	}

	//保存Step
	err = dao.CourseStepDao.InsertList(stepList)
	if err != nil {
		return err
	}

	return nil
}

func (*CourseServiceImpl) DeleteCourse(courseId string, userId string) error {
	var err error

	//监测用户合法性
	course, err := dao.CourseDao.FindCourseById(courseId)
	if err != nil || course == nil || course.UserId != userId {
		return &response.AppException{Code: response.ResultPermissionDenied}
	}

	//级联删除
	err = dao.CourseStepDao.DeleteByCourseId(courseId)
	err = dao.CourseTagDao.DeleteCourseTagByCourseId(courseId)
	err = dao.CourseDao.DeleteCourse(courseId)
	err = dao.CourseDao.DeleteSearchCourse(courseId)

	return err
}

func searchCourseModelToVo(entity *entity.SearchCourseResult) *vo.SearchCourseVO {
	return &vo.SearchCourseVO{
		Id:         entity.Id,
		Name:       entity.Name,
		Detail:     entity.Detail,
		Image:      entity.Image,
		UserId:     entity.UserId,
		UserAvatar: entity.UserAvatar,
		CreateTime: entity.CreateTime,
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
