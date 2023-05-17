// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemAdminDao is the data access object for table test_system_admin.
type SystemAdminDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns SystemAdminColumns // columns contains all the column names of Table for convenient usage.
}

// SystemAdminColumns defines and stores column names for table test_system_admin.
type SystemAdminColumns struct {
	Id         string // ID
	Username   string // 后台管理员账号
	HeadPic    string // 管理员头像
	Password   string // 后台管理员密码
	RealName   string // 后台管理员姓名
	Roles      string // 后台管理员权限
	LastIp     string // 后台管理员最后一次登录ip
	LastTime   string // 后台管理员最后一次登录时间
	LoginCount string // 登录次数
	Level      string // 后台管理员级别
	Status     string // 后台管理员状态 1有效0无效
	DivisionId string // 事业部id
	IsDel      string // 是否删除
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	Operator   string // 操作人
	Phone      string // 手机号码
}

// systemAdminColumns holds the columns for table test_system_admin.
var systemAdminColumns = SystemAdminColumns{
	Id:         "id",
	Username:   "username",
	HeadPic:    "head_pic",
	Password:   "password",
	RealName:   "real_name",
	Roles:      "roles",
	LastIp:     "last_ip",
	LastTime:   "last_time",
	LoginCount: "login_count",
	Level:      "level",
	Status:     "status",
	DivisionId: "division_id",
	IsDel:      "is_del",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	Operator:   "operator",
	Phone:      "phone",
}

// NewSystemAdminDao creates and returns a new DAO object for table data access.
func NewSystemAdminDao() *SystemAdminDao {
	return &SystemAdminDao{
		group:   "default",
		table:   "test_system_admin",
		columns: systemAdminColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemAdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemAdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemAdminDao) Columns() SystemAdminColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemAdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemAdminDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemAdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
