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
		r.GET("/course/search", controller.CourseControllerInstance.SearchCourse)
		r.GET("/course/query", controller.CourseControllerInstance.QueryCourse)
		r.GET("/course/:courseId")
		r.GET("/course/recommend")
		r.POST("/course")
		r.PUT("/course/:courseId")
		r.DELETE("/course/:courseId")

		// /tag/**
		r.GET("/tag/list")
		r.GET("/tag/type/list")

		r.NoRoute(func(ctx *gin.Context) {
			response.Error(ctx, response.ResultNotFound)
		})
	}

	return r
}
