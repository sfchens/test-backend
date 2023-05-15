// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure of table my_admin for DAO operations like Where/Data.
type Admin struct {
	g.Meta     `orm:"table:my_admin, do:true"`
	Id         interface{} //
	Username   interface{} // 角色ID
	Password   interface{} // 密码
	Phone      interface{} // 手机号码
	Email      interface{} // 邮箱
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}