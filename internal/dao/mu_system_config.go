// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe/internal/dao/internal"
)

// internalMuSystemConfigDao is internal type for wrapping internal DAO implements.
type internalMuSystemConfigDao = *internal.MuSystemConfigDao

// muSystemConfigDao is the data access object for table mu_system_config.
// You can define custom methods on it to extend its functionality as you wish.
type muSystemConfigDao struct {
	internalMuSystemConfigDao
}

var (
	// MuSystemConfig is globally public accessible object for table mu_system_config operations.
	MuSystemConfig = muSystemConfigDao{
		internal.NewMuSystemConfigDao(),
	}
)

// Fill with you ideas below.
