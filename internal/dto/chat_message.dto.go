package dto

import "go-backend/internal/common/pagination"

type SendMessageReturn struct {
	MessageText string
	ChatGroupId int
	UserId      int
}

type ChatMessageFindAllFilters struct {
	MessageText string `json:"messageText"`
	ChatGroupId int    `json:"chatGroupId"`
}

type ChatMessageFindAllInput struct {
	pagination.Query
	ChatMessageFindAllFilters
}
