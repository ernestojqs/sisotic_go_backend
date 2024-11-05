package lib

import (
	"otic/models"
	"otic/resolvers/directives"

	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Edge struct {
	Edges     string          `gql:"name=edges"`
	PageInfo  models.PageInfo `gql:"name=pageInfo"`
	FindOpts  []*options.FindOptions
	CountOpts []*options.CountOptions
	Where     bson.D
}

func NewEdge(info resolvers.ResolverInfo, model string) (r Edge) {
	r.Where = bson.D{}
	if info.Directives["paginate"] != nil {
		r.FindOpts = append(r.FindOpts, info.Directives["paginate"].(*directives.PaginateData).FindOptions)
		r.CountOpts = append(r.CountOpts, info.Directives["paginate"].(*directives.PaginateData).CountOptions)
		r.PageInfo.Total = info.Directives["paginate"].(*directives.PaginateData).Total
		r.PageInfo.Shown = info.Directives["paginate"].(*directives.PaginateData).Shown
		r.PageInfo.Page = info.Directives["paginate"].(*directives.PaginateData).Page
		r.PageInfo.Pages = info.Directives["paginate"].(*directives.PaginateData).Pages
		r.PageInfo.Split = info.Directives["paginate"].(*directives.PaginateData).Split
	}
	if info.Directives["search"] != nil {
		r.Where = info.Directives["search"].(bson.D)
	}
	r.Edges = model
	return
}
