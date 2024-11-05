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
package jobtitle

import (
	"otic/lib"
	gqlErrors "otic/lib/gql_errors"
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (o *JobTitle) createJobTitleMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	jobAreaID := input["jobArea"].(primitive.ObjectID)

	jobAreaWhere := map[string]any{"_id": jobAreaID}
	dbJobArea, rerr := o.jobAreaModel.Read(jobAreaWhere, nil)
	if rerr != nil || len(dbJobArea.([]models.JobArea)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_QUERY_JOB_AREA_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_QUERY_JOB_AREA_IN_DB, nil)
		return
	}

	jobTitle, rerr := o.model.Create(input, nil)
	if rerr != nil {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_INSERT_JOB_TITLE_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_INSERT_JOB_TITLE_IN_DB, nil)
		return
	}
	r = jobTitle.(models.JobTitle)

	newJobTitles := append(dbJobArea.([]models.JobArea)[0].JobTitles, jobTitle.(models.JobTitle).Id)
	jobAreaInput := map[string]any{"jobTitles": newJobTitles}
	o.jobAreaModel.Update(jobAreaInput, jobAreaWhere, nil)

	return
}
func (o *JobTitle) updateJobTitleMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	jobTitleWhere := map[string]any{"_id": input["_id"].(primitive.ObjectID)}

	dbJobTitle, rerr := o.model.Read(jobTitleWhere, nil)
	if rerr != nil || len(dbJobTitle.([]models.JobTitle)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_QUERY_JOB_TITLE_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_QUERY_JOB_TITLE_IN_DB, nil)
		return
	}

	updatedJobTitle, rerr := o.model.Update(input, jobTitleWhere, nil)
	if rerr != nil {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_UPDATE_JOB_TITLE_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_UPDATE_JOB_TITLE_IN_DB, nil)
		return
	}
	r = updatedJobTitle.([]models.JobTitle)[0]
	return
}
