package app

import "net/http"

func appHandler(fn func(http.ResponseWriter, *http.Request, *Context)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := NewContext(r)
		defer ctx.Close()
		fn(rw, r, ctx)
	}
}

func restrictedHandler(fn func(http.ResponseWriter, *http.Request, *Context)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := NewContext(r)
		defer ctx.Close()

		session, _ := cookies.Get(r, "mw-session")
		if session.Values["authenticated"] != true {
			http.Redirect(rw, r, "/admin/login", 302)
		}

		fn(rw, r, ctx)
	}
}

// func authHandler(next http.HandlerFunc) http.HandlerFunc {
// 	return func(rw http.ResponseWriter, r *http.Request) {
// 		session, _ := cookies.Get(r, "mw-session")
// 		if session.Values["authenticated"] != true {
// 			http.Redirect(rw, r, "/admin/login", 302)
// 		}
// 		next.ServeHTTP(rw, r)
// 	}
// }
