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
 *
 * Este type fue obtenido del archivo:
 * .schema/directives/ubication/type.gql:1
 */
package models

import (
	"otic/resolvers/scalars"

	"github.com/pjmd89/mongomodel/mongomodel"
)

type Ubication struct {
	mongomodel.Model `bson:"-" gql:"omit=true"`
	Lat              scalars.Latitude  `bson:"lat" gql:"name=lat"`
	Lon              scalars.Longitude `bson:"lon" gql:"name=lon"`
}
