package model

import (
	"401-company/model/entity"
)

type Commodity struct {
	Id         int64  `json:"id"           dc:"商品id"`
	Name       string `json:"name"         dc:"商品名称"    v:"required#请输入商品名称"`
	ClassifyId int64  `json:"classify_id"  dc:"所属分类id"  v:"required#请输入所属分类id"`
	Des        string `json:"des"          dc:"商品描述"`
}

type CommodityInfo entity.Commodities
