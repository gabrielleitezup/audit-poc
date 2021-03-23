package web

import (
	"audit-poc/internal/auditions"
	"audit-poc/util"
	"audit-poc/web/restutil"
	"net/http"
)

func HistoryHandler(methods auditions.ServiceMethods) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := util.FillContext(r, "gabrielleite")

		params := make(map[string]interface{})

		GetParam(r, params, "entity_id", "entityId")
		GetParam(r, params, "username", "username")
		GetParam(r, params, "table_name", "entity")
		GetParam(r, params, "operation", "operation")

		response, err := methods.HistoryList(ctx, params)
		if err != nil {
			restutil.NewResponse(w, http.StatusInternalServerError, err)
			return
		}

		restutil.NewResponse(w, http.StatusOK, response)
	}
}

func GetParam(r *http.Request, parameters map[string]interface{}, key, value string) {
	pr := r.URL.Query().Get(value)
	if pr != "" {
		parameters[key] = pr
	}
}
