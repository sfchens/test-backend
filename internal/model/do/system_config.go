// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemConfig is the golang structure of table test_system_config for DAO operations like Where/Data.
type SystemConfig struct {
	g.Meta    `orm:"table:test_system_config, do:true"`
	Id        interface{} // ID
	Name      interface{} // 名称
	Type      interface{} // 0json配置1基础配置2商城配置3用户配置
	Config    interface{} // 配置
	Operator  interface{} // 操作人
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}