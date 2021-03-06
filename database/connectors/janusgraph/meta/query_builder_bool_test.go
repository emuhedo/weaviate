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
package meta

import (
	"testing"

	gm "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/getmeta"
)

func Test_QueryBuilder_BoolProps(t *testing.T) {
	tests := testCases{
		testCase{
			name: "with only a boolean, with only count",
			inputProps: []gm.MetaProperty{
				gm.MetaProperty{
					Name:                "isCapital",
					StatisticalAnalyses: []gm.StatisticalAnalysis{gm.Count},
				},
			},
			expectedQuery: `
				.union(
					union(
						has("isCapital").count().as("count").project("count").by(select("count"))
					)
					.as("isCapital").project("isCapital").by(select("isCapital"))
				)
			`,
		},

		testCase{
			name: "with count and type",
			inputProps: []gm.MetaProperty{
				gm.MetaProperty{
					Name:                "isCapital",
					StatisticalAnalyses: []gm.StatisticalAnalysis{gm.Count, gm.Type},
				},
			},
			expectedQuery: `
				.union(
					union(
						has("isCapital").count().as("count").project("count").by(select("count"))
					)
					.as("isCapital").project("isCapital").by(select("isCapital"))
				)
			`,
		},

		testCase{
			name: "with only a boolean, with only totalTrue",
			inputProps: []gm.MetaProperty{
				gm.MetaProperty{
					Name:                "isCapital",
					StatisticalAnalyses: []gm.StatisticalAnalysis{gm.TotalTrue},
				},
			},
			expectedQuery: `
				.union(
					union(
						groupCount().by("isCapital")
							.as("boolGroupCount").project("boolGroupCount").by(select("boolGroupCount"))
					)
						.as("isCapital").project("isCapital").by(select("isCapital"))
				)
			`,
		},

		testCase{
			name: "with all boolean props combined",
			inputProps: []gm.MetaProperty{
				gm.MetaProperty{
					Name: "isCapital",
					StatisticalAnalyses: []gm.StatisticalAnalysis{
						gm.Count, gm.TotalTrue, gm.TotalFalse, gm.PercentageTrue, gm.PercentageFalse,
					},
				},
			},
			expectedQuery: `
				.union(
					union(
						has("isCapital").count()
							.as("count").project("count").by(select("count")),
						groupCount().by("isCapital")
							.as("boolGroupCount").project("boolGroupCount").by(select("boolGroupCount"))
					)
						.as("isCapital").project("isCapital").by(select("isCapital"))
				)
			`,
		},
		testCase{
			name: "with only a boolean, with only all true/false props",
			inputProps: []gm.MetaProperty{
				gm.MetaProperty{
					Name: "isCapital",
					StatisticalAnalyses: []gm.StatisticalAnalysis{
						gm.TotalTrue, gm.TotalFalse, gm.PercentageTrue, gm.PercentageFalse,
					},
				},
			},
			expectedQuery: `
				.union(
					union(
						groupCount().by("isCapital")
							.as("boolGroupCount").project("boolGroupCount").by(select("boolGroupCount"))
					)
						.as("isCapital").project("isCapital").by(select("isCapital"))
				)
			`,
		},
	}

	tests.AssertQuery(t, nil)

}

func Test_QueryBuilderWithNamesource(t *testing.T) {
	tests := testCases{
		testCase{
			name: "with only a boolean, with only count",
			inputProps: []gm.MetaProperty{
				gm.MetaProperty{
					Name:                "isCapital",
					StatisticalAnalyses: []gm.StatisticalAnalysis{gm.Count},
				},
			},
			expectedQuery: `.union(` +
				`union(has("prop_20").count().as("count").project("count").by(select("count"))).as("isCapital").project("isCapital").by(select("isCapital"))` +
				`)`,
		},
	}

	tests.AssertQuery(t, &fakeNameSource{})
}
