package fanbox

import (
	_ "github.com/ogen-go/ogen"
)

// TODO: ここで使用している API のスキーマを、バージョン管理された artifact のものを使用するよう修正
//go:generate go run github.com/ogen-go/ogen/cmd/ogen -v --clean --config cfg.yaml --target . --package fanbox openapi.yaml
