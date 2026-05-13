package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rohanv8676/bookings/pkg/config"
	"github.com/rohanv8676/bookings/pkg/handlers"
	"github.com/rohanv8676/bookings/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

// Main application function
func main() {

	// Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.UseCache = false

	app.TemplateCache = templateCache

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting app on port %s", portNumber)

	serv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serv.ListenAndServe()
	log.Fatal(serv)
}
