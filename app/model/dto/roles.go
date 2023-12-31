// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dto

import (
	"ms/app/lib/constants"
	"ms/app/model/resp"
	"time"
)

const TableNameRole = "roles"

// Role mapped from table <roles>
type Role struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Flag      string    `gorm:"column:flag;not null;comment:角色flag" json:"flag"`
	Name      string    `gorm:"column:name;not null;comment:角色名称" json:"name"`
	IsSystem  int       `gorm:"column:is_system;not null;comment:是否是系统内置角色 1：是 0：否" json:"is_system"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:最后更新时间" json:"updated_at"`

	Permissions  []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}

func (role *Role) ToResponse() (roleResponse resp.RoleResponse) {
	roleResponse.ID = role.ID
	roleResponse.Flag = role.Flag
	roleResponse.Name = role.Name
	roleResponse.IsSystem = role.IsSystem
	roleResponse.CreatedAt = role.CreatedAt.Format(constants.DateTimeFormat)
	roleResponse.UpdatedAt = role.UpdatedAt.Format(constants.DateTimeFormat)
	for _, permission := range role.Permissions {
		roleResponse.PermissionIds = append(roleResponse.PermissionIds, permission.ID)
	}
	return
}