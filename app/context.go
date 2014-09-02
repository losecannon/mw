package app

import (
	"net/http"

	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
)

type Context struct {
	// Database *mgo.Database
	DbSession *mgo.Session
	Cookies   *sessions.CookieStore
}

func (c *Context) Close() {
	c.DbSession.Close()
}

func NewContext(req *http.Request) *Context {
	return &Context{
		DbSession: mgoSession.Copy(),
		Cookies:   cookies,
	}
}
