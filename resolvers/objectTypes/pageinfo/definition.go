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
package pageinfo

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type PageInfo struct {
	model    models.PageInfo
	dbModels map[string]PageInfoModel
}

type PageInfoModel struct {
	object dbutils.ModelInterface
	model  interface{}
}

func NewPageInfo(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {

	dbModels := map[string]PageInfoModel{
		"user":               {&models.User{}, models.User{}},
		"collegeDependency":  {&models.CollegeDependency{}, models.CollegeDependency{}},
		"device":             {&models.Device{}, models.Device{}},
		"jobArea":            {&models.JobArea{}, models.JobArea{}},
		"jobTitle":           {&models.JobTitle{}, models.JobTitle{}},
		"resolverActivity":   {&models.ResolverActivity{}, models.ResolverActivity{}},
		"task":               {&models.Task{}, models.Task{}},
		"technicalDiagnosis": {&models.TechnicalDiagnosis{}, models.TechnicalDiagnosis{}},
	}

	for _, v := range dbModels {
		v.object.Init(v.model, db)
	}
	o = &PageInfo{
		dbModels: dbModels,
	}

	return o
}
func (o *PageInfo) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		r, err = o.edgesQuery(info)
	case "mutation":
		break
	}
	return
}
func (o *PageInfo) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
