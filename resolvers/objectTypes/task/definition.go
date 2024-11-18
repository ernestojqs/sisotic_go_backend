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
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type Task struct {
	model        models.Task
	jobAreaModel models.JobArea
	userModel    models.User
}

func NewTask(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &Task{
		model:        models.Task{},
		jobAreaModel: models.JobArea{},
		userModel:    models.User{},
	}
	o.(*Task).model.Init(models.Task{}, db)
	o.(*Task).jobAreaModel.Init(models.JobArea{}, db)
	o.(*Task).userModel.Init(models.User{}, db)
	return o
}
func (o *Task) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "edges":
			r, err = o.edges(info)
		}
	case "mutation":
		switch info.Resolver {
		case "createTask":
			r, err = o.createTaskMutation(info)
		case "updateTask":
			r, err = o.updateTaskMutation(info)
		case "deleteTask":
			r, err = o.deleteTaskMutation(info)
		}
	}
	return
}
func (o *Task) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
