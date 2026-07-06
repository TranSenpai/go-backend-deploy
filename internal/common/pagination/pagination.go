package pagination

import "strconv"

type Query struct {
	Page     int `query:"page" description:"số trang" example:"1"`
	PageSize int `query:"pageSize" description:"số lượng phần tử trong 1 trang" example:"3"`
	Offset   int
}

type PaginationRes[T any] struct {
	Items     T   `json:"items"`
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalItem int `json:"totalItem"`
	TotalPage int `json:"totalPage"`
}

func Get(pageString string, pageSizeString string) Query {
	page, err := strconv.Atoi(pageString)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeString)
	if err != nil || pageSize <= 0 {
		pageSize = 3
	}

	offset := (page - 1) * pageSize

	return Query{
		Page:     page,
		PageSize: pageSize,
		Offset:   offset,
	}
}
