package swagger

import (
	"go-backend/ent"
	"go-backend/internal/common/pagination"
	"go-backend/internal/common/response"
	"go-backend/internal/dto"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
)

func article(reflector *openapi3.Reflector) error {

	{
		operationContext, err := reflector.NewOperationContext(http.MethodGet, "/api/article")
		if err != nil {
			return err
		}
		operationContext.SetTags("Article")
		operationContext.SetSummary("Lấy danh sách article")
		operationContext.SetDescription("có phân trang")
		operationContext.AddReqStructure(new(dto.ArticleFindAllReq))
		operationContext.AddRespStructure(new(response.SuccessFormat[pagination.PaginationRes[*ent.Articles]]))
		reflector.AddOperation(operationContext)
	}

	{
		operationContext, err := reflector.NewOperationContext(http.MethodGet, "/api/article/{id}")
		if err != nil {
			return err
		}
		operationContext.SetTags("Article")
		operationContext.SetSummary("Lấy chi tiết article")
		operationContext.SetDescription("description")
		operationContext.AddReqStructure(new(struct {
			Id string `path:"id" description:"id của bài viết" example:"1"`
		}))
		operationContext.AddRespStructure(new(response.SuccessFormat[*ent.Articles]))
		reflector.AddOperation(operationContext)
	}

	return nil
}
