//go:build generateauth

package generateauth

import (
	"github.com/pjmd89/gogql/cmd/authorization"
	"github.com/pjmd89/gogql/lib/generate"
	"github.com/pjmd89/gqlparser/v2/ast"
)

const Enabled = true

func GenerateAuth(schema *ast.Schema, schemaPath string) {
	var genData = generate.NewGqlGenerate(schema, "schema")
	authorization.Generate(genData)
}
