package impl

import (
	"context"
	"cooking-backend-go/common"
	"cooking-backend-go/common/elastic_config"
	"cooking-backend-go/entity"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
	"time"
)

type CourseDaoImpl struct{}

func (*CourseDaoImpl) SearchCourse(keyword string, pageNum int, pageSize int) (*entity.Page[entity.SearchCourseResult], error) {
	query := elastic.NewBoolQuery()
	query.Should(
		elastic.NewMatchQuery("name", keyword),
		elastic.NewMatchQuery("detail", keyword),
	)
	res, err := elastic_config.ESClient.Search(elastic_config.CourseIndex).Query(query).From((pageNum - 1) * pageSize).Size(pageSize).Do(context.Background())
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return conveyESResultToPage(res, pageNum), nil
}

func (*CourseDaoImpl) GetCourseList(pageNum int, pageSize int) (*entity.Page[entity.SearchCourseResult], error) {
	res, err := elastic_config.ESClient.Search(elastic_config.CourseIndex).From((pageNum - 1) * pageSize).Size(pageSize).Do(context.Background())
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return conveyESResultToPage(res, pageNum), nil
}

func (*CourseDaoImpl) InsertSearchCourse(course *entity.SearchCourse) error {
	course.CreateTime = time.Now().UnixMilli()
	res, err := elastic_config.ESClient.Index().Index(elastic_config.CourseIndex).BodyJson(&course).Do(context.Background())
	if err != nil {
		return err
	}

	course.Id = res.Id
	return nil
}

func (*CourseDaoImpl) InsertCourse(course *entity.Course) {
	common.DB.Create(course)
}

func (*CourseDaoImpl) FindCourseByTagId(tagId string, pageNum int, pageSize int) (*entity.Page[entity.Course], error) {
	var courseIdList []string
	if err := common.DB.Model(&entity.CourseTag{}).Where("tag_id = ?", tagId).Pluck("course_id", &courseIdList).Error; err != nil {
		return nil, err
	}

	var count int64
	if err := common.DB.Model(&entity.Course{}).Where("id in (?)", courseIdList).Count(&count).Error; err != nil {
		return nil, err
	}

	var courseList []entity.Course
	if err := common.DB.Where("id in (?)", courseIdList).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&courseList).Error; err != nil {
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

func (*CourseDaoImpl) FindCourseById(courseId string) (*entity.Course, error) {
	var course entity.Course
	if err := common.DB.Find(&course, courseId).Error; err != nil {
		return nil, err
	}

	if course.Id == "" {
		return nil, nil
	}

	return &course, nil
}

func (*CourseDaoImpl) GetRecommendationCourse() ([]*entity.SearchCourseResult, error) {
	q := elastic.NewFunctionScoreQuery()
	q = q.AddScoreFunc(elastic.NewRandomFunction())
	res, err := elastic_config.ESClient.Search(elastic_config.CourseIndex).Query(q).Size(10).Do(context.Background())
	if err != nil {
		return nil, err
	}

	var data = make([]*entity.SearchCourseResult, len(res.Hits.Hits))
	for i, item := range res.Each(reflect.TypeOf(entity.SearchCourseResult{})) {
		course := item.(entity.SearchCourseResult)
		data[i] = &course
	}

	return data, nil
}

func (*CourseDaoImpl) DeleteCourse(courseId string) error {
	return common.DB.Delete("id = ?", courseId).Error
}

func (*CourseDaoImpl) DeleteSearchCourse(courseId string) error {
	_, err := elastic_config.ESClient.Delete().Id(courseId).Do(context.Background())
	return err
}

func conveyESResultToPage(res *elastic.SearchResult, pageNum int) *entity.Page[entity.SearchCourseResult] {
	var data = make([]*entity.SearchCourseResult, len(res.Hits.Hits))
	for i, hit := range res.Hits.Hits {
		json.Unmarshal(hit.Source, &data[i])
		data[i].Score = *hit.Score
	}

	var page = &entity.Page[entity.SearchCourseResult]{
		PageNum: pageNum,
		Total:   int(res.Hits.TotalHits.Value),
		Data:    data,
	}

	return page
}
