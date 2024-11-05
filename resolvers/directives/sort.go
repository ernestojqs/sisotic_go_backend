package directives

import (
	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SortData struct {
	Field string
	Order int
}
type Sort struct {
}

func NewSort() resolvers.Directive {
	d := &Sort{}
	return d
}
func (o *Sort) Invoke(args map[string]interface{}, typeName string, fieldName string) (r resolvers.DataReturn, err definitionError.GQLError) {
	sort := bson.D{}
	for _, v := range args["input"].([]interface{}) {
		x := v.(map[string]interface{})
		order := 1
		switch x["order"].(string) {
		case "asc":
			order = 1
		case "desc":
			order = -1
		}
		sort = append(sort, bson.E{Key: x["field"].(string), Value: order})
	}
	return options.Find().SetSort(sort), err
}