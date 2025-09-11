package main

import (
	"expkit/internal/config"
	"fmt"
	"log"
	"net/http"
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "auth-server"
)

func main() {
	fmt.Println("Auth service starting ...")

	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("config: %+v\n", cfg)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
