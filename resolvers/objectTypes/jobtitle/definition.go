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
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type JobTitle struct {
	model        models.JobTitle
	jobAreaModel models.JobArea
}

func NewJobTitle(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &JobTitle{
		model:        models.JobTitle{},
		jobAreaModel: models.JobArea{},
	}
	o.(*JobTitle).model.Init(models.JobTitle{}, db)
	o.(*JobTitle).jobAreaModel.Init(models.JobArea{}, db)
	return o
}
func (o *JobTitle) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "edges":
			r, err = o.edges(info)
		case "jobTitle":
			r, err = o.jobTitle(info)
		}
	case "mutation":
		switch info.Resolver {
		case "createJobTitle":
			r, err = o.createJobTitleMutation(info)
		case "updateJobTitle":
			r, err = o.updateJobTitleMutation(info)
		case "jobTitle":
			r, err = o.jobTitle(info)
		}
	}
	return
}
func (o *JobTitle) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
