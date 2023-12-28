package main

import (
	"log"
	"net/http"

	"github.com/donseba/go-htmx"
	"github.com/donseba/go-htmx/middleware"
)

type App struct {
	htmx *htmx.HTMX
}

func main() {
	// new app with htmx instance
	app := &App{
		htmx: htmx.New(),
	}

	mux := http.NewServeMux()
	// wrap the htmx example middleware around the http handler
	mux.Handle("/", middleware.MiddleWare(http.HandlerFunc(app.Home)))

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	// initiate a new htmx handler
	h := a.htmx.NewHandler(w, r)

	// set the headers for the response, see docs for more options
	h.PushURL("http://push.url")
	h.ReTarget("#ReTarged")

	// write the output like you normally do.
	// check inspector tool in browser to see that the headers are set.
	_, _ = h.Write([]byte("OK"))
}
