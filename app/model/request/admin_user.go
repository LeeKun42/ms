package request

type CreateAdminUserParams struct {
	Mobile string `json:"Mobile"`
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
	Name   string `json:"name"`
	Status int    `json:"status"`
	Roles  []int  `json:"roles"`
}

type UpdateAdminUserParams struct {
	ID     int    `json:"id"`
	Mobile string `json:"Mobile"`
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
	Name   string `json:"name"`
	Status int    `json:"status"`
	Roles  []int  `json:"roles"`
}

type AdminUserLoginParams struct {
	Account string `json:"account"`
	Passwd  string `json:"passwd"`
}

type SetUserRoleParams struct {
	AdminUserID int   `json:"admin_user_id"`
	RoleIds     []int `json:"role_ids"`
}

type SearchAdminUserParams struct {
	Mobile string `json:"Mobile"`
	Email  string `json:"email"`
	Name   string `url:"name"`
	RoleId int    `url:"role_id"`
	Status int    `url:"status"`
	PageRequest
}
