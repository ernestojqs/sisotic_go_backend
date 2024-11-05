package pageinfo

import (
	"math"
	"otic/lib"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
)

func (o *PageInfo) edgesQuery(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	var rerr error
	var pages, shown, total, overall int64
	edge := o.dbModels[info.Parent.(lib.Edge).Edges]
	parent := info.Parent.(lib.Edge)
	if edge.object == nil {
		lib.Logs.System.Fatal().Println("model " + info.Parent.(lib.Edge).Edges + " not set into PageInfo.model")
		err = definitionError.NewFatal("model "+info.Parent.(lib.Edge).Edges+" not set into PageInfo.model", nil)
		return
	}

	total, rerr = edge.object.Count(parent.Where, nil)
	if rerr != nil {
		lib.Logs.System.Fatal().Println(rerr.Error())
		err = definitionError.NewFatal(rerr.Error(), nil)
		return
	}
	shown, rerr = edge.object.Count(parent.Where, parent.CountOpts)
	if rerr != nil {
		lib.Logs.System.Fatal().Println(rerr.Error())
		err = definitionError.NewFatal(rerr.Error(), nil)
		return
	}
	overall, rerr = edge.object.Count(parent.Where, nil)
	if rerr != nil {
		lib.Logs.System.Fatal().Println(rerr.Error())
		err = definitionError.NewFatal(rerr.Error(), nil)
		return
	}
	pages = int64(math.Ceil(float64(total) / float64(parent.PageInfo.Split)))
	o.model.Pages = pages
	o.model.Shown = shown
	o.model.Total = total
	o.model.Overall = overall
	o.model.Page = parent.PageInfo.Page
	o.model.Split = parent.PageInfo.Split
	r = o.model
	return
}
