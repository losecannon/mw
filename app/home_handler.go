package app

import (
	"net/http"

	"bitbucket.org/fayub/mobilewala/render"
)

func homeIndex(rw http.ResponseWriter, req *http.Request, ctx *Context) {
	render.HTML(rw, "home/home", "layout", nil)
}
