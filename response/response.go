package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	ResultSuccess             = Result{Code: 200, Message: "行"}
	ResultPermissionDenied    = Result{Code: 403, Message: "你妈个逼"}
	ResultNotFound            = Result{Code: 404, Message: "艹"}
	ResultInternalServerError = Result{Code: 500, Message: "艹你妈"}

	ResultPatternError = Result{Code: 401, Message: "他妈的艹你妈不好好看文档"}
	ResultNoSuchUser   = Result{Code: 40000, Message: "没有这个人艹你妈"}
	ResultNoSuchFile   = Result{Code: 404, Message: "艹你妈逼文件没找到"}
	ResultLoginError   = Result{Code: 40001, Message: "他妈逼的你传的东西有问题艹"}
)

type AppException struct {
	Code    Result
	Message string
}

func (*AppException) Error() string {
	return ""
}

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(ctx *gin.Context, result Result) {
	ctx.JSON(http.StatusOK, result)
}

func ResponseData(ctx *gin.Context, result Result, data interface{}) {
	result.Data = data
	ctx.JSON(http.StatusOK, result)
}

func ResponseMessage(ctx *gin.Context, result Result, message string) {
	result.Message = message
	ctx.JSON(http.StatusOK, result)
}

func responseDataFull[T any](ctx *gin.Context, code int, data *T) {
	ctx.JSON(http.StatusOK, Result{
		Code:    code,
		Message: "",
		Data:    data,
	})
}

func Success(ctx *gin.Context) {
	Response(ctx, ResultSuccess)
}

func SuccessData(ctx *gin.Context, data interface{}) {
	ResponseData(ctx, ResultSuccess, data)
}

func Error(ctx *gin.Context, result Result) {
	Response(ctx, result)
}

func ErrorHandler(ctx *gin.Context, err error) {
	var exception *AppException
	if errors.As(err, &exception) {
		ErrorException(ctx, *exception)
	} else {
		Error(ctx, ResultInternalServerError)
		log.Panicln(err)
	}
	return
}

func ErrorException(ctx *gin.Context, exception AppException) {
	Response(ctx, exception.Code)
}

func ErrorMessage(ctx *gin.Context, result Result, message string) {
	ResponseMessage(ctx, result, message)
}
