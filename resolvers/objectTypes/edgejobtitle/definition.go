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
package edgejobtitle

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type EdgeJobTitle struct {
	model models.EdgeJobTitle
}

func NewEdgeJobTitle(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &EdgeJobTitle{
		model: models.EdgeJobTitle{},
	}
	o.(*EdgeJobTitle).model.Init(models.EdgeJobTitle{}, db)
	return o
}
func (o *EdgeJobTitle) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "getJobTitles", "jobTitles":
			r, err = o.getJobTitlesQuery(info)
		}
	case "mutation":
		switch info.Resolver {
		case "jobTitles":
			r, err = o.getJobTitlesQuery(info)
		}
	}
	return
}
func (o *EdgeJobTitle) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
