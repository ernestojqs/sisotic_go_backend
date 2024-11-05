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
package systeminfo

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type SystemInfo struct {
	model models.SystemInfo
}

func NewSystemInfo(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &SystemInfo{
		model: models.SystemInfo{},
	}
	o.(*SystemInfo).model.Init(models.SystemInfo{}, db)
	return o
}
func (o *SystemInfo) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "systemInfo":
			r, err = o.systemInfoQuery(info)
			break
		}
		break
	case "mutation":
		break
	}
	return
}
func (o *SystemInfo) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
