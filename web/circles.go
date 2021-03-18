package web

import (
	"audit-poc/internal/circle"
	"audit-poc/util"
	"audit-poc/web/restutil"
	"net/http"
)

func SaveCircleHandler(methods circle.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := util.FillContext(r, "gabrielleite", "circles")

		request, err := methods.ParseCircle(r.Body)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		response, err := methods.CreateCircle(ctx, request)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		restutil.NewResponse(w, http.StatusCreated, response)
	}
}