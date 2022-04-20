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

// SearchCourse 搜索菜品
// @Summary      搜索菜品
// @Description  根据关键字搜索菜品
// @Tags         菜品
// @Produce      json
// @Responses    json
// @Param    keyword   query  string  true  "关键词"
// @Param    pageNum   query  int     true  "页数"
// @Param    pageSize  query  int     true  "页面大小"
// @Success      200       {object}  response.Result
// @Router       /course/search [GET]
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

	page, err := service.CourseService.SearchCourse(keyword, pageNum, pageSize)
	if err != nil {
		response.Error(ctx, response.ResultInternalServerError)
		log.Panic(err)
		return
	}

	response.SuccessData(ctx, page)

}

// QueryCourse 查找菜品
// @Summary  查找菜品
// @Tags     菜品
// @Param    by        query  string  true  "查找类型，可选: tag"
// @Param        keyword   query     string  true  "关键词"
// @Param        pageNum   query     int     true  "页数"
// @Param        pageSize  query     int     true  "页面大小"
// @Router   /course/query [GET]
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

	page, err := service.CourseService.GetCourseByTag(keyword, pageNum, pageSize)
	if err != nil {
		response.Error(ctx, response.ResultInternalServerError)
		return
	}

	response.SuccessData(ctx, page)
}

// GetCourseDetail 获取菜品详情
// @Summary  获取菜品详情
// @Tags     菜品
// @Param    courseId  path  string  true  "菜品id"
// @Router   /course/{courseId} [GET]
func (*CourseController) GetCourseDetail(ctx *gin.Context) {
	courseId := ctx.Param("courseId")
	detail, err := service.CourseService.GetCourseDetail(courseId)
	if err != nil {
		response.Error(ctx, response.ResultPatternError)
		return
	}

	response.SuccessData(ctx, detail)
}

func (*CourseController) GetRecommendCourseList(ctx *gin.Context) {

}

func (*CourseController) UploadCourse(ctx *gin.Context) {

}

func (*CourseController) UpdateCourse(ctx *gin.Context) {

}

func (*CourseController) DeleteCourse(ctx *gin.Context) {

}
