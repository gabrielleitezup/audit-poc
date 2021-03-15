package web

import (
	"audit-poc/internal/userworkspace/workspace"
	"audit-poc/web/util"
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(":)"))
}

func SaveWorkspaceHandler(methods workspace.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctxUser := context.WithValue(r.Context(), "jwt", "gabrielleite")
		ctxUserAgent := context.WithValue(ctxUser, "user-agent", r.UserAgent())
		ctxRemoteAddress := context.WithValue(ctxUserAgent, "user-ip", r.RemoteAddr)

		request, err := methods.ParseWorkspace(r.Body)
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		response, err := methods.SaveWorkspace(ctxRemoteAddress, request)
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		util.NewResponse(w, http.StatusCreated, response)
	}
}


func UpdateWorkspaceHandler(methods workspace.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		workspaceId, err := uuid.Parse(params["workspaceId"])
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		ctxUser := context.WithValue(r.Context(), "jwt", "gabrielleite")
		ctxUserAgent := context.WithValue(ctxUser, "user-agent", r.UserAgent())
		ctxRemoteAddress := context.WithValue(ctxUserAgent, "user-ip", r.RemoteAddr)

		request, err := methods.ParseWorkspace(r.Body)
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		response, err := methods.UpdateWorkspace(ctxRemoteAddress, request, workspaceId)
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		util.NewResponse(w, http.StatusOK, response)
	}
}

func DeleteWorkspaceHandler(methods workspace.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		workspaceId, err := uuid.Parse(params["workspaceId"])
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		ctxUser := context.WithValue(r.Context(), "jwt", "gabrielleite")
		ctxUserAgent := context.WithValue(ctxUser, "user-agent", r.UserAgent())
		ctxRemoteAddress := context.WithValue(ctxUserAgent, "user-ip", r.RemoteAddr)


		response, err := methods.DeleteWorkspace(ctxRemoteAddress, workspaceId)
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		util.NewResponse(w, http.StatusOK, response)
	}
}
