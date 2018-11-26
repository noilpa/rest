package main

import (
	"bytes"
	"net/http"
	"restreamTestCase/db"
	"testing"
)

func TestDBConnect(t *testing.T) {
	t.Log("Database availability testing")
	err := db.InitDb()
	checkError(err, t)
}

func TestRegistration(t *testing.T) {
	data := []byte(`{
		"user":{
			"login": "qwe13",
			"password": "qwe4"
		},
		"name":"billy",
		"age":69,
		"phone":"81234567890"
	}`)
	r := bytes.NewReader(data)
	http.NewRequest("POST", "/authorizations", r)

}

func TestAuthorization(t *testing.T) {

	data := []byte(`{"foo":"bar"}`)
	r := bytes.NewReader(data)
	http.NewRequest("POST", "/authorizations", r)
}

func TestFilms(t *testing.T) {
	req, err := http.NewRequest("GET", "/films?offset=0&size=3&date==+1988-03-29", nil)
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}


