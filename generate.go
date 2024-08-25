package fanbox

import (
	_ "github.com/ogen-go/ogen"
)

//go:generate go run github.com/ogen-go/ogen/cmd/ogen -v --clean --config cfg.yaml --target . --package fanbox openapi.yaml
