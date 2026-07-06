package dto

import "go-backend/internal/common/pagination"

type ChatGroupFindAllFilters struct {
	Name string `json:"name"`
}

type ChatGroupFindAllInput struct {
	pagination.Query
	ChatGroupFindAllFilters
}
