package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var (
	TableUser       = "user"
	TableCourseTag  = "course_tag"
	TableCourse     = "course"
	TableCourseStep = "course_step"
)

func InitDatabase(username string, password string, url string, database string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, url, database)
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败%s", err.Error())
		return
	}

	DB = open
}
