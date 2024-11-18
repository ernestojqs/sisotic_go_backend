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
	"time"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (o *Device) checkAdminJobTitle(adminID primitive.ObjectID) (err definitionError.GQLError) {
	user, _ := o.userModel.Read(map[string]any{"_id": adminID}, nil)
	if user.([]models.User)[0].JobTitle == primitive.NilObjectID {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_ADMIN_WITHOUT_JOB_TITLE)
		err = definitionError.NewError(gqlErrors.ERROR_ADMIN_WITHOUT_JOB_TITLE, nil)
	}
	return
}

func (o *Device) createDevicesMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	sess, _ := utils.GetSession(info.SessionID)
	if sess.UserRole == string(enums.ROLEENUM_ADMIN) {
		if err = o.checkAdminJobTitle(sess.UserID); err != nil {
			return
		}
	}
	input := info.Args["input"].(map[string]any)
	r = []models.Device{}
	collegeDependencyID := input["collegeDependency"].(primitive.ObjectID)
	if err = o.readCollegeDependency(collegeDependencyID); err != nil {
		return
	}

	placeOfCare := input["placeOfCare"].(string)
	if placeOfCare == string(enums.PLACEOFCAREENUM_WORKSHOP) {
		if input["transferEvidence"] == nil || input["transferEvidence"].(map[string]any)["id_number"] == nil {
			lib.Logs.System.Warning().Println(gqlErrors.ERROR_TRANSFER__EVIDENCE_REQUERIDED)
			err = definitionError.NewError(gqlErrors.ERROR_TRANSFER__EVIDENCE_REQUERIDED, nil)
			return
		}
		input["transferEvidence"].(map[string]any)["transferType"] = string(enums.TRANSFERTYPEENUM_ENTRY)
	}

	groupID := utils.GenerateTokenFromUUID(12, true)
	for _, v := range input["deviceInfo"].([]any) {
		deviceInfo := v.(map[string]any)
		deviceInfo["groupID"] = groupID
		deviceInfo["receiverUser"] = sess.UserID
		deviceInfo["collegeDependency"] = collegeDependencyID
		deviceInfo["placeOfCare"] = placeOfCare
		if placeOfCare == string(enums.PLACEOFCAREENUM_LOCAL) {
			deviceInfo["currentSupportStatus"] = string(enums.SUPPORTSTATUSENUM_DELIVERED)
			if deviceInfo["isSupport"] == true && ((deviceInfo["localTechnicalDiagnosis"] == nil) || (deviceInfo["localTechnicalDiagnosis"] != nil && len(deviceInfo["localTechnicalDiagnosis"].([]any)) == 0)) {
				lib.Logs.System.Warning().Println(gqlErrors.ERROR_TECHNICAL_DIAGNOSIS_REQUERIDED)
				err = definitionError.NewError(gqlErrors.ERROR_TECHNICAL_DIAGNOSIS_REQUERIDED, nil)
				return
			}
			deviceInfo["technicalDiagnosis"], err = o.createLocalInfo(deviceInfo["localTechnicalDiagnosis"].([]any), sess.UserID)
			if err != nil {
				return
			}
		}
		if placeOfCare == string(enums.PLACEOFCAREENUM_WORKSHOP) {
			deviceInfo["currentSupportStatus"] = string(enums.SUPPORTSTATUSENUM_PENDING)
			if deviceInfo["isSupport"] == false {
				deviceInfo["currentSupportStatus"] = string(enums.SUPPORTSTATUSENUM_COMPLETED)
			}
			deviceInfo["transferEvidences"] = []any{input["transferEvidence"]}
		}
		deviceInfo["supportStatusDetails"] = []any{map[string]any{
			"type":       deviceInfo["currentSupportStatus"].(string),
			"date":       time.Now().Unix(),
			"recordUser": sess.UserID,
		}}
		device, rerr := o.model.Create(deviceInfo, nil)
		if rerr != nil {
			lib.Logs.System.Warning().Println(gqlErrors.ERROR_INSERT_DEVICE_IN_DB)
			err = definitionError.NewError(gqlErrors.ERROR_INSERT_DEVICE_IN_DB, nil)
			return
		}
		r = append(r.([]models.Device), device.(models.Device))
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

func (o *Device) checkTransferNumber() (err definitionError.GQLError) {
	return
}

func (o *Device) readDevice(where map[string]any) (device models.Device, err definitionError.GQLError) {
	result, rerr := o.model.Read(where, nil)
	if rerr != nil || len(result.([]models.Device)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_QUERY_DEVICE_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_QUERY_DEVICE_IN_DB, nil)
		return
	}
	return result.([]models.Device)[0], nil
}

func (o *Device) diagnosticsAreResolved(diagnosticIDs []primitive.ObjectID) (r bool) {
	if len(diagnosticIDs) != 0 {
		for _, v := range diagnosticIDs {
			where := bson.M{"_id": v, "resolverActivities": bson.M{"$size": 0}}
			if result, _ := o.technicalDiagnosisModel.Read(where, nil); len(result.([]models.TechnicalDiagnosis)) != 0 {
				return
			}
		}
		r = true
	}
	return
}

