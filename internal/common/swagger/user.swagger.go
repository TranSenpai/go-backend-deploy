package swagger

import (
	"go-backend/ent"
	"go-backend/internal/common/response"
	"mime/multipart"
	"net/http"

	"github.com/swaggest/openapi-go/openapi3"
)

func user(reflector *openapi3.Reflector) error {
	{
		operationContext, err := reflector.NewOperationContext(http.MethodPost, "/api/user/avatar-local")
		if err != nil {
			return err
		}
		operationContext.SetTags("Users")
		operationContext.SetSummary("Upload Avatar")
		operationContext.SetDescription("description")
		operationContext.AddReqStructure(new(struct {
			Avatar  multipart.File   `formData:"avatar"`
			Avatars []multipart.File `formData:"avatars"`
		}))
		operationContext.AddRespStructure(new(response.SuccessFormat[*ent.Users]))
		reflector.AddOperation(operationContext)
	}
	return nil
}
