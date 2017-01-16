package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"basicGo/config"
	"basicGo/handlers"
)

func HandleRequests(context config.Context){
	r:=mux.NewRouter()
	r.HandleFunc("/login",handlers.Login(context))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/",r)
}