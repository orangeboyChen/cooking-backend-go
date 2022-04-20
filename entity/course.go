package entity

type Course struct {
	Id         string `gorm:"column:id;primaryKey;not null"`
	Name       string `gorm:"column:name;not null"`
	Detail     string `gorm:"column:detail;not null"`
	Image      string `gorm:"column:image"`
	UserId     string `gorm:"column:user_id"`
	CreateTime int64  `gorm:"column:create_time"`
}

type SearchCourseResult struct {
	Id          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Detail      string  `json:"detail,omitempty"`
	Image       string  `json:"image,omitempty"`
	UserId      string  `json:"userId,omitempty"`
	UserAvatar  string  `json:"userAvatar,omitempty"`
	CreateTime  int64   `json:"createTime,omitempty"`
	Score       float64 `json:"score"`
	NameWithHit string  `json:"nameWithHit"`
}

type SearchCourse struct {
	Id         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Detail     string `json:"detail,omitempty"`
	Image      string `json:"image,omitempty"`
	UserId     string `json:"userId,omitempty"`
	UserAvatar string `json:"userAvatar,omitempty"`
	CreateTime int64  `json:"createTime,omitempty"`
}

type CourseTag struct {
	Id       string `gorm:"primaryKey;column:id"`
	CourseId string `gorm:"column:course_id"`
	TagId    string `gorm:"column:tag_id"`
}

type CourseStep struct {
	Id       string `gorm:"primaryKey;column:id"`
	CourseId string `gorm:"column:id"`
	Content  string `gorm:"column:content"`
	Order    int    `gorm:"column:order"`
	Second   int    `gorm:"column:second"`
}
