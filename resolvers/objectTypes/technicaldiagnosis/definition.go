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
package technicaldiagnosis

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type TechnicalDiagnosis struct {
	model       models.TechnicalDiagnosis
	deviceModel models.Device
}

func NewTechnicalDiagnosis(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &TechnicalDiagnosis{
		model:       models.TechnicalDiagnosis{},
		deviceModel: models.Device{},
	}
	o.(*TechnicalDiagnosis).model.Init(models.TechnicalDiagnosis{}, db)
	o.(*TechnicalDiagnosis).deviceModel.Init(models.Device{}, db)
	return o
}
func (o *TechnicalDiagnosis) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "edges":
			r, err = o.edges(info)
		case "technicalDiagnosis":
			r, err = o.technicalDiagnosis(info)
		}
	case "mutation":
		switch info.Resolver {
		case "createTechnicalDiagnosis":
			r, err = o.createTechnicalDiagnosisMutation(info)
		case "technicalDiagnosis":
			r, err = o.technicalDiagnosis(info)
		}
	}
	return
}
func (o *TechnicalDiagnosis) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
