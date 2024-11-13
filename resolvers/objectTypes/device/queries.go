package device

import (
	"otic/lib"
	"otic/lib/utils"
	"otic/models"
	"reflect"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (o *Device) edges(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	var opts []*options.FindOptions
	if info.Parent.(lib.Edge).FindOpts != nil {
		opts = append(opts, info.Parent.(lib.Edge).FindOpts...)
	}
	rr, rerr := o.model.Read(info.Parent.(lib.Edge).Where, opts)
	if rerr != nil {
		lib.Logs.System.Fatal().Println(rerr.Error())
		err = definitionError.NewFatal(rerr.Error(), nil)
		return
	}
	r = rr.([]models.Device)
	return
}
func (o *Device) device(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	deviceIDs := utils.GetGqlField(info.Resolver, info.Parent)

	if reflect.ValueOf(deviceIDs).Kind() == reflect.Slice {
		var Danydevices []models.Device
		for _, id := range deviceIDs.([]primitive.ObjectID) {
			result, _ := o.model.Read(map[string]interface{}{"_id": id}, nil)
			if len(result.([]models.Device)) != 0 {
				Danydevices = append(Danydevices, result.([]models.Device)[0])
			}
		}
		r = Danydevices
		return
	}

	result, _ := o.model.Read(map[string]interface{}{"_id": deviceIDs}, nil)
	if len(result.([]models.Device)) != 0 {
		r = result.([]models.Device)[0]
	}

	return
}
