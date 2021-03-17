package web

import (
	"audit-poc/internal/userworkspace"
	"audit-poc/web/restutil"
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func SaveUserWorkspaceHandler(methods userworkspace.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctxUser := context.WithValue(r.Context(), "jwt", "gabrielleite")
		ctxUserAgent := context.WithValue(ctxUser, "user-agent", r.UserAgent())
		ctxRemoteAddress := context.WithValue(ctxUserAgent, "user-ip", r.RemoteAddr)

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

		response, err := methods.AssociateUserGroupToWorkspace(ctxRemoteAddress, request, workspaceId)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		restutil.NewResponse(w, http.StatusCreated, response)
	}
}
