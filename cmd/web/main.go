package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ivanelgran23/go-course/pkg/config"
	"github.com/ivanelgran23/go-course/pkg/handlers"
	"github.com/ivanelgran23/go-course/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{Addr: portNumber, Handler: routes(&app)}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
