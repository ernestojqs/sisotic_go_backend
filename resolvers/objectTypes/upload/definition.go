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
package upload

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type Upload struct {
	model models.Upload
}

func NewUpload(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &Upload{
		model: models.Upload{},
	}
	o.(*Upload).model.Init(models.Upload{}, db)
	return o
}
func (o *Upload) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "checkFileSize":
			r, err = o.checkFileSizeQuery(info)
			break
		}
		break
	case "mutation":
		switch info.Resolver {
		case "upload":
			r, err = o.uploadMutation(info)
			break
		case "cancelUpload":
			r, err = o.cancelUploadMutation(info)
			break
		}
		break
	}
	return
}
func (o *Upload) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
