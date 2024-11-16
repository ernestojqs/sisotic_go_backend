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
package supportstatusdetails

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
)

type SupportStatusDetails struct {
	model models.SupportStatusDetails
}

func NewSupportStatusDetails() (o resolvers.ObjectTypeInterface) {
	o = &SupportStatusDetails{
		model: models.SupportStatusDetails{},
	}
	return o
}
func (o *SupportStatusDetails) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query", "mutation":
		switch info.Resolver {
		case "supportStatusDetails":
			r = info.Parent.(models.Device).SupportStatusDetails
		}
	}
	return
}
func (o *SupportStatusDetails) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
