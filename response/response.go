package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ResultSuccess             = Result{code: 200, message: "行"}
	ResultPermissionDenied    = Result{code: 403, message: "你妈个逼"}
	ResultNotFound            = Result{code: 404, message: "艹"}
	ResultInternalServerError = Result{code: 500, message: "艹你妈"}

	ResultPatternError = Result{code: 401, message: "他妈的艹你妈不好好看文档"}
)

type AppException struct {
	Code    Result
	Message string
}

func (*AppException) Error() string {
	return ""
}

type Result struct {
	code    int
	message string
	data    interface{}
}

func Response(ctx *gin.Context, result Result) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    result.code,
		"data":    result.data,
		"message": result.message,
	})
}

func ResponseData(ctx *gin.Context, result Result, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    result.code,
		"data":    data,
		"message": result.message,
	})
}

func ResponseMessage(ctx *gin.Context, result Result, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    result.code,
		"data":    nil,
		"message": message,
	})
}

func responseDataFull[T any](ctx *gin.Context, code int, data *T) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    *data,
		"message": nil,
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

func ErrorException(ctx *gin.Context, exception AppException) {
	Response(ctx, exception.Code)
}

func ErrorMessage(ctx *gin.Context, result Result, message string) {
	ResponseMessage(ctx, result, message)
}
