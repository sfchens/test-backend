// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RbacPermission is the golang structure of table my_rbac_permission for DAO operations like Where/Data.
type RbacPermission struct {
	g.Meta      `orm:"table:my_rbac_permission, do:true"`
	Id          interface{} //
	Name        interface{} // 名称
	Path        interface{} // 路径
	Description interface{} // 描述
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
}
