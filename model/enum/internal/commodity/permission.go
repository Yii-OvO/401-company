package commodity

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/utility/permission"
)

type PermissionTypeEnum = *sys_model.SysPermissionTree

type permissionType struct {
	Query  PermissionTypeEnum
	Create PermissionTypeEnum
	Update PermissionTypeEnum
	Delete PermissionTypeEnum
}

var PermissionType = permissionType{
	Query:  permission.New(6413883105738821, "ViewDetail", "查看商品", "查看某个商品"),
	Create: permission.New(6413888678854725, "Create", "创建商品", "创建一个新商品"),
	Update: permission.New(6413884825010245, "Update", "更新商品信息", "更新某个商品信息"),
	Delete: permission.New(6413886221582405, "Delete", "删除商品", "删除某个商品"),
}
