package controller

import (
	"cooking-backend-go/response"
	"cooking-backend-go/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type CourseController struct{}

var CourseControllerInstance = CourseController{}

// SearchCourse 搜索课程
func (*CourseController) SearchCourse(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	pageNumStr := ctx.Query("pageNum")
	pageSizeStr := ctx.Query("pageSize")
	pageNum, err := strconv.Atoi(pageNumStr)
	pageSize, err := strconv.Atoi(pageSizeStr)

	if err != nil || keyword == "" {
		response.Error(ctx, response.ResultPatternError)
		return
	}

	page, err := service.CourseServiceInstance.SearchCourse(keyword, pageNum, pageSize)
	if err != nil {
		response.Error(ctx, response.ResultInternalServerError)
		log.Panic(err)
		return
	}

	response.SuccessData(ctx, page)

}

// QueryCourse 查找课程
func (*CourseController) QueryCourse(ctx *gin.Context) {
	queryType := ctx.Query("by")
	keyword := ctx.Query("keyword")
	pageNumStr := ctx.Query("pageNum")
	pageSizeStr := ctx.Query("pageSize")
	pageNum, err := strconv.Atoi(pageNumStr)
	pageSize, err := strconv.Atoi(pageSizeStr)

	if err != nil || queryType != "tag" || keyword == "" {
		response.Error(ctx, response.ResultPatternError)
		return
	}

	page, err := service.CourseServiceInstance.GetCourseByTag(keyword, pageNum, pageSize)
	if err != nil {
		response.Error(ctx, response.ResultInternalServerError)
		return
	}

	response.SuccessData(ctx, page)

}
