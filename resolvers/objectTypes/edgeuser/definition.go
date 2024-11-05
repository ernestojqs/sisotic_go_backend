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
package edgeuser

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type EdgeUser struct {
	model models.EdgeUser
}

func NewEdgeUser(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &EdgeUser{
		model: models.EdgeUser{},
	}
	o.(*EdgeUser).model.Init(models.EdgeUser{}, db)
	return o
}
func (o *EdgeUser) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "getUsers":
			r, err = o.getUsersQuery(info)
		}
	case "mutation":
		break
	}
	return
}
func (o *EdgeUser) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
