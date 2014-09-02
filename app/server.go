package app

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
)

var (
	database   string
	mgoSession *mgo.Session
	cookies    *sessions.CookieStore
)

func init() {
	// TODO: change cookie key
	cookies = sessions.NewCookieStore([]byte("youdontknow"))

}

func New() {

	var err error
	mgoSession, err = mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/", appHandler(homeIndex)).Methods("GET")
	// Admin routes
	r.HandleFunc("/admin", restrictedHandler(adminIndex)).Methods("GET")
	r.HandleFunc("/admin/login", appHandler(adminLogin)).Methods("GET")
	r.HandleFunc("/admin/logout", appHandler(adminLogout)).Methods("GET")
	// API

	http.Handle("/", r)
	http.ListenAndServe(":3001", nil)
	log.Println("Listening 3001")

}
