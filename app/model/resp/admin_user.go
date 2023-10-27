package resp

type AdminUserResponse struct {
	ID          int            `json:"id"`
	Mobile      string         `json:"mobile"`
	Email       string         `json:"email"`
	Name        string         `json:"name"`
	Avatar      string         `json:"avatar"`
	Status      int            `json:"status"`
	StatusText  string         `json:"status_text"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
	Roles       []RoleResponse `json:"roles"`
	Permissions []string       `json:"permissions,omitempty"`
}

type AdminUserListResponse struct {
	PageMeta
	Data []AdminUserResponse `json:"data"`
}
