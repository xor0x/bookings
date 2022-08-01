package main

import (
	"fmt"
	"log"
	"github.com/xor0x/bookings/pkg/config"
	"github.com/xor0x/bookings/pkg/handlers"
	"github.com/xor0x/bookings/pkg/render"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)
var portNumber = ":8000"
var app config.AppConfig
var session *scs.SessionManager

// main is the entry point of the program
func main() {
	

	// change this true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteStrictMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

 	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can't create template cache")
	}
	
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)


	fmt.Println("Server is running on port", portNumber)


	srv := &http.Server{Addr: portNumber, Handler: routes(&app)}

	err = srv.ListenAndServe()
	log.Fatal(err)
}