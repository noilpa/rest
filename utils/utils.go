package utils

import "time"

type User struct {
	Id		 int64   `json:"id"`
	Login 	 string `json:"login"`
	Password string `json:"password"`
}

type UserInfo struct {
	Usr	  User   `json:"user"`
	Name  string `json:"name"`
	Age   uint    `json:"age"`
	Phone string `json:"phone"`
}

type Film struct {
	Id	  int64	    `json:"id"`
	Name  string	`json:"name"`
	Date  time.Time `json:"date"`
	Genre string    `json:"genre"`
}

type Genre struct {
	Id   int64   `json:"id"`
	Name string `json:"name"`
}

// check for an element in a slice
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return false
}

