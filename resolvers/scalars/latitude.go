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
	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
)

type Latitude float64

func NewLatitudeScalar() (r resolvers.Scalar) {
	var scalar *Latitude
	r = scalar
	return
}

func (o *Latitude) Set(value interface{}) (r interface{}, err definitionError.GQLError) {
	r = value
	return
}

func (o *Latitude) Assess(resolved resolvers.ScalarResolved) (r interface{}, err definitionError.GQLError) {
	r = resolved.Value
	return
}
