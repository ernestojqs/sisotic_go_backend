package task

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

func (o *Task) edges(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
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
	r = rr.([]models.Task)
	return
}
func (o *Task) task(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	taskIDs := utils.GetGqlField(info.Resolver, info.Parent)

	if reflect.ValueOf(taskIDs).Kind() == reflect.Slice {
		var manyTasks []models.Task
		for _, id := range taskIDs.([]primitive.ObjectID) {
			result, _ := o.model.Read(map[string]interface{}{"_id": id}, nil)
			if len(result.([]models.Task)) != 0 {
				manyTasks = append(manyTasks, result.([]models.Task)[0])
			}
		}
		r = manyTasks
		return
	}

	result, _ := o.model.Read(map[string]interface{}{"_id": taskIDs}, nil)
	if len(result.([]models.Task)) != 0 {
		r = result.([]models.Task)[0]
	}

	return
}
