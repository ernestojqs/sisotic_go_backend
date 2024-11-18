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
 * .schema/objectTypes/device/type.gql:1
 */
package models

import (
	"otic/models/enums"
	"otic/resolvers/scalars"

	"github.com/pjmd89/mongomodel/mongomodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	mongomodel.Model     `bson:"-" gql:"omit=true"`
	Id                   primitive.ObjectID      `bson:"_id,omitempty" gql:"name=_id,id=true,objectID=true"`
	PlaceOfCare          enums.PlaceOfCareEnum   `bson:"placeOfCare" gql:"name=placeOfCare"`
	CollegeDependency    primitive.ObjectID      `bson:"collegeDependency" gql:"name=collegeDependency,objectID=true"`
	IsSupport            bool                    `bson:"isSupport" gql:"name=isSupport"`
	AssetCode            string                  `bson:"assetCode" gql:"name=assetCode"`
	SerialCode           string                  `bson:"serialCode" gql:"name=serialCode"`
	Type                 string                  `bson:"type" gql:"name=type"`
	Modell               string                  `bson:"model" gql:"name=model"`
	GroupID              string                  `bson:"groupID" gql:"name=groupID"`
	Observations         []Observation           `bson:"observations" gql:"name=observations"`
	TransferEvidences    []TransferEvidence      `bson:"transferEvidences" gql:"name=transferEvidences"`
	ReceiverUser         primitive.ObjectID      `bson:"receiverUser" gql:"name=receiverUser,objectID=true"`
	TechnicalDiagnosis   []primitive.ObjectID    `bson:"technicalDiagnosis" gql:"name=technicalDiagnosis,objectID=true"`
	CurrentSupportStatus enums.SupportStatusEnum `bson:"currentSupportStatus" gql:"name=currentSupportStatus"`
	SupportStatusDetails []SupportStatusDetails  `bson:"supportStatusDetails" gql:"name=supportStatusDetails"`
	ResolvedDate         scalars.DateTime        `bson:"estimatedDeliveryDate" gql:"name=estimatedDeliveryDate"`
	Created              scalars.DateTime        `bson:"created" gql:"name=created,created=true"`
	Updated              scalars.DateTime        `bson:"updated" gql:"name=updated,updated=true"`
}
