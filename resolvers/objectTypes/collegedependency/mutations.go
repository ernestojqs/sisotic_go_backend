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
package collegedependency

import (
	"otic/lib"
	gqlErrors "otic/lib/gql_errors"
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (o *CollegeDependency) createCollegeDependencyMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	dependency, rerr := o.model.Create(input, nil)
	if rerr != nil {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_INSERT_JOB_AREA_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_INSERT_JOB_AREA_IN_DB, nil)
		return
	}
	r = dependency.(models.CollegeDependency)
	return
}

func (o *CollegeDependency) readCollegeDependency(id primitive.ObjectID) (collegeDependency models.CollegeDependency, err definitionError.GQLError) {
	dependencyWhere := map[string]any{"_id": id}
	result, rerr := o.model.Read(dependencyWhere, nil)
	if rerr != nil || len(result.([]models.CollegeDependency)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_INSERT_COLLEGE_DEPENDENCY_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_INSERT_COLLEGE_DEPENDENCY_IN_DB, nil)
		return
	}
	collegeDependency = result.([]models.CollegeDependency)[0]

	return
}

func (o *CollegeDependency) updateCollegeDependencyMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	dependencyID := input["_id"].(primitive.ObjectID)
	if _, err = o.readCollegeDependency(dependencyID); err != nil {
		return
	}

	dependency, rerr := o.model.Update(input, map[string]any{"_id": dependencyID}, nil)
	if rerr != nil {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_UPDATE_COLLEGE_DEPENDENCY_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_UPDATE_COLLEGE_DEPENDENCY_IN_DB, nil)
		return
	}

	r = dependency.([]models.CollegeDependency)[0]

	return
}
func (o *CollegeDependency) deleteCollegeDependencyMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	dependencyID := info.Args["_id"].(primitive.ObjectID)
	if _, err = o.readCollegeDependency(dependencyID); err != nil {
		return
	}

	deleteWhere := map[string]any{"_id": dependencyID}
	dependency, rerr := o.model.Delete(deleteWhere, nil)
	if rerr != nil {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_DELETE_COLLEGE_DEPENDENCY_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_DELETE_COLLEGE_DEPENDENCY_IN_DB, nil)
		return
	}

	r = dependency.([]models.CollegeDependency)[0]

	return
}
