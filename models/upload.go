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
 * .schema/objectTypes/upload/type.gql:1
 */
package models

import (
	"github.com/pjmd89/mongomodel/mongomodel"
)

type Upload struct {
	mongomodel.Model `bson:"-" gql:"omit=true"`
	Name             string  `bson:"-" gql:"name=name"`
	Size             float64 `bson:"-" gql:"name=size"`
	Type             string  `bson:"-" gql:"name=type"`
	Folder           string  `bson:"-" gql:"name=folder"`
	SizeUploaded     float64 `bson:"-" gql:"name=sizeUploaded"`
	Display          string  `bson:"-" gql:"name=display"`
}
