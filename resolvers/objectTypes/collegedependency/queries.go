package collegedependency

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

func (o *CollegeDependency) edges(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
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
	r = rr.([]models.CollegeDependency)
	return
}
func (o *CollegeDependency) collegeDependency(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	collegeDependencyIDs := utils.GetGqlField(info.Resolver, info.Parent)

	if reflect.ValueOf(collegeDependencyIDs).Kind() == reflect.Slice {
		var collegeDependencies []models.CollegeDependency
		for _, id := range collegeDependencyIDs.([]primitive.ObjectID) {
			result, _ := o.model.Read(map[string]interface{}{"_id": id}, nil)
			if len(result.([]models.CollegeDependency)) != 0 {
				collegeDependencies = append(collegeDependencies, result.([]models.CollegeDependency)[0])
			}
		}
		r = collegeDependencies
		return
	}

	result, _ := o.model.Read(map[string]interface{}{"_id": collegeDependencyIDs}, nil)
	if len(result.([]models.CollegeDependency)) != 0 {
		r = result.([]models.CollegeDependency)[0]
	}

	return
}
