package controller

import (
	"cooking-backend-go/response"
	"cooking-backend-go/service"
	"github.com/gin-gonic/gin"
)

type TagController struct{}

var TagControllerInstance = TagController{}

// GetTagList 获取标签列表
// @Tags      标签
// @Summary   获取标签列表
// @Security  ApiAuthToken
// @Param     typeId  path  string  true  "tagTypeId"
// @Router    /tag/type/{tagTypeId} [GET]
func (*TagController) GetTagList(ctx *gin.Context) {

	tagTypeId := ctx.Param("tagTypeId")
	if tagTypeId == "" {
		response.Error(ctx, response.ResultPatternError)
		return
	}

	list, err := service.TagService.GetTagList(tagTypeId)
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}

	response.SuccessData(ctx, list)
}

// GetTagTypeList 获取标签种类列表
// @Tags      标签
// @Summary   获取标签种类列表
// @Security  ApiAuthToken
// @Router    /tag/type/list [GET]
func (*TagController) GetTagTypeList(ctx *gin.Context) {
	list, err := service.TagService.GetTagTypeList()
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}

	response.SuccessData(ctx, list)
}
