package api

import (
	"fmt"
	"net/http"
	"test/repository"
)

type API struct {
	usersRepo repository.UserRepository
	mux       *http.ServeMux
}

func NewAPI(usersRepo repository.UserRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo, mux,
	}

	mux.Handle("/api/user/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/user/logout", api.POST(http.HandlerFunc(api.logout)))

	mux.Handle("/api/user/register", api.POST(http.HandlerFunc(api.register)))

	mux.Handle("/api/user/profile", api.GET(http.HandlerFunc(api.profile)))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("Starting Web Server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}
