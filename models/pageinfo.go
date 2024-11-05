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
 * .schema/directives/paginate/type.gql:1
 */
package models

import (
	"github.com/pjmd89/mongomodel/mongomodel"
)

type PageInfo struct {
	mongomodel.Model `bson:"-" gql:"omit=true"`
	Page             int64 `bson:"page" gql:"name=page"`
	Pages            int64 `bson:"pages" gql:"name=pages"`
	Shown            int64 `bson:"shown" gql:"name=shown"`
	Total            int64 `bson:"total" gql:"name=total"`
	Overall          int64 `bson:"overall" gql:"name=overall"`
	Split            int64 `bson:"split" gql:"name=split"`
}
