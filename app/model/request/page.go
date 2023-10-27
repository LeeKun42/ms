package request

type PageRequest struct {
	// Index 当前页码
	PageIndex int `url:"page_index"`
	// PageSize 每页显示数据条数
	PageSize int `url:"page_size"`
}
