package web

import (
	"audit-poc/internal/userworkspace/usergroup"
	"audit-poc/web/util"
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func SaveUserGroupHandler(methods usergroup.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctxUser := context.WithValue(r.Context(), "jwt", "gabrielleite")
		ctxUserAgent := context.WithValue(ctxUser, "user-agent", r.UserAgent())
		ctxRemoteAddress := context.WithValue(ctxUserAgent, "user-ip", r.RemoteAddr)

		request, err := methods.ParseUserGroup(r.Body)
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		response, err := methods.SaveUserGroup(ctxRemoteAddress, request)
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		util.NewResponse(w, http.StatusCreated, response)
	}
}

func UpdateUserGroupHandler(methods usergroup.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		groupId, err := uuid.Parse(params["groupId"])
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		ctxUser := context.WithValue(r.Context(), "jwt", "gabrielleite")
		ctxUserAgent := context.WithValue(ctxUser, "user-agent", r.UserAgent())
		ctxRemoteAddress := context.WithValue(ctxUserAgent, "user-ip", r.RemoteAddr)

		request, err := methods.ParseUserGroup(r.Body)
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		response, err := methods.UpdateUserGroup(ctxRemoteAddress, request, groupId)
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		util.NewResponse(w, http.StatusOK, response)
	}
}

func DeleteUserGroupHandler(methods usergroup.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		groupId, err := uuid.Parse(params["groupId"])
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		ctxUser := context.WithValue(r.Context(), "jwt", "gabrielleite")
		ctxUserAgent := context.WithValue(ctxUser, "user-agent", r.UserAgent())
		ctxRemoteAddress := context.WithValue(ctxUserAgent, "user-ip", r.RemoteAddr)


		response, err := methods.DeleteUserGroup(ctxRemoteAddress, groupId)
		if err != nil {
			util.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		util.NewResponse(w, http.StatusOK, response)
	}
}