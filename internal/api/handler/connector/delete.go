// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package connector

import (
	"net/http"

	"github.com/harness/gitness/internal/api/controller/connector"
	"github.com/harness/gitness/internal/api/render"
	"github.com/harness/gitness/internal/api/request"
	"github.com/harness/gitness/internal/paths"
)

func HandleDelete(connectorCtrl *connector.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		session, _ := request.AuthSessionFrom(ctx)
		connectorRef, err := request.GetConnectorRefFromPath(r)
		if err != nil {
			render.TranslatedUserError(w, err)
			return
		}
		spaceRef, connectorUID, err := paths.DisectLeaf(connectorRef)
		if err != nil {
			render.TranslatedUserError(w, err)
			return
		}

		err = connectorCtrl.Delete(ctx, session, spaceRef, connectorUID)
		if err != nil {
			render.TranslatedUserError(w, err)
			return
		}

		render.DeleteSuccessful(w)
	}
}