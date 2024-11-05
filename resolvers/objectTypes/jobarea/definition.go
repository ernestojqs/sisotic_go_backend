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
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type JobArea struct {
	model models.JobArea
}

func NewJobArea(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &JobArea{
		model: models.JobArea{},
	}
	o.(*JobArea).model.Init(models.JobArea{}, db)
	return o
}
func (o *JobArea) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "edges":
			r, err = o.edges(info)
		case "jobArea":
			r, err = o.jobArea(info)
		}
	case "mutation":
		switch info.Resolver {
		case "createJobArea":
			r, err = o.createJobAreaMutation(info)
		case "updateJobArea":
			r, err = o.updateJobAreaMutation(info)
		case "jobArea":
			r, err = o.jobArea(info)
		}
	}
	return
}
func (o *JobArea) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
