/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */
package common_filters

import (
	test_helper "github.com/creativesoftwarefdn/weaviate/graphqlapi/test/helper"
	"github.com/graphql-go/graphql"
)

type mockResolver struct {
	test_helper.MockResolver
}

func newMockResolver() *mockResolver {
	// Build a FakeGet.
	fakeGet := &graphql.Field{
		Name:        "SomeAction",
		Description: "Fake Some Action",
		Args: graphql.FieldConfigArgument{
			"where": &graphql.ArgumentConfig{
				Description: "Filter options for the Get search, to convert the data to the filter input",
				Type: graphql.NewInputObject(
					graphql.InputObjectConfig{
						Name:        "WeaviateLocalGetWhereInpObj",
						Fields:      BuildNew("WeaviateLocalGet"),
						Description: "",
					},
				),
			},
		},
		Type: graphql.Int,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			resolver := p.Source.(map[string]interface{})["Resolver"].(*mockResolver)
			filters, err := ExtractFilters(p.Args, p.Info.FieldName)
			if err != nil {
				return nil, err
			}

			result, err := resolver.ReportFilters(filters)
			return result, err
		},
	}

	mocker := &mockResolver{}
	mocker.RootFieldName = "SomeAction"
	mocker.RootField = fakeGet
	mocker.RootObject = map[string]interface{}{"Resolver": mocker}
	return mocker
}

func (m *mockResolver) ReportFilters(filter *LocalFilter) (interface{}, error) {
	args := m.Called(filter)
	return args.Get(0), args.Error(1)
}
