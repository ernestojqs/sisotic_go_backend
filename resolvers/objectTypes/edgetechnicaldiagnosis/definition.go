/*
 * Generado por gqlgenerate.
 *
 * Este archivo puede contener errores, de ser asi, coloca el issue en el repositorio de github
 * https://github.com/pjmd89/gogql
 *
 * Estos arvhivos corren riesgo de sobreescritura, por ese motivo gqlgnerate crea una carpeta llamada generate, asi que,
 * copia todas las carpetas que estan dentro de la carpeta generate y pegalas en la carpeta raiz de tu proyecto.
 *
 * gqlgenerate no creara archivos en la carpeta raiz de tu modulo porque puedes sufrir perdida de informacion.
 */
package edgetechnicaldiagnosis

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type EdgeTechnicalDiagnosis struct {
	model models.EdgeTechnicalDiagnosis
}

func NewEdgeTechnicalDiagnosis(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &EdgeTechnicalDiagnosis{
		model: models.EdgeTechnicalDiagnosis{},
	}
	o.(*EdgeTechnicalDiagnosis).model.Init(models.EdgeTechnicalDiagnosis{}, db)
	return o
}
func (o *EdgeTechnicalDiagnosis) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "getTechnicalDiagnosis":
			r, err = o.getTechnicalDiagnosisQuery(info)
			break
		}
		break
	case "mutation":
		break
	}
	return
}
func (o *EdgeTechnicalDiagnosis) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
