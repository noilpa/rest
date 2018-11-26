package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
	"github.com/noilpa/rest/db"
	"github.com/noilpa/rest/routers"
	"net/http"
	"os"
)

// Entry point
func main() {

	err := db.InitDb()
	if err != nil {
		panic(err)
	}

	db := db.DB
	defer db.Close()

	r := routers.InitRouter()
	fmt.Println("Listening...")
	http.ListenAndServe(":3333", handlers.LoggingHandler(os.Stdout, r))

}
