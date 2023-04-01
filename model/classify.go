package model

import (
	"401-company/model/entity"
)

type Classify struct {
	Id       int64  `json:"id"        dc:"分类id"`
	Name     string `json:"name"      dc:"分类名称"      v:"required#请输入分类名称"`
	ParentId int64  `json:"parent_id" dc:"父级id"`
}

type ClassifyInfo entity.Classify
