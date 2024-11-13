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
package device

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

func (o *Device) createDevicesMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	sess, _ := utils.GetSession(info.SessionID)
	input := info.Args["input"].(map[string]any)
	r = []models.Device{}
	collegeDependencyID := input["collegeDependency"].(primitive.ObjectID)
	if err = o.readCollegeDependency(collegeDependencyID); err != nil {
		return
	}
	for _, v := range input["deviceInfo"].([]any) {
		deviceInfo := v.(map[string]any)
		deviceInfo["receiverUser"] = sess.UserID
		deviceInfo["groupID"] = utils.GenerateTokenFromUUID(12, true)
		deviceInfo["collegeDependency"] = collegeDependencyID
		deviceInfo["placeOfCare"] = input["placeOfCare"].(string)
		deviceInfo["supportStatus"] = string(enums.SUPPORTSTATUSENUM_DELIVERED)
		if deviceInfo["placeOfCare"].(string) == string(enums.PLACEOFCAREENUM_LOCAL) {
			if deviceInfo["isSupport"] == true && ((deviceInfo["localTechnicalDiagnosis"] == nil) || (deviceInfo["localTechnicalDiagnosis"] != nil && len(deviceInfo["localTechnicalDiagnosis"].([]any)) == 0)) {
				lib.Logs.System.Warning().Println(gqlErrors.ERROR_TECHNICAL_DIAGNOSIS_REQUERIDED)
				err = definitionError.NewError(gqlErrors.ERROR_TECHNICAL_DIAGNOSIS_REQUERIDED, nil)
				return
			}
			var techDiagnosisIDs []primitive.ObjectID
			for _, v := range deviceInfo["localTechnicalDiagnosis"].([]any) {
				techDiagnosis := v.(map[string]any)
				if (techDiagnosis["localResolverActivities"] == nil) || (techDiagnosis["localResolverActivities"] != nil && len(techDiagnosis["localResolverActivities"].([]any)) == 0) {
					lib.Logs.System.Warning().Println(gqlErrors.ERROR_TECHNICAL_DIAGNOSIS_REQUERIDED)
					err = definitionError.NewError(gqlErrors.ERROR_TECHNICAL_DIAGNOSIS_REQUERIDED, nil)
					return
				}
				var resolverActivityIDs []primitive.ObjectID
				for _, v := range techDiagnosis["localResolverActivities"].([]any) {
					resolverActivityInput := v.(map[string]any)
					resolverActivityInput["resolverUser"] = sess.UserID
					result, rerr := o.resolverActivityModel.Create(resolverActivityInput, nil)
					if rerr != nil {
						lib.Logs.System.Warning().Println(gqlErrors.ERROR_INSERT_RESOLVER_ACTIVITY_IN_DB)
						err = definitionError.NewError(gqlErrors.ERROR_INSERT_RESOLVER_ACTIVITY_IN_DB, nil)
						return
					}
					resolverActivityIDs = append(resolverActivityIDs, result.(models.ResolverActivity).Id)
				}
				techDiagnosis["resolverActivities"] = resolverActivityIDs
				techDiagnosis["diagnosticUser"] = sess.UserID
				createdDiagnosis, rerr := o.technicalDiagnosisModel.Create(techDiagnosis, nil)
				if rerr != nil {
					lib.Logs.System.Warning().Println(gqlErrors.ERROR_INSERT_TECHNICAL_DIAGNOSIS_IN_DB)
					err = definitionError.NewError(gqlErrors.ERROR_INSERT_TECHNICAL_DIAGNOSIS_IN_DB, nil)
					return
				}
				techDiagnosisIDs = append(techDiagnosisIDs, createdDiagnosis.(models.TechnicalDiagnosis).Id)
			}
			deviceInfo["technicalDiagnosis"] = techDiagnosisIDs
			device, rerr := o.model.Create(deviceInfo, nil)
			if rerr != nil {
				lib.Logs.System.Warning().Println(gqlErrors.ERROR_INSERT_DEVICE_IN_DB)
				err = definitionError.NewError(gqlErrors.ERROR_INSERT_DEVICE_IN_DB, nil)
				return
			}
			r = append(r.([]models.Device), device.(models.Device))
		}
	}

	return
}

func (o *Device) readCollegeDependency(id primitive.ObjectID) (err definitionError.GQLError) {
	collegeDependency, rerr := o.collegeDependencyModel.Read(map[string]any{"_id": id}, nil)
	if rerr != nil || len(collegeDependency.([]models.CollegeDependency)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_QUERY_COLLEGE_DEPENDENCY_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_QUERY_COLLEGE_DEPENDENCY_IN_DB, nil)
	}
	return
}

func (o *Device) updateDeviceMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {

	return
}
func (o *Device) deleteDeviceMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {

	return
}
