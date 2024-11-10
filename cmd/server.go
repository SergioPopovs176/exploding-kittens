package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SergioPopovs176/exploding-kittens/app"
)

func main() {
	app, _ := app.New()

	app.Logger.Println("Hello, Kitty. Version", app.Config.Version)
	app.Logger.Println("Start server")

	// Канал для приема сигнала остановки
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

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

	go func() {
		err := app.Game.Start()
		if err != nil {
			app.Logger.Fatal(err)
		}
	}()

	// Ожидаем сигнала остановки
	<-stop
	app.Logger.Println("Shutting down server...")
	// Контекст с тайм-аутом для корректного завершения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Завершаем работу сервера
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	app.Logger.Println("Server stopped gracefully")
	select {}
}

func getRouter(app *app.Application) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v0/healthcheck", app.HealthcheckHandler)
	mux.HandleFunc("GET /v0/client/{id}", app.ClientHandler)
	mux.HandleFunc("GET /v0/app/status", app.StatusHandler)

	mux.HandleFunc("GET /v0/game/status", app.Game.GetStatusHandler)
	mux.HandleFunc("POST /v0/game/add", app.Game.AddPlayerHandler)

	return mux
}
