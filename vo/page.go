package vo

import "cooking-backend-go/entity"

type PageVO[T any] struct {
	PageNum int  `json:"pageNum,omitempty"`
	Total   int  `json:"total,omitempty"`
	Data    []*T `json:"data,omitempty"`
}

type EntityToVOInterface[T any, K any] interface {
	convey(entity T) K
}

func ConveyPageToPageVO[T any, K any](entityPage *entity.Page[T], convey func(*T) *K) *PageVO[K] {
	var entityData = entityPage.Data
	var voData = make([]*K, len(entityData))

	for i, element := range entityData {
		voData[i] = convey(element)
	}

	var vo = PageVO[K]{
		PageNum: entityPage.PageNum,
		Total:   entityPage.Total,
		Data:    voData,
	}

	return &vo
}