func (o *Device) updateDeviceMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	deviceID := input["_id"].(primitive.ObjectID)
	where := bson.M{"_id": deviceID}
	dbDevice, err := o.readDevice(where)
	if err != nil {
		return
	}

	if dbDevice.CurrentSupportStatus == enums.SUPPORTSTATUSENUM_COMPLETED || dbDevice.CurrentSupportStatus == enums.SUPPORTSTATUSENUM_DELIVERED {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_CANNOT_UPDATE_DEVICE_INFO)
		err = definitionError.NewError(gqlErrors.ERROR_CANNOT_UPDATE_DEVICE_INFO, nil)
		return
	}

	if input["collegeDependency"] != nil {
		if err = o.readCollegeDependency(input["collegeDependency"].(primitive.ObjectID)); err != nil {
			return
		}
	}

	if input["observations"] == nil || (input["observations"] != nil && len(input["observations"].([]any)) == 0) {
		input["observations"] = utils.ParseArayDBObj(dbDevice.Observations)
	}

	if input["currentSupportStatus"] != nil && input["currentSupportStatus"].(string) != "" {
		inputSupportStatus := input["currentSupportStatus"].(string)
		if dbDevice.CurrentSupportStatus == enums.SUPPORTSTATUSENUM_PROCESS && inputSupportStatus == enums.SUPPORTSTATUSENUM_PENDING {
			lib.Logs.System.Warning().Println(gqlErrors.ERROR_BAD_STATUS_ORDER_1)
			err = definitionError.NewError(gqlErrors.ERROR_BAD_STATUS_ORDER_1, nil)
			return
		}

		if dbDevice.CurrentSupportStatus != enums.SUPPORTSTATUSENUM_PROCESS && inputSupportStatus == enums.SUPPORTSTATUSENUM_COMPLETED {
			lib.Logs.System.Warning().Println(gqlErrors.ERROR_BAD_STATUS_ORDER_2)
			err = definitionError.NewError(gqlErrors.ERROR_BAD_STATUS_ORDER_2, nil)
			return
		}

		if inputSupportStatus == enums.SUPPORTSTATUSENUM_COMPLETED && !o.diagnosticsAreResolved(dbDevice.TechnicalDiagnosis) {
			lib.Logs.System.Warning().Println(gqlErrors.ERROR_DIAGNOSIS_NOT_RESOLVED)
			err = definitionError.NewError(gqlErrors.ERROR_DIAGNOSIS_NOT_RESOLVED, nil)
			return
		}

		if inputSupportStatus == enums.SUPPORTSTATUSENUM_DELIVERED {
			lib.Logs.System.Warning().Println(gqlErrors.ERROR_BAD_STATUS_ORDER_4)
			err = definitionError.NewError(gqlErrors.ERROR_BAD_STATUS_ORDER_4, nil)
			return
		}

		if string(dbDevice.CurrentSupportStatus) == inputSupportStatus {
			lib.Logs.System.Warning().Println(gqlErrors.ERROR_BAD_STATUS_ORDER_3)
			err = definitionError.NewError(gqlErrors.ERROR_BAD_STATUS_ORDER_3, nil)
			return
		}
		recordingStatus := utils.ParseArayDBObj(dbDevice.SupportStatusDetails)
		sess, _ := utils.GetSession(info.SessionID)
		recordingStatus = append(recordingStatus, map[string]any{
			"type":       inputSupportStatus,
			"date":       time.Now().Unix(),
			"recordUser": sess.UserID,
		},
		)
		input["supportStatusDetails"] = recordingStatus
	}
	r, err = o.updateDevice(deviceID, input)
	return
}

func (o *Device) updateDevice(id primitive.ObjectID, inputValues map[string]any) (device models.Device, err definitionError.GQLError) {
	result, rerr := o.model.Update(inputValues, map[string]any{"_id": id}, nil)
	if rerr != nil || len(result.([]models.Device)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_UPDATE_DEVICE_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_UPDATE_DEVICE_IN_DB, nil)
		return
	}
	return result.([]models.Device)[0], nil
}

func (o *Device) deleteDeviceMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {

	return
}

func (o *Device) createTask(resolvedDate int64, description string, userID primitive.ObjectID) {
	user, _ := o.userModel.Read(map[string]any{"_id": userID}, nil)
	jobAreaWhere := bson.M{"jobTitles": bson.M{"$in": bson.A{user.([]models.User)[0].JobTitle}}}
	jobArea, _ := o.jobAreaModel.Read(jobAreaWhere, nil)
	taskInput := map[string]any{
		"description":  description,
		"jobArea":      jobArea.([]models.JobArea)[0].Id,
		"autor":        userID,
		"resolvedDate": resolvedDate,
	}
	o.taskModel.Create(taskInput, nil)
}

