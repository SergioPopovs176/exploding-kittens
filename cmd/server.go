package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SergioPopovs176/exploding-kittens/app"
)

func main() {
	app, _ := app.New()

	app.Logger.Println("Hello, Kitty. Version", app.Config.Version)
	app.Logger.Println("Start server")

	mux := getRouter(app)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		app.Logger.Printf("starting %s server on %s", app.Config.Env, srv.Addr)
		err := srv.ListenAndServe()
		app.Logger.Fatal(err)
	}()

	err := app.Game.Start()
	if err != nil {
		app.Logger.Fatal(err)
	}

	select {}
}

func getRouter(app *app.Application) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v0/healthcheck", app.HealthcheckHandler)
	mux.HandleFunc("GET /v0/client/{id}", app.ClientHandler)
	mux.HandleFunc("GET /v0/app/status", app.StatusHandler)

	mux.HandleFunc("GET /v0/game/status", app.Game.GetStatusHandler)

	return mux
}
