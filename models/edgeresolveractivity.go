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
 * .schema/objectTypes/resolver_activity/type.gql:14
 */
package models

import (
	"github.com/pjmd89/mongomodel/mongomodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EdgeResolverActivity struct {
	mongomodel.Model `bson:"-" gql:"omit=true"`
	Edges            []primitive.ObjectID `bson:"edges" gql:"name=edges"`
	PageInfo         PageInfo             `bson:"pageInfo" gql:"name=pageInfo"`
}
