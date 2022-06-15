package service

import (
	"cooking-backend-go/dto"
	"cooking-backend-go/entity"
	"cooking-backend-go/service/impl"
	"cooking-backend-go/vo"
)

var (
	CourseService     courseService     = &impl.CourseServiceImpl{}
	TagService        tagService        = &impl.TagServiceImpl{}
	UserService       userService       = &impl.UserServiceImpl{}
	IngredientService ingredientService = &impl.IngredientServiceImpl{}
)

type courseService interface {
	SearchCourse(keyword string, pageNum int, pageSize int) (*vo.PageVO[vo.SearchCourseVO], error)
	GetCourseByTag(tagId string, pageNum int, pageSize int) (*vo.PageVO[vo.CourseVO], error)
	GetCourseDetail(courseId string) (*vo.CourseDetailVO, error)
	GetCourseRecommendation() ([]*vo.SearchCourseVO, error)
	InsertCourse(courseDto dto.CourseDto, userId string) (string, error)
	UpdateCourse(courseDto dto.CourseDto, courseId string, userId string) error
	DeleteCourse(courseId string, userId string) error
}

type tagService interface {
	GetTagList(tagTypeId string) ([]*vo.TagVO, error)
	GetTagTypeList() ([]*vo.TagTypeVO, error)
}

type userService interface {
	Login(dto dto.UserLoginDto) (string, error)
	SetAvatar(userId string, avatarFilePath string) error
	GetAvatar(userId string) (string, error)
	UpdateUserInfo(userInfoDto dto.UserInfoDto, userId string) error
	FindUserById(userId string) (*vo.UserInfoVO, error)
}

type ingredientService interface {
	GetIngredientByCourseId(courseId string) ([]*entity.Ingredient, error)
	ConveyModelToVo(model *entity.Ingredient) *vo.IngredientVO
}

type mealService interface {
}
