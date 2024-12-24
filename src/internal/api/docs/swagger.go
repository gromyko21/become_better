package swagger_docs

import (
	"net/http"
	"os"
)

func SwaggerPage(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	swaggerSpec, err := os.ReadFile("internal/api/docs/proto/main.swagger.json")
	if err != nil {
		panic(err)
	}
	w.Write(swaggerSpec)
}
