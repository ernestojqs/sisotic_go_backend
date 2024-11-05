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
package ubication

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type Ubication struct {
	model models.Ubication
}

func NewUbication(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &Ubication{
		model: models.Ubication{},
	}
	o.(*Ubication).model.Init(models.Ubication{}, db)
	return o
}
func (o *Ubication) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		break
	case "mutation":
		break
	}
	return
}
func (o *Ubication) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
