// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dto

const TableNameAdminUserRole = "admin_user_roles"

// AdminUserRole mapped from table <admin_user_roles>
type AdminUserRole struct {
	AdminUserID int `gorm:"column:admin_user_id;primaryKey" json:"admin_user_id"`
	RoleID      int `gorm:"column:role_id;primaryKey" json:"role_id"`
}

// TableName AdminUserRole's table name
func (*AdminUserRole) TableName() string {
	return TableNameAdminUserRole
}
