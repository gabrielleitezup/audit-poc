package web

import (
	"audit-poc/internal/deployment"
	"audit-poc/util"
	"audit-poc/web/restutil"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func SaveDeploymentHandler(methods deployment.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		circleId, err := uuid.Parse(params["circleId"])
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		ctx := util.FillContext(r, "gabrielleite")

		request, err := methods.ParseDeployment(r.Body)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		response, err := methods.CreateDeployment(ctx, request, circleId)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		restutil.NewResponse(w, http.StatusCreated, response)
	}
}
