package directives

import (
	"reflect"
	"strconv"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PaginateData struct {
	Page         int64
	Split        int64
	Total        int64
	Pages        int64
	Shown        int64
	Overall      int64
	CountOptions *options.CountOptions
	FindOptions  *options.FindOptions
}
type Paginate struct{}

func NewPaginate() resolvers.Directive {
	d := &Paginate{}
	return d
}
func (o *Paginate) Invoke(args map[string]interface{}, typeName string, fieldName string) (r resolvers.DataReturn, err definitionError.GQLError) {
	paginate := &PaginateData{}
	var rTypePage reflect.Kind
	var rTypeSplit reflect.Kind

	if args["page"] != nil {
		rTypePage = reflect.TypeOf(args["page"]).Kind()
	}
	if args["split"] != nil {
		rTypeSplit = reflect.TypeOf(args["split"]).Kind()
	}

	switch rTypePage {
	case reflect.Int:
		paginate.Page = int64(args["page"].(int))
	case reflect.Int32:
		paginate.Page = int64(args["page"].(int32))
	case reflect.Int64:
		paginate.Page = int64(args["page"].(int64))
	case reflect.String:
		i, _ := strconv.Atoi(args["page"].(string))
		paginate.Page = int64(i)
	default:
		paginate.Page = int64(1)
	}
	switch rTypeSplit {
	case reflect.Int:
		paginate.Split = int64(args["split"].(int))
	case reflect.Int32:
		paginate.Split = int64(args["split"].(int32))
	case reflect.Int64:
		paginate.Split = int64(args["split"].(int64))
	case reflect.String:
		i, _ := strconv.Atoi(args["split"].(string))
		paginate.Split = int64(i)
	default:
		paginate.Split = int64(10)
	}
	paginate.CountOptions = new(options.CountOptions)
	paginate.CountOptions.SetSkip(int64((paginate.Page - 1) * paginate.Split))
	paginate.CountOptions.SetLimit(paginate.Split)
	paginate.FindOptions = new(options.FindOptions)
	paginate.FindOptions.SetSkip(int64((paginate.Page - 1) * paginate.Split))
	paginate.FindOptions.SetLimit(paginate.Split)
	r = paginate
	return r, err
}
