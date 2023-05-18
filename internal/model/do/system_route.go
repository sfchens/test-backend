// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemRoute is the golang structure of table test_system_route for DAO operations like Where/Data.
type SystemRoute struct {
	g.Meta     `orm:"table:test_system_route, do:true"`
	Id         interface{} // 菜单ID
	Pid        interface{} // 父级id
	Title      interface{} // 名称
	Methods    interface{} // 提交方式POST GET PUT DELETE
	Access     interface{} // 子管理员是否可用
	Path       interface{} // 路径
	UniqueAuth interface{} // 前台唯一标识
	Status     interface{} // 0关闭1开启
	Sort       interface{} // 排序
	Operator   interface{} // 操作人
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
