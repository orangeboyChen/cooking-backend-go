package dao

import (
	"cooking-backend-go/dao/impl"
	"cooking-backend-go/entity"
)

var (
	CourseDao           courseDao           = &impl.CourseDaoImpl{}
	TagDao              tagDao              = &impl.TagDaoImpl{}
	UserDao             userDao             = &impl.UserDaoImpl{}
	CourseStepDao       courseStepDao       = &impl.CourseStepDaoImpl{}
	CourseTagDao        courseTagDao        = &impl.CourseTagDaoImpl{}
	IngredientDao       ingredientDao       = &impl.IngredientDaoImpl{}
	IngredientCourseDao ingredientCourseDao = &impl.IngredientCourseDaoImpl{}
	MealDao             mealDao             = &impl.MealDaoImpl{}
	MealCourseDao       mealCourseDao       = &impl.MealCourseDao{}
)

type courseDao interface {
	SearchCourse(keyword string, pageNum int, pageSize int) (*entity.Page[entity.SearchCourseResult], error)
	GetCourseList(pageNum int, pageSize int) (*entity.Page[entity.SearchCourseResult], error)
	InsertSearchCourse(course *entity.SearchCourse) error
	InsertCourse(course *entity.Course)
	FindCourseByTagId(tagId string, pageNum int, pageSize int) (*entity.Page[entity.Course], error)
	FindCourseById(courseId string) (*entity.Course, error)
	GetRecommendationCourse() ([]*entity.SearchCourseResult, error)
	DeleteCourse(courseId string) error
	DeleteSearchCourse(courseId string) error
}

type courseStepDao interface {
	FindCourseStepByCourseId(courseId string) ([]*entity.CourseStep, error)
	InsertList(list []*entity.CourseStep) error
	DeleteByCourseId(courseId string) error
}

type tagDao interface {
	GetTagList(tagTypeId string) ([]*entity.Tag, error)
	GetTagTypeList() ([]*entity.TagType, error)
	GetTagListByIdList(tagIdList []string) ([]*entity.Tag, error)
}

type userDao interface {
	InsertUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	FindUserById(id string) (*entity.User, error)
	FindUserByUserIdList(idList []string) ([]*entity.User, error)
	FindUserByOpenid(openid string) (*entity.User, error)
}

type courseTagDao interface {
	InsertCourseTagList(list []*entity.CourseTag) error
	DeleteCourseTagByCourseId(courseId string) error
}

type ingredientDao interface {
	InsertIngredient(ingredient *entity.Ingredient) error
	FindIngredientByIdList(idList []string) ([]*entity.Ingredient, error)
}

type ingredientCourseDao interface {
	FindIngredientCourseByCourseId(courseId string) ([]*entity.IngredientCourse, error)
	FindIngredientCourseByCourseIdList(courseIdList []string) ([]*entity.IngredientCourse, error)
}

type mealDao interface {
	InsertMeal(meal *entity.Meal) error
}

type mealCourseDao interface {
}
