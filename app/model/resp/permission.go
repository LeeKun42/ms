package resp

type PermissionResponse struct {
	ID        int                  `json:"id"`
	ParentID  int                  `json:"parent_id"`
	Flag      string               `json:"flag"`
	Name      string               `json:"name"`
	Desc      string               `json:"desc"`
	Type      int                  `json:"type"`
	TypeText  string               `json:"type_text"`
	CreatedAt string               `json:"created_at"`
	UpdatedAt string               `json:"updated_at"`
	Children  []PermissionResponse `json:"children,omitempty"`
}

type PermissionListResponse struct {
	PageMeta
	Data []PermissionResponse `json:"data"`
}
