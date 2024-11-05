package jobarea

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

func (o *JobArea) edges(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
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
	r = rr.([]models.JobArea)
	return
}
func (o *JobArea) jobArea(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	jobAreaIDs := utils.GetGqlField(info.Resolver, info.Parent)

	if reflect.ValueOf(jobAreaIDs).Kind() == reflect.Slice {
		var manyJobAreas []models.JobArea
		for _, id := range jobAreaIDs.([]primitive.ObjectID) {
			result, _ := o.model.Read(map[string]interface{}{"_id": id}, nil)
			if len(result.([]models.JobArea)) != 0 {
				manyJobAreas = append(manyJobAreas, result.([]models.JobArea)[0])
			}
		}
		r = manyJobAreas
		return
	}

	result, _ := o.model.Read(map[string]interface{}{"_id": jobAreaIDs}, nil)
	if len(result.([]models.JobArea)) != 0 {
		r = result.([]models.JobArea)[0]
	}

	return
}
