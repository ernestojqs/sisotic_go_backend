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
package technicaldiagnosis

import (
	"otic/lib"
	gqlErrors "otic/lib/gql_errors"
	"otic/lib/utils"
	"otic/models"
	"otic/models/enums"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (o *TechnicalDiagnosis) readDevice(id primitive.ObjectID) (device models.Device, err definitionError.GQLError) {
	result, rerr := o.model.Read(map[string]any{"_id": id, "currentSupportStatus": enums.SUPPORTSTATUSENUM_PROCESS}, nil)
	if rerr != nil || len(result.([]models.Device)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_QUERY_DEVICE_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_QUERY_DEVICE_IN_DB, nil)
		return
	}
	return result.([]models.Device)[0], nil
}

func (o *TechnicalDiagnosis) updateDevice(id primitive.ObjectID, inputValues map[string]any) (device models.Device, err definitionError.GQLError) {
	result, rerr := o.model.Update(inputValues, map[string]any{"_id": id}, nil)
	if rerr != nil || len(result.([]models.Device)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_UPDATE_DEVICE_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_UPDATE_DEVICE_IN_DB, nil)
		return
	}
	return result.([]models.Device)[0], nil
}

func (o *TechnicalDiagnosis) createTechnicalDiagnosisMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	deviceID := input["deviceID"].(primitive.ObjectID)
	sess, _ := utils.GetSession(info.SessionID)
	input["diagnosticUser"] = sess.UserID
	dbDevice, err := o.readDevice(deviceID)
	if err != nil {
		return
	}

	r, rerr := o.model.Create(input, nil)
	if rerr != nil {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_INSERT_TECHNICAL_DIAGNOSIS_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_INSERT_TECHNICAL_DIAGNOSIS_IN_DB, nil)
		return
	}

	newDiagnosis := append(dbDevice.TechnicalDiagnosis, r.(models.TechnicalDiagnosis).Id)
	o.updateDevice(deviceID, map[string]any{"technicalDiagnosis": newDiagnosis})
	return
}
