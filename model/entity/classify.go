// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Classify is the golang structure for table classify.
type Classify struct {
	Id        int64       `json:"id"        description:"分类ID"`
	Name      string      `json:"name"      description:"分类名称"`
	ParentId  int64       `json:"parentId"  description:"父级ID"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
