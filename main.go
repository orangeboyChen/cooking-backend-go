package main

import (
	"cooking-backend-go/common"
	"cooking-backend-go/common/elastic_config"
	"cooking-backend-go/entity"
	"cooking-backend-go/route"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
)

// @title                       不叫外卖后端API文档
// @version                     1.0
// @description                 如有问题，请联系orangeboy
// @contact.name                orangeboyChen
// @securityDefinitions.apikey  ApiAuthToken
// @in                          header
// @name                        Authorization
func main() {
	initConfig()
	initDatabase()

	r := gin.Default()
	r = route.CollectRoute(r)

	//courseDao := dao.CourseDao{}
	//courseDao.SearchCourse("中国", 1, 1)
	//courseDao.GetCourseList(1, 100)
	panic(r.Run())
}

func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	viper.ReadInConfig()
}

func initDatabase() {
	common.InitDatabase(
		viper.GetString("go.datasource.username"),
		viper.GetString("go.datasource.password"),
		viper.GetString("go.datasource.url"),
		viper.GetString("go.datasource.database"))
	elastic_config.InitElasticSearch()

	func() {
		//开始建表
		err := common.DB.AutoMigrate(&entity.User{}, &entity.Course{}, &entity.CourseTag{}, &entity.CourseStep{}, &entity.Tag{}, &entity.TagType{})
		if err != nil {
			log.Fatal("妈的数据库建表失败了，你他妈去写建表语句去自己去建表吧艹", err)
		}
		log.Print("已执行建表")
	}()
}
