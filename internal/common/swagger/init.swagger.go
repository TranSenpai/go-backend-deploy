package swagger

import (
	"encoding/json"
	"fmt"
	"go-backend/internal/common/env"

	"github.com/gin-gonic/gin"
	"github.com/swaggest/openapi-go/openapi3"
)

func Start(ginEngine *gin.Engine, env *env.Env) {

	midBasicAuth := gin.BasicAuth(gin.Accounts{
		"test@gmail.com":  "12345",
		"test2@gmail.com": "12345",
	})

	ginEngine.GET("docs", midBasicAuth, func(ctx *gin.Context) {
		googleLoginUrl := fmt.Sprintf("%s/api/auth/google/login?redirect=%s/docs", env.DomainBe, env.DomainBe)

		html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <meta name="color-scheme" content="light dark" />
  <title>MegaPro Chat Bot Swagger</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.32.2/swagger-ui.css" />
</head>
<body style="background: #f8f9fa;">
  <a href="%s">Google Login</a>
  <div id="swagger-ui"></div>

  <script src="https://unpkg.com/swagger-ui-dist@5.32.2/swagger-ui-bundle.js"></script>
  <script src="https://unpkg.com/swagger-ui-dist@5.32.2/swagger-ui-standalone-preset.js"></script>

  <script>
    window.onload = function () {
      window.ui = SwaggerUIBundle({
        url: "/docs.json",
        dom_id: "#swagger-ui",
        deepLinking: true,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        layout: "StandaloneLayout"
      });
    };
  </script>
</body>
</html>
`, googleLoginUrl)

		ctx.Data(200, "text/html charset=utf-8", []byte(html))
	})

	ginEngine.GET("/docs.json", midBasicAuth, func(ctx *gin.Context) {
		userAuthAny, exists := ctx.Get("user")
		if !exists {
			ctx.JSON(400, "Lỗi")
		}

		userAuth := userAuthAny.(string)

		fmt.Printf("type %T,userAuth: %+v \n\n", userAuth, userAuth)

		reflector := &openapi3.Reflector{}

		reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}

		reflector.Spec.Info.
			WithTitle("Golang Backend").
			WithVersion("1.0.0").
			WithDescription("Description")

		URLDescription := "ok"
		reflector.Spec.Servers = []openapi3.Server{
			{
				URL:         "http://localhost:3069",
				Description: &URLDescription,
			},
			{
				URL: "https://production.com",
			},
		}

		// Thêm module mới vào đây
		modules := []func(reflector *openapi3.Reflector) error{
			article,
			auth,
			user,
		}

		for _, item := range modules {
			err := item(reflector)
			if err != nil {
				return
			}
		}

		docJson, err := json.MarshalIndent(reflector.Spec, "", "     ")
		if err != nil {
			ctx.JSON(400, "Lỗi")
		}

		ctx.Data(200, "application/json charset=utf-8", docJson)
	})

}
