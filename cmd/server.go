package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SergioPopovs176/exploding-kittens/game"
)

const version = "0.0.5"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
	game   *game.Game
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logger.Println("Hello, Kitty. Version", version)
	logger.Println("Start server")

	var cfg config

	flag.IntVar(&cfg.port, "port", 8000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	flag.Parse()

	game := game.Ini(logger)

	app := &application{
		config: cfg,
		logger: logger,
		game:   game,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v0/healthcheck", app.healthcheckHandler)
	mux.HandleFunc("GET /v0/client/{id}", app.clientHandler)
	mux.HandleFunc("GET /v0/app/status", app.gameStatusHandler)

	mux.HandleFunc("GET /v0/game/status", app.game.GetStatusHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
		err := srv.ListenAndServe()
		logger.Fatal(err)
	}()

	err := game.Start()
	if err != nil {
		logger.Fatal(err)
	}

	select {}
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Println("healthcheckHandler ...")

	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) clientHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Println("clientHandler ...")

	clientId := r.PathValue("id")

	fmt.Fprintf(w, "Client: %s\n", clientId)
}

func (app *application) gameStatusHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Println("gameStatusHandler ...")

	fmt.Fprintln(w, "App status: ready")
	fmt.Fprintf(w, "Client amount: %d\n", 0)
	fmt.Fprintf(w, "version: %s\n", version)
}
