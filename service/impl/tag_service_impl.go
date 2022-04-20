package impl

import (
	"cooking-backend-go/dao"
	"cooking-backend-go/entity"
	"cooking-backend-go/vo"
)

type TagServiceImpl struct{}

func (*TagServiceImpl) GetTagList(tagTypeId string) ([]*vo.TagVO, error) {
	list, err := dao.TagDao.GetTagList(tagTypeId)
	if err != nil {
		return nil, err
	}

	result := make([]*vo.TagVO, len(list))
	for i := range list {
		result[i] = tagModelToVO(list[i])
	}

	return result, nil
}

func (*TagServiceImpl) GetTagTypeList() ([]*vo.TagTypeVO, error) {
	list, err := dao.TagDao.GetTagTypeList()
	if err != nil {
		return nil, err
	}

	result := make([]*vo.TagTypeVO, len(list))
	for i := range list {
		result[i] = tagTypeModelToVO(list[i])
	}

	return result, nil
}

func tagModelToVO(entity *entity.Tag) *vo.TagVO {
	return &vo.TagVO{
		Id:   entity.Id,
		Name: entity.Name,
	}
}

func tagTypeModelToVO(entity *entity.TagType) *vo.TagTypeVO {
	return &vo.TagTypeVO{
		Id:   entity.Id,
		Name: entity.Name,
	}
}
