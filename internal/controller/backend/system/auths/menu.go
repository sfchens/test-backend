package auths

import (
	"context"
	"goframe/api/v1/backend/system/auths"
	"goframe/internal/service"
)

type cMenu struct {
}

var NewMenu = &cMenu{}

// Add 添加菜单
// @summary 添加菜单
// @description 添加菜单
// @tags 设置管理
// @accept application/json
// @produce application/json
// @params body body auth.MenuAddReq
// @success 200 {object} response.ResponseRes={data=auth.AdminLoginRes} "code错误码 msg操作信息 data返回信息"
// @route /backend/setting/auth_menu/add [post]
func (c *cMenu) Add(ctx context.Context, req *auths.MenuAddReq) (res *auths.MenuAddRes, err error) {
	err = service.AuthsMenu().Add(ctx, req)
	return
}

// // Edit 编辑菜单
// // @summary 编辑菜单
// // @description 编辑菜单
// // @tags 设置管理
// // @accept application/json
// // @produce application/json
// // @params body body auth.MenuEditReq
// // @success 200 {object} response.ResponseRes={data=auth.MenuEditRes} "code错误码 msg操作信息 data返回信息"
// // @route /backend/setting/auth_menu/edit [post]
// func (c *cMenu) Edit(ctx context.Context, req *auth.MenuEditReq) (res *auth.MenuEditRes, err error) {
// 	err = service.authMenu().Edit(ctx, req)
// 	return
// }

// // List 菜单列表
// // @summary 菜单列表
// // @description 菜单列表
// // @tags 设置管理
// // @accept application/json
// // @produce application/json
// // @params body body auth.MenuListReq
// // @success 200 {object} response.ResponseRes={data=auth.MenuListRes} "code错误码 msg操作信息 data返回信息"
// // @route /backend/setting/auth_menu/list [get]
// func (c *cMenu) List(ctx context.Context, req *auth.MenuListReq) (res *auth.MenuListRes, err error) {
// 	var (
// 		out []setting.MenuListItem
// 	)
// 	out, err = service.authMenu().List(ctx, req)
// 	if err != nil {
// 		return
// 	}
// 	res = &auth.MenuListRes{
// 		List: out,
// 	}
// 	return
// }