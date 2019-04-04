// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/keyauth.api.v1+json",
    "application/json"
  ],
  "produces": [
    "application/keyauth.api.v1+json",
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the REST API providing for the Szeretetfoldje daily message app",
    "title": "Szeretetfoldje daily message API",
    "version": "0.1.0"
  },
  "basePath": "/api",
  "paths": {
    "/": {
      "get": {
        "tags": [
          "daily"
        ],
        "operationId": "getDaily",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "name": "from",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list of daily items since given id",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/daily"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "daily"
        ],
        "operationId": "addOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/daily"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/daily"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{id}": {
      "get": {
        "tags": [
          "daily"
        ],
        "operationId": "getOne",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/daily"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "daily"
        ],
        "operationId": "updateOne",
        "parameters": [
          {
            "name": "daily",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/daily"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/daily"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "daily"
        ],
        "operationId": "destroyOne",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "id",
            "in": "query"
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "daily": {
      "type": "object",
      "required": [
        "message",
        "verse",
        "pray",
        "title"
      ],
      "properties": {
        "counter": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "date": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "message": {
          "type": "string",
          "minLength": 1
        },
        "pray": {
          "type": "string",
          "minLength": 1
        },
        "title": {
          "type": "string",
          "minLength": 1
        },
        "verse": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "apiKey",
      "name": "x-token",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/keyauth.api.v1+json",
    "application/json"
  ],
  "produces": [
    "application/keyauth.api.v1+json",
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the REST API providing for the Szeretetfoldje daily message app",
    "title": "Szeretetfoldje daily message API",
    "version": "0.1.0"
  },
  "basePath": "/api",
  "paths": {
    "/": {
      "get": {
        "tags": [
          "daily"
        ],
        "operationId": "getDaily",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "name": "from",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list of daily items since given id",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/daily"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "daily"
        ],
        "operationId": "addOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/daily"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/daily"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{id}": {
      "get": {
        "tags": [
          "daily"
        ],
        "operationId": "getOne",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/daily"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "daily"
        ],
        "operationId": "updateOne",
        "parameters": [
          {
            "name": "daily",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/daily"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/daily"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "daily"
        ],
        "operationId": "destroyOne",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "id",
            "in": "query"
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "daily": {
      "type": "object",
      "required": [
        "message",
        "verse",
        "pray",
        "title"
      ],
      "properties": {
        "counter": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "date": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        },
        "message": {
          "type": "string",
          "minLength": 1
        },
        "pray": {
          "type": "string",
          "minLength": 1
        },
        "title": {
          "type": "string",
          "minLength": 1
        },
        "verse": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "principal": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "apiKey",
      "name": "x-token",
      "in": "header"
    }
  }
}`))
}