func (o *Device) createLocalInfo(localTechnicalDiagnosis []any, loggedUserID primitive.ObjectID) (techDiagnosisIDs []primitive.ObjectID, err definitionError.GQLError) {
	for _, v := range localTechnicalDiagnosis {
		techDiagnosis := v.(map[string]any)
		if (techDiagnosis["localResolverActivities"] == nil) || (techDiagnosis["localResolverActivities"] != nil && len(techDiagnosis["localResolverActivities"].([]any)) == 0) {
			lib.Logs.System.Warning().Println(gqlErrors.ERROR_TECHNICAL_DIAGNOSIS_REQUERIDED)
			err = definitionError.NewError(gqlErrors.ERROR_TECHNICAL_DIAGNOSIS_REQUERIDED, nil)
			return
		}
		var resolverActivityIDs []primitive.ObjectID
		for _, v := range techDiagnosis["localResolverActivities"].([]any) {
			resolverActivityInput := v.(map[string]any)
			resolverActivityInput["resolverUser"] = loggedUserID
			result, rerr := o.resolverActivityModel.Create(resolverActivityInput, nil)
			if rerr != nil {
				lib.Logs.System.Warning().Println(gqlErrors.ERROR_INSERT_RESOLVER_ACTIVITY_IN_DB)
				err = definitionError.NewError(gqlErrors.ERROR_INSERT_RESOLVER_ACTIVITY_IN_DB, nil)
				return
			}
			o.createTask(resolverActivityInput["resolvedDate"].(int64), resolverActivityInput["description"].(string), loggedUserID)
			resolverActivityIDs = append(resolverActivityIDs, result.(models.ResolverActivity).Id)
		}
		techDiagnosis["resolverActivities"] = resolverActivityIDs
		techDiagnosis["diagnosticUser"] = loggedUserID
		createdDiagnosis, rerr := o.technicalDiagnosisModel.Create(techDiagnosis, nil)
		if rerr != nil {
			lib.Logs.System.Warning().Println(gqlErrors.ERROR_INSERT_TECHNICAL_DIAGNOSIS_IN_DB)
			err = definitionError.NewError(gqlErrors.ERROR_INSERT_TECHNICAL_DIAGNOSIS_IN_DB, nil)
			return
		}
		techDiagnosisIDs = append(techDiagnosisIDs, createdDiagnosis.(models.TechnicalDiagnosis).Id)
	}
	return
}

func (o *Device) deliveryDevices(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	var updatedDevices, dbDevices []models.Device
	input["transferEvidence"].(map[string]any)["transferType"] = string(enums.TRANSFERTYPEENUM_EXIT)
	sess, _ := utils.GetSession(info.SessionID)
	var groupID string
	for i, deviceID := range input["deviceIDs"].([]any) {
		dbDevice, rerr := o.readDevice(map[string]any{"_id": deviceID.(primitive.ObjectID)})
		if rerr != nil {
			o.updateDeviceToPreviousValues(dbDevices)
			err = rerr
			return
		}

		if i == 0 {
			groupID = dbDevice.GroupID
		}

		if dbDevice.GroupID != groupID {
			lib.Logs.System.Warning().Println(gqlErrors.ERROR_DEVICES_FROM_DIFFERENT_GROUP_ID)
			o.updateDeviceToPreviousValues(dbDevices)
			err = definitionError.NewError(gqlErrors.ERROR_DEVICES_FROM_DIFFERENT_GROUP_ID, nil)
			return
		}
		if dbDevice.CurrentSupportStatus != enums.SUPPORTSTATUSENUM_COMPLETED {
			lib.Logs.System.Warning().Println(gqlErrors.ERROR_BAD_STATUS_ORDER_5)
			o.updateDeviceToPreviousValues(dbDevices)
			err = definitionError.NewError(gqlErrors.ERROR_BAD_STATUS_ORDER_5, nil)
			return
		}
		dbDevices = append(dbDevices, dbDevice)
		tranferEvidences := []any{
			utils.ParseDBObj(dbDevice.TransferEvidences[0]),
			input["transferEvidence"],
		}
		recordingStatus := utils.ParseArayDBObj(dbDevice.SupportStatusDetails)
		recordingStatus = append(recordingStatus, map[string]any{
			"type":       string(enums.SUPPORTSTATUSENUM_DELIVERED),
			"date":       time.Now().Unix(),
			"recordUser": sess.UserID,
		})
		deviceInput := map[string]any{
			"transferEvidences":    tranferEvidences,
			"currentSupportStatus": string(enums.SUPPORTSTATUSENUM_DELIVERED),
			"supportStatusDetails": recordingStatus,
		}
		result, _ := o.updateDevice(dbDevice.Id, deviceInput)
		updatedDevices = append(updatedDevices, result)
	}
	r = updatedDevices
	return
}

func (o *Device) updateDeviceToPreviousValues(devices []models.Device) {
	for _, v := range devices {
		input := map[string]any{
			"transferEvidences":    utils.ParseArayDBObj(v.TransferEvidences),
			"currentSupportStatus": string(v.CurrentSupportStatus),
			"supportStatusDetails": utils.ParseArayDBObj(v.SupportStatusDetails),
		}
		o.updateDevice(v.Id, input)
	}
}
