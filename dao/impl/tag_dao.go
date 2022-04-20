package impl

import (
	"cooking-backend-go/common"
	"cooking-backend-go/entity"
	"cooking-backend-go/utils"
)

type TagDaoImpl struct{}

func (*TagDaoImpl) GetTagList(tagTypeId string) ([]*entity.Tag, error) {
	var tagList []entity.Tag
	if err := common.DB.Table(common.TableTag).Select("tag_type_id = ?", tagTypeId).Find(&tagList).Error; err != nil {
		return nil, err
	}

	return utils.ToPointerList(tagList), nil
}

func (*TagDaoImpl) GetTagListByIdList(tagIdList []string) ([]*entity.Tag, error) {
	var tagList []entity.Tag
	if err := common.DB.Table(common.TableTag).Select("tag_type_id in (?)", tagIdList).Find(&tagList).Error; err != nil {
		return nil, err
	}

	return utils.ToPointerList(tagList), nil
}

func (*TagDaoImpl) GetTagTypeList() ([]*entity.TagType, error) {
	var tagTypeList []entity.TagType
	if err := common.DB.Table(common.TableTagType).Find(&tagTypeList).Error; err != nil {
		return nil, err
	}

	return utils.ToPointerList(tagTypeList), nil
}
