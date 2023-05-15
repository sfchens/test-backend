// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MuRabcAdmin is the golang structure for table mu_rabc_admin.
type MuRabcAdmin struct {
	Id        uint        `json:"id"        description:"ID"`
	Username  string      `json:"username"  description:"账号"`
	Phone     string      `json:"phone"     description:"手机号码"`
	Password  string      `json:"password"  description:"密码"`
	Email     string      `json:"email"     description:"邮箱"`
	PicUrl    string      `json:"picUrl"    description:"头像"`
	Status    uint        `json:"status"    description:"状态 0未激活 1激活 2禁用 3删除"`
	Operator  string      `json:"operator"  description:"操作人"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
}
