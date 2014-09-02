package app

import (
	"log"
	"net/http"

	"bitbucket.org/fayub/mobilewala/render"
	"bitbucket.org/fayub/mobilewala/repo/user"
)

var count = 1

func adminIndex(rw http.ResponseWriter, r *http.Request, ctx *Context) {

	render.HTML(rw, "admin/home", "admin/layout", nil)
}

func adminLogin(rw http.ResponseWriter, r *http.Request, ctx *Context) {
	u := user.NewUser()
	u.Username = "fa"
	u.Password = "xyz"
	u.Add(ctx.DbSession)
	render.HTML(rw, "admin/login", "admin/layout", nil)
}

func adminLogout(rw http.ResponseWriter, r *http.Request, ctx *Context) {
	u := user.NewUser()
	u.Username = "fa"
	u.Password = "xyz"
	result, _ := u.Find(ctx.DbSession)
	count++
	log.Print(count)
	render.JSON(rw, 200, result)
}
