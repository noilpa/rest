package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/noilpa/rest/db"
	"github.com/noilpa/rest/utils"
	"net/http"
	"strconv"
)

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

var Registrations = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	var u utils.UserInfo
	if r.Body == nil {
		http.Error(w, "Error bad request with empty body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if utils.IsEmpty(u.Usr.Login) || utils.IsEmpty(u.Usr.Password) {
		http.Error(w, "Error bad request with empty login/password", 400)
		return
	}

	response, err := db.Registrations(u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
	}

	w.WriteHeader(201)
	fmt.Fprintln(w, string(resp))
})

var Authorizations = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	var u utils.User
	if r.Body == nil {
		http.Error(w, "Error bad request with empty body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if utils.IsEmpty(u.Login) || utils.IsEmpty(u.Password) {
		http.Error(w, "Error bad request with empty login/password", 400)
		return
	}

	isAuth, err := db.Authorizations(u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var token string
	if isAuth {
		token, err = Encode(u)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	response := map[string]interface{} {"JWT": token}
	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintln(w, string(resp))
})

var Films = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	offset, err := strconv.ParseUint(r.URL.Query().Get("offset"), 10, 32)
	if err != nil || offset < 0 {
		http.Error(w, "Error bad request", 400)
		return
	}

	size, err := strconv.ParseUint(r.URL.Query().Get("size"), 10, 32)
	if err != nil || size < 1 {
		http.Error(w, "Error bad request", 400)
		return
	}

	date := r.URL.Query().Get("date")
	genre := r.URL.Query().Get("genre")

	fmt.Println(offset, size, date, genre)
	films, err := db.Films(uint(size), uint(offset), date, genre)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	jFilms, err := json.Marshal(films)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(jFilms))
})
