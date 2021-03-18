package main

import (
	"audit-poc/internal/configuration"
	"audit-poc/internal/members"
	"audit-poc/internal/userworkspace"
	"audit-poc/internal/userworkspace/usergroup"
	"audit-poc/internal/userworkspace/workspace"
	"audit-poc/web"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func main() {

	godotenv.Load()

	fmt.Println("It's natural!")

	db, err := configuration.GetDBConnection("migrations")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.DB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database up")

	workspaceMain := workspace.NewMain(db)
	userGroupMain := usergroup.NewMain(db)
	userWorkspaceMain := userworkspace.NewMain(db)
	memberMain := members.NewMain(db)
	r := CreateRouter(workspaceMain, userGroupMain, userWorkspaceMain, memberMain)

	Start(r)
}

func CreateRouter(workspace workspace.ServiceMethods, usergroup usergroup.ServiceMethods, userworkspace userworkspace.ServiceMethods, member members.ServiceMethods) *mux.Router {
	r := mux.NewRouter().PathPrefix("/v1").Subrouter()
	{
		r.HandleFunc("/", web.HomeHandler).Methods("GET")
		r.HandleFunc("/workspaces", web.SaveWorkspaceHandler(workspace)).Methods("POST")
		r.HandleFunc("/workspaces/{workspaceId}", web.UpdateWorkspaceHandler(workspace)).Methods("PATCH")
		r.HandleFunc("/workspaces/{workspaceId}", web.DeleteWorkspaceHandler(workspace)).Methods("DELETE")
	}
	{
		r.HandleFunc("/user-groups", web.SaveUserGroupHandler(usergroup)).Methods("POST")
		r.HandleFunc("/user-groups/{groupId}", web.UpdateUserGroupHandler(usergroup)).Methods("PATCH")
		r.HandleFunc("/user-groups/{groupId}", web.DeleteUserGroupHandler(usergroup)).Methods("DELETE")
	}
	{
		r.HandleFunc("/workspaces/{workspaceId}/user-groups", web.SaveUserWorkspaceHandler(userworkspace)).Methods("POST")
	}
	{
		r.HandleFunc("/user-groups/{groupId}/members", web.SaveMemberHandler(member)).Methods("POST")
	}

	return r
}

func Start(r *mux.Router) {
	server := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
