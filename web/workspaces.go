package web

import (
	"audit-poc/internal/workspace"
	"audit-poc/util"
	"audit-poc/web/restutil"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func SaveWorkspaceHandler(methods workspace.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := util.FillContext(r, "gabrielleite", "workspaces")

		request, err := methods.ParseWorkspace(r.Body)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		response, err := methods.SaveWorkspace(ctx, request)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		restutil.NewResponse(w, http.StatusCreated, response)
	}
}

func UpdateWorkspaceHandler(methods workspace.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		workspaceId, err := uuid.Parse(params["workspaceId"])
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		ctx := util.FillContext(r, "gabrielleite", "workspaces")

		request, err := methods.ParseWorkspace(r.Body)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		response, err := methods.UpdateWorkspace(ctx, request, workspaceId)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		restutil.NewResponse(w, http.StatusOK, response)
	}
}

func DeleteWorkspaceHandler(methods workspace.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		workspaceId, err := uuid.Parse(params["workspaceId"])
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		ctx := util.FillContext(r, "gabrielleite", "workspaces")

		response, err := methods.DeleteWorkspace(ctx, workspaceId)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		restutil.NewResponse(w, http.StatusNoContent, response)
	}
}
