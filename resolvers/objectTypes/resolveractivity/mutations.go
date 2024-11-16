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

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (o *ResolverActivity) readDiagnosis(id primitive.ObjectID) (technicalDiagnosis models.TechnicalDiagnosis, err definitionError.GQLError) {
	result, rerr := o.model.Read(map[string]any{"_id": id}, nil)
	if rerr != nil || len(result.([]models.TechnicalDiagnosis)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_QUERY_TECHNICAL_DIAGNOSIS_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_QUERY_TECHNICAL_DIAGNOSIS_IN_DB, nil)
		return
	}
	return result.([]models.TechnicalDiagnosis)[0], nil
}

func (o *ResolverActivity) updateDiagnosis(id primitive.ObjectID, inputValues map[string]any) (technicalDiagnosis models.TechnicalDiagnosis, err definitionError.GQLError) {
	result, rerr := o.model.Update(inputValues, map[string]any{"_id": id}, nil)
	if rerr != nil || len(result.([]models.TechnicalDiagnosis)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_UPDATE_TECHNICAL_DIAGNOSIS_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_UPDATE_TECHNICAL_DIAGNOSIS_IN_DB, nil)
		return
	}
	return result.([]models.TechnicalDiagnosis)[0], nil
}

func (o *ResolverActivity) createResolverActivityMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	diagnosisID := input["technicalDiagnosisID"].(primitive.ObjectID)
	sess, _ := utils.GetSession(info.SessionID)
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

	return
}
