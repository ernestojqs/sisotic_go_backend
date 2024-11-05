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
package edgejobarea

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type EdgeJobArea struct {
	model models.EdgeJobArea
}

func NewEdgeJobArea(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &EdgeJobArea{
		model: models.EdgeJobArea{},
	}
	o.(*EdgeJobArea).model.Init(models.EdgeJobArea{}, db)
	return o
}
func (o *EdgeJobArea) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "getJobAreas":
			r, err = o.getJobAreasQuery(info)
		}
	case "mutation":
		break
	}
	return
}
func (o *EdgeJobArea) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
