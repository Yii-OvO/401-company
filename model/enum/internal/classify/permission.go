package classify

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
	Query:  permission.New(6413863559036997, "ViewDetail", "查看分类", "查看某个分类"),
	Create: permission.New(6413867020976197, "Create", "创建分类", "创建一个新分类"),
	Update: permission.New(6413869167083589, "Update", "更新分类信息", "更新某个分类信息"),
	Delete: permission.New(6413870771011653, "Delete", "删除分类", "删除某个分类"),
}
