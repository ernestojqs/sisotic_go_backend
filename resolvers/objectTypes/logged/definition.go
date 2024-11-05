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
package logged

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type Logged struct {
	model     models.Logged
	userModel models.User
}

func NewLogged(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &Logged{
		model:     models.Logged{},
		userModel: models.User{},
	}
	o.(*Logged).model.Init(models.Logged{}, db)
	o.(*Logged).userModel.Init(models.User{}, db)
	return o
}
func (o *Logged) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "login":
			r, err = o.loginQuery(info)
		case "logout":
			r, err = o.logoutQuery(info)
		}
	case "mutation":
		break
	}
	return
}
func (o *Logged) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
