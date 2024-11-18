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
package resolveractivity

import (
	"otic/lib"
	gqlErrors "otic/lib/gql_errors"
	"otic/lib/utils"
	"otic/models"
	"otic/models/enums"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (o *ResolverActivity) readDiagnosis(id primitive.ObjectID) (technicalDiagnosis models.TechnicalDiagnosis, err definitionError.GQLError) {
	result, rerr := o.technicalDiagnosis.Read(map[string]any{"_id": id}, nil)
	if rerr != nil || len(result.([]models.TechnicalDiagnosis)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_QUERY_TECHNICAL_DIAGNOSIS_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_QUERY_TECHNICAL_DIAGNOSIS_IN_DB, nil)
		return
	}
	deviceWhere := bson.M{
		"technicalDiagnosis":   bson.M{"$in": bson.A{id}},
		"currentSupportStatus": enums.SUPPORTSTATUSENUM_PROCESS,
		"placeOfCare":          enums.PLACEOFCAREENUM_WORKSHOP,
		"isSupport":            true,
	}
	device, _ := o.deviceModel.Read(deviceWhere, nil)
	if len(device.([]models.Device)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_DEVICE_IS_NOT_IN_PROCESS)
		err = definitionError.NewError(gqlErrors.ERROR_DEVICE_IS_NOT_IN_PROCESS, nil)
		return
	}

	return result.([]models.TechnicalDiagnosis)[0], nil
}

func (o *ResolverActivity) updateDiagnosis(id primitive.ObjectID, inputValues map[string]any) (technicalDiagnosis models.TechnicalDiagnosis, err definitionError.GQLError) {
	result, rerr := o.technicalDiagnosis.Update(inputValues, map[string]any{"_id": id}, nil)
	if rerr != nil || len(result.([]models.TechnicalDiagnosis)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_UPDATE_TECHNICAL_DIAGNOSIS_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_UPDATE_TECHNICAL_DIAGNOSIS_IN_DB, nil)
		return
	}
	return result.([]models.TechnicalDiagnosis)[0], nil
}

func (o *ResolverActivity) createTask(description string, userID primitive.ObjectID) {
	user, _ := o.userModel.Read(map[string]any{"_id": userID}, nil)
	jobAreaWhere := bson.M{"jobTitles": bson.M{"$in": bson.A{user.([]models.User)[0].JobTitle}}}
	jobArea, _ := o.jobAreaModel.Read(jobAreaWhere, nil)
	taskInput := map[string]any{
		"description": description,
		"jobArea":     jobArea.([]models.JobArea)[0].Id,
		"autor":       userID,
	}
	o.taskModel.Create(taskInput, nil)
}

func (o *ResolverActivity) checkAdminJobTitle(adminID primitive.ObjectID) (err definitionError.GQLError) {
	user, _ := o.userModel.Read(map[string]any{"_id": adminID}, nil)
	if user.([]models.User)[0].JobTitle == primitive.NilObjectID {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_ADMIN_WITHOUT_JOB_TITLE)
		err = definitionError.NewError(gqlErrors.ERROR_ADMIN_WITHOUT_JOB_TITLE, nil)
	}
	return
}

func (o *ResolverActivity) createResolverActivityMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	diagnosisID := input["technicalDiagnosisID"].(primitive.ObjectID)
	sess, _ := utils.GetSession(info.SessionID)
	if sess.UserRole == string(enums.ROLEENUM_ADMIN) {
		if err = o.checkAdminJobTitle(sess.UserID); err != nil {
			return
		}
	}
	input["resolverUser"] = sess.UserID
	dbDiagnosis, err := o.readDiagnosis(diagnosisID)
	if err != nil {
		return
	}

	r, rerr := o.model.Create(input, nil)
	if rerr != nil {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_INSERT_RESOLVER_ACTIVITY_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_INSERT_RESOLVER_ACTIVITY_IN_DB, nil)
		return
	}

	newActivities := append(dbDiagnosis.ResolverActivities, r.(models.ResolverActivity).Id)
	o.updateDiagnosis(diagnosisID, map[string]any{"resolverActivities": newActivities})
	o.createTask(input["description"].(string), sess.UserID)
	return
}
