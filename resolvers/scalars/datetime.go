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
	"log"
	gqlErrors "otic/lib/gql_errors"
	"strconv"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
)

type DateTime int64

func NewDateTimeScalar() (r resolvers.Scalar) {
	var scalar *DateTime
	r = scalar
	return
}

func (o *DateTime) Set(value interface{}) (r interface{}, err definitionError.GQLError) {
	r = value
	if value != nil {
		var er error
		r, er = strconv.ParseInt(value.(string), 10, 64)
		if er != nil {
			log.Println("Wrong Datetime: ", value)
			errExtension := definitionError.ExtensionError{
				"code": gqlErrors.ERROR_INVALID_DATETIME.Code,
			}
			err = definitionError.NewFatal(gqlErrors.ERROR_INVALID_DATETIME.Message, errExtension)
			return
		}
		r = r.(int64)
	}
	return
}

func (o *DateTime) Assess(resolved resolvers.ScalarResolved) (r interface{}, err definitionError.GQLError) {
	r = resolved.Value
	return
}
