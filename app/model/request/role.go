package request

type SearchRoleParams struct {
	Flag string `url:"flag"`
	Name string `url:"name"`
	PageRequest
}

type CreateRoleParams struct {
	ID   int    `json:"id"`
	Flag string `json:"flag"`
	Name string `json:"name"`
}

type SetRolePermissionsParams struct {
	RoleID        int   `json:"role_id"`
	PermissionIds []int `json:"permission_ids"`
}
