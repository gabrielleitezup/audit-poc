package web

import (
	"audit-poc/internal/userworkspace"
	"audit-poc/util"
	"audit-poc/web/restutil"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func SaveUserWorkspaceHandler(methods userworkspace.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := util.FillContext(r, "gabrielleite", "user_group_workspaces")

		params := mux.Vars(r)
		workspaceId, err := uuid.Parse(params["workspaceId"])
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		request, err := methods.ParseUserWorkspace(r.Body)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		response, err := methods.AssociateUserGroupToWorkspace(ctx, request, workspaceId)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		restutil.NewResponse(w, http.StatusCreated, response)
	}
}
