package main

import (
	"cooking-backend-go/common"
	"cooking-backend-go/docs"
	"cooking-backend-go/route"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title         学厨艺后端API文档
// @version       1.0
// @description   如有问题，请联系orangeboy
// @contact.name  orangeboyChen
func main() {
	common.InitElasticSearch()

	r := gin.Default()
	r = route.CollectRoute(r)

	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//courseDao := dao.CourseDao{}
	//courseDao.SearchCourse("中国", 1, 1)
	//courseDao.GetCourseList(1, 100)
	panic(r.Run())
}
