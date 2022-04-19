package vo

type CourseVO struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Detail     string `json:"detail"`
	Image      string `json:"image"`
	UserId     string `json:"user_id"`
	UserAvatar string `json:"userAvatar"`
	CreateTime int64  `json:"create_time"`
}

type SearchCourseVO struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Detail      string `json:"detail,omitempty"`
	Image       string `json:"image,omitempty"`
	UserId      string `json:"userId,omitempty"`
	UserAvatar  string `json:"userAvatar,omitempty"`
	NameWithHit string `json:"nameWithHit,omitempty"`
	CreateTime  int64  `json:"createTime,omitempty"`
}

type CourseDetailVO struct {
	Id         string
	Name       string
	Detail     string
	Image      string
	UserId     string
	UserAvatar string
	Step       []CourseStepVO
	CreateTime int64
}

type CourseStepVO struct {
	Order   int
	Content string
	Second  int
}
