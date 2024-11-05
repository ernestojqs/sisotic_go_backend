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
package collegedependency

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type CollegeDependency struct {
	model models.CollegeDependency
}

func NewCollegeDependency(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &CollegeDependency{
		model: models.CollegeDependency{},
	}
	o.(*CollegeDependency).model.Init(models.CollegeDependency{}, db)
	return o
}
func (o *CollegeDependency) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "edges":
			r, err = o.edges(info)
		case "collegeDependency":
			r, err = o.collegeDependency(info)
		}
	case "mutation":
		switch info.Resolver {
		case "createCollegeDependency":
			r, err = o.createCollegeDependencyMutation(info)
		case "updateCollegeDependency":
			r, err = o.updateCollegeDependencyMutation(info)
		case "deleteCollegeDependency":
			r, err = o.deleteCollegeDependencyMutation(info)
		case "collegeDependency":
			r, err = o.collegeDependency(info)
		}
	}
	return
}
func (o *CollegeDependency) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
