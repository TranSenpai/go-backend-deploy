package dto

import "go-backend/internal/common/pagination"

type ArticleCreateReq struct {
	Title    string  `json:"title" form:"title"`
	Content  *string `json:"content" form:"content" binding:"required"`
	ImageUrl *string `json:"image_url" form:"image_url"`
}

type ArticleFindAllFilters struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Views   *int   `json:"views"`
}

type ArticleFindAllInput struct {
	pagination.Query
	ArticleFindAllFilters
}

type ArticleFindAllReq struct {
	pagination.Query
	Filters string `query:"filters" example:"{ \"id\":1,\"content\":\"string\",\"views\":55}"`
}
