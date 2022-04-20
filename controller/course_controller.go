package controller

import (
	"cooking-backend-go/dto"
	"cooking-backend-go/response"
	"cooking-backend-go/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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
		response.ErrorHandler(ctx, err)
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
		response.ErrorHandler(ctx, err)
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
		response.ErrorHandler(ctx, err)
		return
	}

	response.SuccessData(ctx, detail)
}

// GetRecommendCourseList 获取推荐列表
// @Summary  获取推荐列表
// @Tags     菜品
// @Router   /course/recommend [GET]
func (*CourseController) GetRecommendCourseList(ctx *gin.Context) {
	courseList, err := service.CourseService.GetCourseRecommendation()
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}

	response.SuccessData(ctx, courseList)
}

// UploadCourse 上传菜品
// @Summary  上传菜品
// @Tags     菜品
// @Param    dto  body  dto.CourseDto  true  "菜品详情"
// @Router   /course [POST]
func (*CourseController) UploadCourse(ctx *gin.Context) {
	request := ctx.Request
	bodyByte, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Error(ctx, response.ResultPatternError)
		return
	}

	var courseDto dto.CourseDto
	err = json.Unmarshal(bodyByte, &courseDto)
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}

	userId := request.Header.Get("userId")

	userId, err = service.CourseService.InsertCourse(courseDto, userId)
}

// UpdateCourse 更新菜品
// @Summary  更新菜品
// @Tags     菜品
// @Param    courseId  path  string         true  "courseId"
// @Param    dto       body  dto.CourseDto  true  "菜品详情"
// @Router   /course/{courseId} [PUT]
func (*CourseController) UpdateCourse(ctx *gin.Context) {
	request := ctx.Request
	courseId := ctx.Param("courseId")
	userId := request.Header.Get("userId")
	bodyByte, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.Error(ctx, response.ResultPatternError)
		return
	}

	var courseDto dto.CourseDto
	err = json.Unmarshal(bodyByte, &courseDto)
	if err != nil {
		response.Error(ctx, response.ResultPatternError)
		return
	}
	err = service.CourseService.UpdateCourse(courseDto, courseId, userId)
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}
	response.Success(ctx)
}

// DeleteCourse 删除菜品
// @Summary  删除菜品
// @Tags     菜品
// @Param    courseId  path  string  true  "courseId"
// @Router   /course/{courseId} [DELETE]
func (*CourseController) DeleteCourse(ctx *gin.Context) {
	courseId := ctx.Param("courseId")
	userId := ctx.Request.Header.Get("userId")
	err := service.CourseService.DeleteCourse(courseId, userId)
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}
	response.Success(ctx)
}
