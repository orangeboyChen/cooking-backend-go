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

	ResultPatternError = Result{code: 23333, message: "妈的日了狗了"}
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
}

func responseFull(ctx *gin.Context, code int, data gin.H, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}

func responseMessage(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    nil,
		"message": message,
	})
}

func responseData(ctx *gin.Context, code int, data *gin.H) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": nil,
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
	responseMessage(ctx, ResultSuccess.code, ResultSuccess.message)
}

func SuccessJson(ctx *gin.Context, data *gin.H) {
	responseData(ctx, ResultSuccess.code, data)
}

func SuccessData[T any](ctx *gin.Context, data *T) {
	responseDataFull(ctx, ResultSuccess.code, data)
}

func Error(ctx *gin.Context, result Result) {
	responseMessage(ctx, result.code, result.message)
}

func ErrorMessage(ctx *gin.Context, result Result, message string) {
	responseMessage(ctx, result.code, message)
}
