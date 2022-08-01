package main

import (
	"net/http"
	"github.com/justinas/nosurf"
)


// NoSurf adds CSRF protection to the request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure:   app.InProduction,
		Path:    "/",
		SameSite: http.SameSiteStrictMode,
	})
	return csrfHandler
}

// SessionLoad is a middleware that loads the session and save it to the request
func SessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}