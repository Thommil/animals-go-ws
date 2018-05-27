package users

import (
	"github.com/gorilla/mux"
	"io"
	//"log"
	"net/http"
)

//Mongo : https://godoc.org/github.com/globalsign/mgo

//API implementation --> defines subroutes for /user
func GetRoutes(router *mux.Router) error {
	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		io.WriteString(response, "OK OK")
	})
	return nil
}
