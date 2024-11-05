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
package edgedevice

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type EdgeDevice struct {
	model models.EdgeDevice
}

func NewEdgeDevice(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &EdgeDevice{
		model: models.EdgeDevice{},
	}
	o.(*EdgeDevice).model.Init(models.EdgeDevice{}, db)
	return o
}
func (o *EdgeDevice) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "getDevices":
			r, err = o.getDevicesQuery(info)
			break
		}
		break
	case "mutation":
		break
	}
	return
}
func (o *EdgeDevice) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
