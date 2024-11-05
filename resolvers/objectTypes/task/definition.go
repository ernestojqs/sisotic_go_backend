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
	model models.Task
}

func NewTask(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &Task{
		model: models.Task{},
	}
	o.(*Task).model.Init(models.Task{}, db)
	return o
}
func (o *Task) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		break
	case "mutation":
		switch info.Resolver {
		case "createTask":
			r, err = o.createTaskMutation(info)
			break
		case "updateTask":
			r, err = o.updateTaskMutation(info)
			break
		}
		break
	}
	return
}
func (o *Task) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
