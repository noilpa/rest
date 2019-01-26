package routers

import (
	"github.com/gorilla/mux"
	"github.com/noilpa/rest/handles"
	"fmt"
)

func InitRouter() *mux.Router {

	r := mux.NewRouter()
	fmt.Sptint("asd")
	fmt.Sptint("asd")
	fmt.Sptint("asd")
	addHandlers(r)
	return r

}


func addHandlers(r *mux.Router) {

	// PUBLIC API

	// handle user registration requests
	// Example curl -H "Content-Type: application/json" -d
	// '{
	//	"user":{
	//    "login":"qwe",
	//	  "password":"asd"
	//  },
	//	"name":"bill",
	//	"age":69,
	//	"phone":"81234567890"
	// }' http://localhost:3333/registrations
	//
	// return new user id or error
	//
	// User registration in DB
	// Parameters
	// login - must have parameter, use for authorizations
	// password - must have parameter, use for authorizations
	// name - optional parameter, reference information
 	// age - optional parameter, reference information, format: only digits (12, 56, 89)
	// phone - optional parameter, reference information, format: only digits (12345, 81234567890)
	r.Handle("/registrations", handles.Registrations).Methods("POST")

	// handle user authorization requests
	// Example curl -H "Content-Type: application/json" -d
	// '{
	//  "login":"qwe",
	//	"password":"asd"
	// }' http://localhost:3333/authorizations
	//
	// return JWT token or error
	//
	// Parameters
	// login - must have parameter
	// password - must have parameter
	r.Handle("/authorizations", handles.Authorizations).Methods("POST")

	// PROTECTED API

	// handle films requests
	// Example: curl -H "Content-Type: application/json"
	// http://localhost:3333/films?date=>+11-10-2017&genre=horror+comedy&offset=1&size=10
	//
	// return json with list of films or error
	//
	// date - optional parameter with 1 or 2 args, used for filtering, (>, <, =, >=, <=) YYYY-MM-DD,
	//        default first value is "="
	// genre - optional parameter with list of args, used for filtering, Example: &genre=genre1+genre2+...
	// offset - must have parameter, used for pagination, value >= 0
	// size - must have parameter, used for pagination, value > 0
	r.Handle("/films", handles.JwtAuth(handles.Films)).Methods("GET")


	// у jwt существую ассиметричные и неассиметричные алгоритмы, но тк в ФТ нет повышенных требований к защищенности

}