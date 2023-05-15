// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MuRbacRole is the golang structure of table mu_rbac_role for DAO operations like Where/Data.
type MuRbacRole struct {
	g.Meta      `orm:"table:mu_rbac_role, do:true"`
	Id          interface{} //
	Name        interface{} // 名称
	Description interface{} // 描述
	Operator    interface{} // 操作人
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
}
