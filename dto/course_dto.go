package dto

type CourseDto struct {
	Name   string
	Detail string
	Image  string
	Tags   []string
	Step   []CourseStepDto
}

type CourseStepDto struct {
	Order   int
	Content string
	Second  int
}