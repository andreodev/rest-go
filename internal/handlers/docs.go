package handlers

import "net/http"

func (h Handlers) registerDocsEndpoints() {
	http.HandleFunc("GET /docs", h.docsPage)
	http.HandleFunc("GET /openapi.json", h.openapiJSON)
}

func (h Handlers) docsPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`<!doctype html>
<html lang="pt-BR">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>REST Go API Docs</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.onload = function () {
      SwaggerUIBundle({
        url: "/openapi.json",
        dom_id: "#swagger-ui"
      });
    };
  </script>
</body>
</html>`))
}

func (h Handlers) openapiJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{
  "openapi": "3.0.3",
  "info": {
    "title": "REST Go API",
    "description": "Documentacao das rotas da API REST Go.",
    "version": "1.0.0"
  },
  "servers": [
    {
      "url": "http://localhost:8080",
      "description": "Servidor local"
    }
  ],
  "paths": {
    "/products": {
      "get": {
        "summary": "Lista produtos",
        "parameters": [
          {
            "name": "page",
            "in": "query",
            "schema": {
              "type": "integer",
              "default": 1,
              "minimum": 1
            }
          },
          {
            "name": "limit",
            "in": "query",
            "schema": {
              "type": "integer",
              "default": 10,
              "minimum": 1
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Produtos retornados com sucesso",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GetAllProductsResponse"
                }
              }
            }
          },
          "400": {
            "$ref": "#/components/responses/BadRequest"
          },
          "500": {
            "$ref": "#/components/responses/InternalServerError"
          }
        }
      },
      "post": {
        "summary": "Cria um produto",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/CreateProductRequest"
              },
              "example": {
                "name_product": "Teclado mecanico",
                "price": 149.9,
                "description": "Teclado compacto com switches azuis"
              }
            }
          }
        },
        "responses": {
          "200": {
            "description": "Produto criado com sucesso",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/CreateProductResponse"
                }
              }
            }
          },
          "400": {
            "$ref": "#/components/responses/BadRequest"
          },
          "500": {
            "$ref": "#/components/responses/InternalServerError"
          }
        }
      }
    },
    "/users": {
      "get": {
        "summary": "Lista usuarios",
        "responses": {
          "200": {
            "description": "Usuarios retornados com sucesso",
            "content": {
              "application/json": {
                "schema": {
                  "type": "array",
                  "items": {
                    "$ref": "#/components/schemas/User"
                  }
                }
              }
            }
          },
          "500": {
            "$ref": "#/components/responses/InternalServerError"
          }
        }
      },
      "post": {
        "summary": "Cria um usuario",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/UserCreateRequest"
              },
              "example": {
                "name": "Andre",
                "email": "andre@example.com"
              }
            }
          }
        },
        "responses": {
          "201": {
            "description": "Usuario criado com sucesso",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/UserCreateResponse"
                }
              }
            }
          },
          "400": {
            "$ref": "#/components/responses/BadRequest"
          },
          "500": {
            "$ref": "#/components/responses/InternalServerError"
          }
        }
      }
    },
    "/users/{id}": {
      "get": {
        "summary": "Busca usuario por ID",
        "parameters": [
          {
            "$ref": "#/components/parameters/UserID"
          }
        ],
        "responses": {
          "200": {
            "description": "Usuario retornado com sucesso",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/User"
                }
              }
            }
          },
          "400": {
            "$ref": "#/components/responses/BadRequest"
          },
          "404": {
            "$ref": "#/components/responses/NotFound"
          },
          "500": {
            "$ref": "#/components/responses/InternalServerError"
          }
        }
      },
      "delete": {
        "summary": "Remove usuario por ID",
        "parameters": [
          {
            "$ref": "#/components/parameters/UserID"
          }
        ],
        "responses": {
          "200": {
            "description": "Usuario removido com sucesso",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/UserDeleteResponse"
                }
              }
            }
          },
          "400": {
            "$ref": "#/components/responses/BadRequest"
          },
          "404": {
            "$ref": "#/components/responses/NotFound"
          },
          "500": {
            "$ref": "#/components/responses/InternalServerError"
          }
        }
      }
    }
  },
  "components": {
    "parameters": {
      "UserID": {
        "name": "id",
        "in": "path",
        "required": true,
        "schema": {
          "type": "string",
          "format": "uuid"
        }
      }
    },
    "responses": {
      "BadRequest": {
        "description": "Requisicao invalida",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/ErrorResponse"
            }
          }
        }
      },
      "NotFound": {
        "description": "Recurso nao encontrado",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/ErrorResponse"
            }
          }
        }
      },
      "InternalServerError": {
        "description": "Erro interno",
        "content": {
          "application/json": {
            "schema": {
              "$ref": "#/components/schemas/ErrorResponse"
            }
          }
        }
      }
    },
    "schemas": {
      "ErrorResponse": {
        "type": "object",
        "properties": {
          "reason": {
            "type": "string"
          }
        }
      },
      "Product": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string",
            "format": "uuid"
          },
          "name_product": {
            "type": "string"
          },
          "price": {
            "type": "number",
            "format": "double"
          },
          "description": {
            "type": "string"
          },
          "created_at": {
            "type": "string",
            "format": "date-time"
          }
        }
      },
      "CreateProductRequest": {
        "type": "object",
        "required": [
          "name_product",
          "price",
          "description"
        ],
        "properties": {
          "name_product": {
            "type": "string"
          },
          "price": {
            "type": "number",
            "format": "double"
          },
          "description": {
            "type": "string"
          }
        }
      },
      "CreateProductResponse": {
        "type": "object",
        "properties": {
          "id": {
            "type": "string",
            "format": "uuid"
          },
          "name_product": {
            "type": "string"
          },
          "price": {
            "type": "number",
            "format": "double"
          },
          "description": {
            "type": "string"
          },
          "created_at": {
            "type": "string",
            "format": "date-time"
          }
        }
      },
      "GetAllProductsResponse": {
        "type": "object",
        "properties": {
          "products": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/Product"
            }
          },
          "page": {
            "type": "integer"
          },
          "limit": {
            "type": "integer"
          },
          "total": {
            "type": "integer"
          }
        }
      },
      "User": {
        "type": "object",
        "properties": {
          "ID": {
            "type": "string",
            "format": "uuid"
          },
          "Name": {
            "type": "string"
          },
          "Email": {
            "type": "string",
            "format": "email"
          }
        }
      },
      "UserCreateRequest": {
        "type": "object",
        "required": [
          "name",
          "email"
        ],
        "properties": {
          "name": {
            "type": "string"
          },
          "email": {
            "type": "string",
            "format": "email"
          }
        }
      },
      "UserCreateResponse": {
        "type": "object",
        "properties": {
          "newUserId": {
            "type": "string",
            "format": "uuid"
          }
        }
      },
      "UserDeleteResponse": {
        "type": "object",
        "properties": {
          "message": {
            "type": "string"
          },
          "id": {
            "type": "string",
            "format": "uuid"
          }
        }
      }
    }
  }
}`))
}
