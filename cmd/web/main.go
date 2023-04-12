package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/TheHarryFer/go-bookings/internal/config"
	"github.com/TheHarryFer/go-bookings/internal/handlers"
	"github.com/TheHarryFer/go-bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = "localhost:8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	startServerMsg := fmt.Sprintf("Starting application on port %s", portNumber)
	fmt.Println(startServerMsg)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
