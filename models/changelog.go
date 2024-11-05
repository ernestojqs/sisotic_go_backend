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
 * .schema/types.gql:7
 */
package models

import (
	"github.com/pjmd89/mongomodel/mongomodel"
)

type ChangeLog struct {
	mongomodel.Model `bson:"-" gql:"omit=true"`
	Version          string `bson:"version" gql:"name=version"`
	Date             string `bson:"date" gql:"name=date"`
	Description      string `bson:"description" gql:"name=description"`
}
