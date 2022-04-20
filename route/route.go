package route

import (
	"cooking-backend-go/controller"
	"cooking-backend-go/docs"
	"cooking-backend-go/response"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"runtime/debug"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(
		//middleware.AuthMiddleware(),
		func(ctx *gin.Context) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("panic: %v\n", r)
					debug.PrintStack()
					response.Error(ctx, response.ResultInternalServerError)
				}
			}()
			ctx.Next()
		})
	// /login
	r.POST("/login", controller.UserControllerImpl.Login)

	// /courseRoute/**
	courseRoute := r.Group("/course")
	{
		courseRoute.GET("/search", controller.CourseControllerInstance.SearchCourse)
		courseRoute.GET("/query", controller.CourseControllerInstance.QueryCourse)
		courseRoute.GET("/:courseId", controller.CourseControllerInstance.GetCourseDetail)
		courseRoute.GET("/recommend", controller.CourseControllerInstance.GetRecommendCourseList)
		courseRoute.POST("", controller.CourseControllerInstance.UploadCourse)
		courseRoute.PUT("/:courseId", controller.CourseControllerInstance.UpdateCourse)
		courseRoute.DELETE("/:courseId", controller.CourseControllerInstance.DeleteCourse)
	}

	// /tag/**
	tagRoute := r.Group("/tag")
	{
		tagRoute.GET("/type/list", controller.TagControllerInstance.GetTagTypeList)
		tagRoute.GET("/type/:tagTypeId", controller.TagControllerInstance.GetTagList)
	}

	//404
	r.NoRoute(func(ctx *gin.Context) {
		response.Error(ctx, response.ResultNotFound)
	})

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
