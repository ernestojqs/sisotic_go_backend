//go:build !generateauth

package generateauth

import (
	"github.com/pjmd89/gqlparser/v2/ast"
)

const Enabled = false

func GenerateAuth(schema *ast.Schema, schemaPath string) {}
