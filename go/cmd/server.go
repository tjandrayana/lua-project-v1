package main

import (
	"fmt"
	"github.com/tjandrayana/lua-project-v1/go/handler"
	"github.com/tjandrayana/lua-project-v1/go/services"

	"net/http"
	grace "gopkg.in/paytm/grace.v1"
)

func startServer(handler http.Handler) error {

	fmt.Println("Serving HTTP on port :8007")
	return grace.Serve(":8007",handler)
}



func startApp()error{

	bm := services.NewCheckBlacklist()
	bh := handler.NewBlacklistHandler(bm)
	router := handler.NewRouter(bh)
	return startServer(router)
}