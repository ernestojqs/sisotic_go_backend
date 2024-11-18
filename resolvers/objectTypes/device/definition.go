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
package device

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type Device struct {
	model                   models.Device
	resolverActivityModel   models.ResolverActivity
	technicalDiagnosisModel models.TechnicalDiagnosis
	collegeDependencyModel  models.CollegeDependency
	taskModel               models.Task
	userModel               models.User
	jobAreaModel            models.JobArea
}

func NewDevice(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &Device{
		model:                   models.Device{},
		resolverActivityModel:   models.ResolverActivity{},
		technicalDiagnosisModel: models.TechnicalDiagnosis{},
		collegeDependencyModel:  models.CollegeDependency{},
		taskModel:               models.Task{},
	}
	o.(*Device).model.Init(models.Device{}, db)
	o.(*Device).resolverActivityModel.Init(models.ResolverActivity{}, db)
	o.(*Device).technicalDiagnosisModel.Init(models.TechnicalDiagnosis{}, db)
	o.(*Device).collegeDependencyModel.Init(models.CollegeDependency{}, db)
	o.(*Device).taskModel.Init(models.Task{}, db)
	o.(*Device).userModel.Init(models.User{}, db)
	o.(*Device).jobAreaModel.Init(models.JobArea{}, db)
	return o
}
func (o *Device) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "edges":
			r, err = o.edges(info)
		}
	case "mutation":
		switch info.Resolver {
		case "createDevices":
			r, err = o.createDevicesMutation(info)
		case "updateDevice":
			r, err = o.updateDeviceMutation(info)
		case "deleteDevice":
			r, err = o.deleteDeviceMutation(info)
		case "deliveryDevices":
			r, err = o.deliveryDevices(info)
		}
	}
	return
}
func (o *Device) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
