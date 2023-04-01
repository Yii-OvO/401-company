package demo

import (
	"401-company/model"
	"github.com/gogf/gf/v2/frame/g"
)

type CreateCommodityReq struct {
	g.Meta `path:"/create"  method:"post"  summary:"创建商品"  tags:"商品"`
	model.Commodity
}

type UpdateCommodityReq struct {
	g.Meta `path:"/update"  method:"post"  summary:"更新商品"  tags:"商品"`
	model.Commodity
}

type QueryCommodityReq struct {
	g.Meta `path:"/query"  method:"post"  summary:"查询商品"  tags:"商品"`
	Id     int64 `json:"id" v:"required#缺少商品id参数"`
}

type DeleteCommodityReq struct {
	g.Meta `path:"/delete"  method:"post"  summary:"删除商品"  tags:"商品"`
	Id     int64 `json:"id" v:"required#缺少商品id参数"`
}

type CommodityInfoRes model.CommodityInfo
