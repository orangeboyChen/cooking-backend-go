package dto

type CourseDto struct {
	Name   string
	Detail string
	Image  string
	Tags   []string `example:"id1,id2"`
	Step   []CourseStepDto
}

type CourseStepDto struct {
	Order   int
	Content string
	Second  int
}
