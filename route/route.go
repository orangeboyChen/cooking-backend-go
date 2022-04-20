package route

import (
	"cooking-backend-go/controller"
	"cooking-backend-go/response"
	"github.com/gin-gonic/gin"
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
	{
		// /login
		r.POST("/login", controller.Login)

		// /course/**
		r.Group("/course")
		{
			r.GET("/search", controller.CourseControllerInstance.SearchCourse)
			r.GET("/query", controller.CourseControllerInstance.QueryCourse)
			r.GET("/:courseId", controller.CourseControllerInstance.GetCourseDetail)
			r.GET("/recommend", controller.CourseControllerInstance.GetRecommendCourseList)
			r.POST("", controller.CourseControllerInstance.UploadCourse)
			r.PUT("/:courseId", controller.CourseControllerInstance.UpdateCourse)
			r.DELETE("/:courseId", controller.CourseControllerInstance.DeleteCourse)
		}

		// /tag/**
		r.Group("/tag")
		{
			r.GET("/type/:tagTypeId", controller.TagControllerInstance.GetTagList)
			r.GET("/type/list", controller.TagControllerInstance.GetTagTypeList)
		}

		//404
		r.NoRoute(func(ctx *gin.Context) {
			response.Error(ctx, response.ResultNotFound)
		})
	}

	return r
}
