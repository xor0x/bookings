package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/xor0x/bookings/internal/config"
	"github.com/xor0x/bookings/internal/handlers"
	"github.com/xor0x/bookings/internal/helpers"
	"github.com/xor0x/bookings/internal/models"
	"github.com/xor0x/bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)
var portNumber = ":8000"
var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the entry point of the program
func main() {

	gob.Register(models.Reservation{})
	

	// change this true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

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
	helpers.NewHlpers(&app)


	fmt.Println("Server is running on port", portNumber)


	srv := &http.Server{Addr: portNumber, Handler: routes(&app)}

	err = srv.ListenAndServe()
	log.Fatal(err)
}