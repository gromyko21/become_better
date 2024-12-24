package swagger_docs

import (
	"net/http"
)

func SwaggerFile(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Swagger UI</title>
		<link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/4.18.3/swagger-ui.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/4.18.3/swagger-ui-bundle.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/4.18.3/swagger-ui-standalone-preset.js"></script>
	</head>
	<body>
		<div id="swagger-ui"></div>
		<script>
			window.onload = function() {
				SwaggerUIBundle({
					url: "/swagger.json",
					dom_id: '#swagger-ui',
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
	`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
