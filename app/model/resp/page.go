package resp

type PageResponse struct {
	// Index 当前页码
	PageIndex int `json:"page_index"`
	// PageSize 每页显示数据条数
	PageSize int `json:"page_size"`
	// PageCount 总页数
	PageCount int `json:"page_count"`
	// Total 总数据条数
	Total int64 `json:"total"`
}

type PageMeta struct {
	Meta PageResponse `json:"meta"`
}
