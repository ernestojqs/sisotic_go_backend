package jobtitle

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

func (o *JobTitle) edges(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
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
	r = rr.([]models.JobTitle)
	return
}
func (o *JobTitle) jobTitle(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	jobTitleIDs := utils.GetGqlField(info.Resolver, info.Parent)

	if reflect.ValueOf(jobTitleIDs).Kind() == reflect.Slice {
		var manyJobTitles []models.JobTitle
		for _, id := range jobTitleIDs.([]primitive.ObjectID) {
			result, _ := o.model.Read(map[string]interface{}{"_id": id}, nil)
			if len(result.([]models.JobTitle)) != 0 {
				manyJobTitles = append(manyJobTitles, result.([]models.JobTitle)[0])
			}
		}
		r = manyJobTitles
		return
	}

	result, _ := o.model.Read(map[string]interface{}{"_id": jobTitleIDs}, nil)
	if len(result.([]models.JobTitle)) != 0 {
		r = result.([]models.JobTitle)[0]
	}

	return
}
