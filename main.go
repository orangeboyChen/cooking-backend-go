package main

import (
	"cooking-backend-go/common"
	"cooking-backend-go/route"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitElasticSearch()

	r := gin.Default()
	r = route.CollectRoute(r)

	//courseDao := dao.CourseDao{}
	//courseDao.SearchCourse("中国", 1, 1)
	//courseDao.GetCourseList(1, 100)
	panic(r.Run())
}
