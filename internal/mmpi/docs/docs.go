package docs

import "github.com/swaggo/swag"

const docTemplate = `{
  "openapi": "3.0.0",
  "info": {
    "title": "MMPI Test API",
    "version": "1.0",
    "description": "API for MMPI psychological test administration",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "name": "API Support",
      "email": "support@fittin.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    }
  },
  "servers": [
    {
      "url": "http://localhost:8080",
      "description": "Development server"
    }
  ],
  "paths": {},
  "components": {
    "schemas": {}
  }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MMPI Test API",
	Description:      "API for MMPI psychological test administration",
	InfoInstanceName: "swag",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
