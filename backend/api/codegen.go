//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target ../pkg/oas -package oas --clean --convenient-errors=on ./openapi.yml
package api

import _ "github.com/ogen-go/ogen"
