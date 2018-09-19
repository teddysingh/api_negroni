package main

import (
	"net/http"

	"github.com/teddysingh/api_negroni/UserService/db"
	"github.com/teddysingh/api_negroni/UserService/utils"

	"github.com/teddysingh/api_negroni/UserService/handlers"
	"github.com/teddysingh/api_negroni/UserService/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/urfave/negroni"
)

var dbconn = db.Connect()
var logger = utils.Logger

func init() {
	logger.Println("Running migrations now...")
	dbconn.AutoMigrate(&models.User{})
}

func main() {
	defer dbconn.Close()

	serverDescription := map[string]string{
		"version": "1.0",
		"name":    "Hello World Server",
		"port":    ":3000",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handlers.CreateUser(w, r)
		} else {
			handlers.FindUser(w, r)
		}
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		utils.Success(w, serverDescription)
	})

	n := negroni.Classic()
	n.UseHandler(mux)

	logger.Printf("Starting %s on port %s\n", serverDescription["name"], serverDescription["port"])
	http.ListenAndServe(serverDescription["port"], n)
}
