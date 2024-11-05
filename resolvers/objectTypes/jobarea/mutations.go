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
package jobarea

import (
	"otic/lib"
	gqlErrors "otic/lib/gql_errors"
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (o *JobArea) createJobAreaMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	jobArea, rerr := o.model.Create(input, nil)
	if rerr != nil {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_INSERT_JOB_AREA_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_INSERT_JOB_AREA_IN_DB, nil)
		return
	}
	r = jobArea.(models.JobArea)
	return
}
func (o *JobArea) updateJobAreaMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	jobAreaWhere := map[string]any{"_id": input["_id"].(primitive.ObjectID)}

	dbJobArea, rerr := o.model.Read(jobAreaWhere, nil)
	if rerr != nil || len(dbJobArea.([]models.JobArea)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_QUERY_JOB_AREA_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_QUERY_JOB_AREA_IN_DB, nil)
		return
	}

	updatedJobArea, rerr := o.model.Update(input, jobAreaWhere, nil)
	if rerr != nil {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_UPDATE_JOB_AREA_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_UPDATE_JOB_AREA_IN_DB, nil)
		return
	}
	r = updatedJobArea.([]models.JobArea)[0]

	return
}
