package controller

import (
	"401-company/api/demo"
	"401-company/internal/service"
	"401-company/model/enum"
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

type cCommodity struct{}

var Commodity = cCommodity{}

func (c *cCommodity) CreateCommodity(ctx context.Context, req *demo.CreateCommodityReq) (*demo.CommodityInfoRes, error) {

	//权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, enum.Commodity.PermissionType.Create); has != true {
		return nil, err
	}

	result, err := service.Commodity().CreateCommodity(ctx, req.Commodity)
	return (*demo.CommodityInfoRes)(result), err
}

func (c *cCommodity) UpdateCommodity(ctx context.Context, req *demo.UpdateCommodityReq) (*demo.CommodityInfoRes, error) {

	//权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, enum.Commodity.PermissionType.Update); has != true {
		return nil, err
	}

	result, err := service.Commodity().UpdateCommodity(ctx, req.Commodity)
	return (*demo.CommodityInfoRes)(result), err
}

func (c *cCommodity) QueryCommodity(ctx context.Context, req *demo.QueryCommodityReq) (*demo.CommodityInfoRes, error) {

	//权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, enum.Commodity.PermissionType.Query); has != true {
		return nil, err
	}

	result, err := service.Commodity().QueryCommodity(ctx, req.Id)
	return (*demo.CommodityInfoRes)(result), err
}

func (c *cCommodity) DeleteCommodity(ctx context.Context, req *demo.DeleteCommodityReq) (api_v1.BoolRes, error) {

	//权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, enum.Commodity.PermissionType.Delete); has != true {
		return false, err
	}

	result, err := service.Commodity().DeleteCommodity(ctx, req.Id)
	return result == true, err
}
