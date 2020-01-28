package handler

import (
	"github.com/go-chi/chi"

)


func NewRouter(bh *blacklistCheckHandler)*chi.Mux{
	router := chi.NewRouter()

	router.Get( "/health", healthCheck)
	router.Post( "/blacklist/check", bh.blacklistCheck)

	return router
}
