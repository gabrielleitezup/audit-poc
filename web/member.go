package web

import (
	"audit-poc/internal/members"
	"audit-poc/util"
	"audit-poc/web/restutil"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func SaveMemberHandler(methods members.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		groupId, err := uuid.Parse(params["groupId"])
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		ctx := util.FillContext(r, "gabrielleite")

		request, err := methods.ParseMember(r.Body)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		response, err := methods.AssociateMemberToUserGroup(ctx, request, groupId)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		restutil.NewResponse(w, http.StatusCreated, response)
	}
}