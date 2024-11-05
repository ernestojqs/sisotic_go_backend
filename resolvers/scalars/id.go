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
package scalars

import (
	gqlErrors "otic/lib/gql_errors"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ID string

func NewIDScalar() (r resolvers.Scalar) {
	var scalar *ID
	r = scalar
	return
}

func (o *ID) Set(value interface{}) (r interface{}, err definitionError.GQLError) {

	if value != nil {
		newObjectId, mistake := primitive.ObjectIDFromHex(value.(string))
		if mistake != nil {
			//log.Println("invalid ID: ", value)
			errExtension := definitionError.ExtensionError{
				"code": gqlErrors.ERROR_EMPTY_OR_INVALID_ID.Code,
			}
			return nil, definitionError.NewFatal(gqlErrors.ERROR_EMPTY_OR_INVALID_ID.Message, errExtension)
		}
		r = newObjectId
	}
	return
}

func (o *ID) Assess(resolved resolvers.ScalarResolved) (r interface{}, err definitionError.GQLError) {
	r = resolved.Value
	return
}
