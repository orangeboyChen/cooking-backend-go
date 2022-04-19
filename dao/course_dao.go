package dao

import (
	"context"
	"cooking-backend-go/common"
	"cooking-backend-go/entity"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"log"
	"time"
)

type CourseDao struct {
}

var CourseDaoInstance = CourseDao{}

func (*CourseDao) SearchCourse(keyword string, pageNum int, pageSize int) (*entity.Page[entity.SearchCourse], error) {
	query := elastic.NewBoolQuery()
	query.Should(
		elastic.NewMatchQuery("name", keyword),
		elastic.NewMatchQuery("detail", keyword),
	)
	res, err := common.ESClient.Search(common.CourseIndex).Query(query).From((pageNum - 1) * pageSize).Size(pageSize).Do(context.Background())
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return conveyESResultToPage(res, pageNum), nil
}

func (*CourseDao) GetCourseList(pageNum int, pageSize int) (*entity.Page[entity.SearchCourse], error) {
	res, err := common.ESClient.Search(common.CourseIndex).From((pageNum - 1) * pageSize).Size(pageSize).Do(context.Background())
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return conveyESResultToPage(res, pageNum), nil
}

func (*CourseDao) InsertCourse(course *entity.Course) error {
	course.CreateTime = time.Now().UnixMilli()
	res, err := common.ESClient.Index().Index(common.CourseIndex).BodyJson(&course).Do(context.Background())
	if err != nil {
		return err
	}

	course.Id = res.Id
	return nil
}

func (*CourseDao) InsertSearchCourse(course *entity.SearchCourse) {
	common.DB.Create(course)
}

func (*CourseDao) FindCourseByTagId(tagId string, pageNum int, pageSize int) (*entity.Page[entity.Course], error) {
	var courseIdList []string
	if err := common.DB.Table(common.TableCourseTag).Where("tag_id = ?", tagId).Pluck("course_id", &courseIdList).Error; err != nil {
		return nil, err
	}

	var count int64
	if err := common.DB.Table(common.TableCourse).Where("id in (?)").Count(&count).Error; err != nil {
		return nil, err
	}

	var courseList []entity.Course
	if err := common.DB.Table(common.TableCourse).Where("id in (?)", courseIdList).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&courseList).Error; err != nil {
		return nil, err
	}

	var coursePointerList = make([]*entity.Course, len(courseList))
	for i := range courseList {
		coursePointerList[i] = &courseList[i]
	}

	return &entity.Page[entity.Course]{
		PageNum: pageNum,
		Total:   int(count),
		Data:    coursePointerList,
	}, nil
}

func (*CourseDao) FindCourseById(courseId string) (*entity.Course, error) {
	var course entity.Course
	if err := common.DB.Table(common.TableCourse).Find(&course, courseId).Error; err != nil {
		return nil, err
	}

	if course.Id == "" {
		return nil, nil
	}

	return &course, nil
}

func (*CourseDao) FindCourseStepByCourseId(courseId string) ([]*entity.CourseStep, error) {
	var courseStepList []entity.CourseStep
	if err := common.DB.Table(common.TableCourseStep).Where("course_id = ?", courseId).Find(&courseStepList).Error; err != nil {
		return nil, err
	}

	var courseStepPointerList = make([]*entity.CourseStep, len(courseStepList))
	for i := range courseStepList {
		courseStepPointerList[i] = &courseStepList[i]
	}

	return courseStepPointerList, nil
}

func conveyESResultToPage(res *elastic.SearchResult, pageNum int) *entity.Page[entity.SearchCourse] {
	var data = make([]*entity.SearchCourse, len(res.Hits.Hits))
	for i, hit := range res.Hits.Hits {
		json.Unmarshal(hit.Source, &data[i])
		data[i].Score = *hit.Score
	}

	var page = &entity.Page[entity.SearchCourse]{
		PageNum: pageNum,
		Total:   int(res.Hits.TotalHits.Value),
		Data:    data,
	}

	return page
}
