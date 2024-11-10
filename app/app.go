package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SergioPopovs176/exploding-kittens/game"
)

const version = "0.0.6"
const serverName = "elvis"

type config struct {
	Port    int
	Env     string
	Version string
}

type Application struct {
	Config config
	Logger *log.Logger
	Game   *game.Game
}

func New() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	var cfg config

	flag.IntVar(&cfg.Port, "port", 8000, "API server port")
	flag.StringVar(&cfg.Env, "env", "dev", "Environment (dev|stage|prod)")
	flag.Parse()

	cfg.Version = version

	game := game.Ini(logger)

	app := &Application{
		Config: cfg,
		Logger: logger,
		Game:   game,
	}

	return app, nil
}

func (app *Application) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("healthcheckHandler ...")

	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.Config.Env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *Application) ClientHandler(w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("clientHandler ...")

	clientId := r.PathValue("id")

	fmt.Fprintf(w, "Client: %s\n", clientId)
}

func (app *Application) StatusHandler(w http.ResponseWriter, r *http.Request) {
	app.Logger.Println("gameStatusHandler ...")

	fmt.Fprintf(w, "server: %s\n", serverName)
	fmt.Fprintf(w, "version: %s\n", version)
	fmt.Fprintln(w, "App status: ready")
	fmt.Fprintf(w, "Client amount: %d\n", 0)
}
