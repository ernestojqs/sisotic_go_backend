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
 * .schema/objectTypes/technical_diagnosis/type.gql:1
 */
package models

import (
	"otic/resolvers/scalars"

	"github.com/pjmd89/mongomodel/mongomodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TechnicalDiagnosis struct {
	mongomodel.Model   `bson:"-" gql:"omit=true"`
	Id                 primitive.ObjectID   `bson:"_id,omitempty" gql:"name=_id,id=true,objectID=true"`
	Description        string               `bson:"description" gql:"name=description"`
	ResolverActivities []primitive.ObjectID `bson:"resolverActivities" gql:"name=resolverActivities,objectID=true"`
	DiagnosticUser     primitive.ObjectID   `bson:"diagnosticUser" gql:"name=diagnosticUser,objectID=true"`
	Created            scalars.DateTime     `bson:"created" gql:"name=created,created=true"`
	Updated            scalars.DateTime     `bson:"updated" gql:"name=updated,updated=true"`
}
