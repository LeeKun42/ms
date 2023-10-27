package request

type CreatePermissionParams struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Flag     string `json:"flag"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Type     int    `json:"type"`
}

type SearchPermissionParams struct {
	ParentID int    `url:"parent_id"`
	Flag     string `url:"flag"`
	Name     string `url:"name"`
	Type     int    `url:"type"`
	PageRequest
}
