package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// noSurf adds CSRF protection on every request
func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// sessionLoad loads and save session on every request
func sessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
