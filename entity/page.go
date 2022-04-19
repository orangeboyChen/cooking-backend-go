package entity

type Page[T any] struct {
	PageNum int
	Total   int
	Data    []*T
}
