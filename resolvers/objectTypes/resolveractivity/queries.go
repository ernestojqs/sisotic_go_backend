package resolveractivity

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

func (o *ResolverActivity) edges(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
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
	r = rr.([]models.ResolverActivity)
	return
}
func (o *ResolverActivity) resolverActivity(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	resolverActivityIDs := utils.GetGqlField(info.Resolver, info.Parent)

	if reflect.ValueOf(resolverActivityIDs).Kind() == reflect.Slice {
		var DanyResolverActivitys []models.ResolverActivity
		for _, id := range resolverActivityIDs.([]primitive.ObjectID) {
			result, _ := o.model.Read(map[string]interface{}{"_id": id}, nil)
			if len(result.([]models.ResolverActivity)) != 0 {
				DanyResolverActivitys = append(DanyResolverActivitys, result.([]models.ResolverActivity)[0])
			}
		}
		r = DanyResolverActivitys
		return
	}

	result, _ := o.model.Read(map[string]interface{}{"_id": resolverActivityIDs}, nil)
	if len(result.([]models.ResolverActivity)) != 0 {
		r = result.([]models.ResolverActivity)[0]
	}

	return
}
