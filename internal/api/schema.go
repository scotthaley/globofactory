package api

import (
	"github.com/graphql-go/graphql"
	"github.com/scotthaley/globofactory/pkg/entity"
)

var Schema graphql.Schema

func init() {
	entityType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Entity",
		Description: "Any entity that can be traded",
		Fields: graphql.Fields{
			"code": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if e, ok := p.Source.(entity.Entity); ok {
						return e.Code, nil
					}
					return nil, nil
				},
			},
			"display": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if e, ok := p.Source.(entity.Entity); ok {
						return e.Display, nil
					}
					return nil, nil
				},
			},
		},
	})
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"entity": &graphql.Field{
				Type: entityType,
				Args: graphql.FieldConfigArgument{
					"code": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					code := p.Args["code"].(string)
					return entity.GetEntity(code), nil
				},
			},
		},
	})

	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
