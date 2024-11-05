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
package transferevidence

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type TransferEvidence struct {
	model models.TransferEvidence
}

func NewTransferEvidence(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &TransferEvidence{
		model: models.TransferEvidence{},
	}
	o.(*TransferEvidence).model.Init(models.TransferEvidence{}, db)
	return o
}
func (o *TransferEvidence) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		break
	case "mutation":
		break
	}
	return
}
func (o *TransferEvidence) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
