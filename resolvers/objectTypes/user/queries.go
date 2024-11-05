package user

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

func (o *User) edges(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
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
	r = rr.([]models.User)
	return
}
func (o *User) user(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {

	if reflect.TypeOf(info.Parent).Name() == "User" {
		r = info.Parent
		return
	}

	userIDs := utils.GetGqlField(info.Resolver, info.Parent)

	if reflect.ValueOf(userIDs).Kind() == reflect.Slice {
		var manyUsers []models.User
		for _, id := range userIDs.([]primitive.ObjectID) {
			result, _ := o.model.Read(map[string]interface{}{"_id": id}, nil)
			if len(result.([]models.User)) != 0 {
				manyUsers = append(manyUsers, result.([]models.User)[0])
			}
		}
		r = manyUsers
		return
	}

	result, _ := o.model.Read(map[string]interface{}{"_id": userIDs}, nil)
	if len(result.([]models.User)) != 0 {
		r = result.([]models.User)[0]
	}

	return
}
