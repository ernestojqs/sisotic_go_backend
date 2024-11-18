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
package task

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

func (o *Task) checkAdminJobTitle(adminID primitive.ObjectID) (err definitionError.GQLError) {
	user, _ := o.userModel.Read(map[string]any{"_id": adminID}, nil)
	if user.([]models.User)[0].JobTitle == primitive.NilObjectID {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_ADMIN_WITHOUT_JOB_TITLE)
		err = definitionError.NewError(gqlErrors.ERROR_ADMIN_WITHOUT_JOB_TITLE, nil)
	}
	return
}
func (o *Task) createTaskMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	session, _ := utils.GetSession(info.SessionID)
	if session.UserRole == string(enums.ROLEENUM_ADMIN) {
		if err = o.checkAdminJobTitle(session.UserID); err != nil {
			return
		}
	}
	user, _ := o.userModel.Read(map[string]any{"_id": session.UserID}, nil)
	jobAreaWhere := bson.M{"jobTitles": bson.M{"$in": bson.A{user.([]models.User)[0].JobTitle}}}
	jobArea, _ := o.jobAreaModel.Read(jobAreaWhere, nil)
	input["jobArea"] = jobArea.([]models.JobArea)[0].Id
	input["autor"] = session.UserID
	r, rerr := o.model.Create(input, nil)
	if rerr != nil {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_INSERT_TASK_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_INSERT_TASK_IN_DB, nil)
	}

	return
}
func (o *Task) updateTaskMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	taskWhere := map[string]any{"_id": input["_id"].(primitive.ObjectID)}
	dbTask, rerr := o.model.Read(taskWhere, nil)
	if rerr != nil || len(dbTask.([]models.Task)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_QUERY_TASK_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_QUERY_TASK_IN_DB, nil)
		return
	}

	updatedTask, rerr := o.model.Update(input, taskWhere, nil)
	if rerr != nil {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_UPDATE_TASK_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_UPDATE_TASK_IN_DB, nil)
		return
	}

	r = updatedTask.([]models.Task)[0]

	return
}

func (o *Task) deleteTaskMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	inputID := info.Args["taskID"].(primitive.ObjectID)
	taskWhere := map[string]any{"_id": inputID}
	dbTask, rerr := o.model.Read(taskWhere, nil)
	if rerr != nil || len(dbTask.([]models.Task)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_QUERY_TASK_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_QUERY_TASK_IN_DB, nil)
		return
	}

	deletedTask, rerr := o.model.Delete(taskWhere, nil)
	if rerr != nil {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_DELETE_TASK_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_DELETE_TASK_IN_DB, nil)
		return
	}

	r = deletedTask.([]models.Task)[0]

	return
}
