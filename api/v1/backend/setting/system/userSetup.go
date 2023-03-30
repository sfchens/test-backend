package system

import "github.com/gogf/gf/v2/frame/g"

type UserSetupAddReq struct {
	g.Meta `path:"system/user_setup/index"  tags:"管理员管理" method:"get" description:""`
	Name   string `json:"name" summary:"测试"`
	Age    int    `json:"age" summary:"cwes"`
}
type UserSetupAddRes struct {
	g.Meta `mime:"application/json" tags:"管理员管理" example:"string"`
	Name   string `json:"name" summary:"测试"`
	Age    int    `json:"age" summary:"cwes"`
}
