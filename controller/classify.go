package controller

import (
	"401-company/api/demo"
	"401-company/internal/service"
	"401-company/model/enum"
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_service"
)

type cClassify struct{}

var Classify = cClassify{}

func (c *cClassify) CreateClassify(ctx context.Context, req *demo.CreateClassifyReq) (*demo.ClassifyInfoRes, error) {

	//权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, enum.Classify.PermissionType.Create); has != true {
		return nil, err
	}

	result, err := service.Classify().CreateClassify(ctx, req.Classify)
	return (*demo.ClassifyInfoRes)(result), err
}

func (c *cClassify) UpdateClassify(ctx context.Context, req *demo.UpdateClassifyReq) (*demo.ClassifyInfoRes, error) {

	//权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, enum.Classify.PermissionType.Update); has != true {
		return nil, err
	}

	result, err := service.Classify().UpdateClassify(ctx, req.Classify)
	return (*demo.ClassifyInfoRes)(result), err
}

func (c *cClassify) QueryClassify(ctx context.Context, req *demo.QueryClassifyReq) (*demo.ClassifyInfoRes, error) {

	//权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, enum.Classify.PermissionType.Query); has != true {
		return nil, err
	}

	result, err := service.Classify().QueryClassify(ctx, req.Id)
	return (*demo.ClassifyInfoRes)(result), err
}

func (c *cClassify) DeleteClassify(ctx context.Context, req *demo.DeleteClassifyReq) (api_v1.BoolRes, error) {

	//权限判断
	if has, err := sys_service.SysPermission().CheckPermission(ctx, enum.Classify.PermissionType.Delete); has != true {
		return false, err
	}

	result, err := service.Classify().DeleteClassify(ctx, req.Id)
	return result == true, err
}
