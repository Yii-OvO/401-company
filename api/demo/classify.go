package demo

import (
	"401-company/model"
	"github.com/gogf/gf/v2/frame/g"
)

type CreateClassifyReq struct {
	g.Meta `path:"/create"  method:"post"  summary:"创建分类"  tags:"分类"`
	model.Classify
}

type UpdateClassifyReq struct {
	g.Meta `path:"/update"  method:"post"  summary:"更新分类"  tags:"分类"`
	model.Classify
}

type QueryClassifyReq struct {
	g.Meta `path:"/query"  method:"post"  summary:"查询分类"  tags:"分类"`
	Id     int64 `json:"id" v:"required#缺少分类id参数" dc:"分类ID"`
}

type DeleteClassifyReq struct {
	g.Meta `path:"/delete"  method:"post"  summary:"删除分类"  tags:"分类"`
	Id     int64 `json:"id" v:"required#缺少分类id参数" dc:"分类ID"`
}

type ClassifyInfoRes model.ClassifyInfo
