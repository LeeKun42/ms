// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dto

import (
	"time"
)

const TableNameAdminUserSocial = "admin_user_socials"

// AdminUserSocial mapped from table <admin_user_socials>
type AdminUserSocial struct {
	AdminUserID int       `gorm:"column:admin_user_id;primaryKey;comment:管理员账号id" json:"admin_user_id"`
	Type        string    `gorm:"column:type;primaryKey;comment:社会化登录类型:wechat、dingding、enterprisewechat" json:"type"`
	Unionid     string    `gorm:"column:unionid;not null" json:"unionid"`
	Openid      string    `gorm:"column:openid;not null" json:"openid"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:最后更新时间" json:"updated_at"`
}

// TableName AdminUserSocial's table name
func (*AdminUserSocial) TableName() string {
	return TableNameAdminUserSocial
}
