package resp

type RoleResponse struct {
	ID            int    `json:"id"`
	Flag          string `json:"flag"`
	Name          string `json:"name"`
	IsSystem      int    `json:"is_system"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	PermissionIds []int  `json:"permission_ids"`
}

type RoleListResponse struct {
	PageMeta
	Data []RoleResponse `json:"data"`
}
