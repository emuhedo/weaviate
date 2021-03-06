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
 */package batch

import "github.com/creativesoftwarefdn/weaviate/restapi/state"

// New creates all REST api handlers for batching things and actions
func New(appState *state.State) *Batch {
	return &Batch{
		appState: appState,
	}
}

// Batch provides various Handlers around batching things and actions
type Batch struct {
	appState *state.State
}
