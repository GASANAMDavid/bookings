package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/alexedwards/scs/v2"
	"github.com/gasanamdavid/bookings/pkg/config"
	"github.com/gasanamdavid/bookings/pkg/handlers"
	"github.com/gasanamdavid/bookings/pkg/render"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {


	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Println("Starting application on port", portNumber, ".........")
	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
