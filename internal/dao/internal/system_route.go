// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemRouteDao is the data access object for table test_system_route.
type SystemRouteDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns SystemRouteColumns // columns contains all the column names of Table for convenient usage.
}

// SystemRouteColumns defines and stores column names for table test_system_route.
type SystemRouteColumns struct {
	Id         string // 菜单ID
	Pid        string // 父级id
	Title      string // 名称
	Methods    string // 提交方式POST GET PUT DELETE
	Access     string // 子管理员是否可用
	Path       string // 路径
	UniqueAuth string // 前台唯一标识
	IsShow     string // 0关闭1开启
	IsDel      string // 是否删除
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	Operator   string // 操作人
}

// systemRouteColumns holds the columns for table test_system_route.
var systemRouteColumns = SystemRouteColumns{
	Id:         "id",
	Pid:        "pid",
	Title:      "title",
	Methods:    "methods",
	Access:     "access",
	Path:       "path",
	UniqueAuth: "unique_auth",
	IsShow:     "is_show",
	IsDel:      "is_del",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	Operator:   "operator",
}

// NewSystemRouteDao creates and returns a new DAO object for table data access.
func NewSystemRouteDao() *SystemRouteDao {
	return &SystemRouteDao{
		group:   "default",
		table:   "test_system_route",
		columns: systemRouteColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemRouteDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemRouteDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemRouteDao) Columns() SystemRouteColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemRouteDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemRouteDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemRouteDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}