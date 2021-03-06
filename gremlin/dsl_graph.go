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
package gremlin

import (
	"fmt"
)

type Graph struct{}

// This is the starting point for building queries.
var G Graph

func (g *Graph) V() *Query {
	q := Query{query: "g.V()"}
	return &q
}

func (g *Graph) E() *Query {
	q := Query{query: "g.E()"}
	return &q
}

func Current() *Query {
	return &Query{query: "__"}
}

func (g *Graph) AddV(label string) *Query {
	query := fmt.Sprintf(`g.addV("%s")`, EscapeString(label))
	q := Query{query: query}
	return &q
}

func (g *Graph) AddE(label string) *Query {
	query := fmt.Sprintf(`g.addE("%s")`, EscapeString(label))
	q := Query{query: query}
	return &q
}
