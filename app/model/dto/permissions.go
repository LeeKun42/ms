// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dto

import (
	"ms/app/lib/constants"
	"ms/app/model/resp"
	"time"
)

const TableNamePermission = "permissions"

var permissionTypeMap = map[int]string{
	10: "页面权限",
	20: "操作权限",
}

// Permission mapped from table <permissions>
type Permission struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ParentID  int       `gorm:"column:parent_id;not null;comment:父级权限id：操作权限的父级id为所属页面权限id" json:"parent_id"`
	Flag      string    `gorm:"column:flag;not null;comment:权限标识" json:"flag"`
	Name      string    `gorm:"column:name;not null;comment:权限名称" json:"name"`
	Desc      string    `gorm:"column:desc;not null;comment:权限说明" json:"desc"`
	Type      int       `gorm:"column:type;not null;comment:权限类型：10：页面权限  20：操作权限" json:"type"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:最后更新时间" json:"updated_at"`
}

// TableName Permission's table name
func (*Permission) TableName() string {
	return TableNamePermission
}

func (permission *Permission) ToResponse() (roleResponse resp.PermissionResponse) {
	roleResponse.ID = permission.ID
	roleResponse.ParentID = permission.ParentID
	roleResponse.Flag = permission.Flag
	roleResponse.Name = permission.Name
	roleResponse.Desc = permission.Desc
	roleResponse.Type = permission.Type
	roleResponse.TypeText = permissionTypeMap[permission.Type]
	roleResponse.CreatedAt = permission.CreatedAt.Format(constants.DateTimeFormat)
	roleResponse.UpdatedAt = permission.UpdatedAt.Format(constants.DateTimeFormat)
	return
}