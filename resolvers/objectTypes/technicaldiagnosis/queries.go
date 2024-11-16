package technicaldiagnosis

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

func (o *TechnicalDiagnosis) edges(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
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
	r = rr.([]models.TechnicalDiagnosis)
	return
}
func (o *TechnicalDiagnosis) technicalDiagnosis(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	technicalDiagnosisIDs := utils.GetGqlField(info.Resolver, info.Parent)

	if reflect.ValueOf(technicalDiagnosisIDs).Kind() == reflect.Slice {
		var manyTechnicalDiagnosis []models.TechnicalDiagnosis
		for _, id := range technicalDiagnosisIDs.([]primitive.ObjectID) {
			result, _ := o.model.Read(map[string]interface{}{"_id": id}, nil)
			if len(result.([]models.TechnicalDiagnosis)) != 0 {
				manyTechnicalDiagnosis = append(manyTechnicalDiagnosis, result.([]models.TechnicalDiagnosis)[0])
			}
		}
		r = manyTechnicalDiagnosis
		return
	}

	result, _ := o.model.Read(map[string]interface{}{"_id": technicalDiagnosisIDs}, nil)
	if len(result.([]models.TechnicalDiagnosis)) != 0 {
		r = result.([]models.TechnicalDiagnosis)[0]
	}

	return
}
