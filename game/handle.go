package game

import (
	"fmt"
	"net/http"
)

func (g *Game) GetStatusHandler(w http.ResponseWriter, r *http.Request) {
	g.logger.Println("getStatusHandler ...")

	fmt.Fprintf(w, "Current Game status %s\n", g.status)
}

func (g *Game) AddPlayerHandler(w http.ResponseWriter, r *http.Request) {
	g.logger.Println("addPlayerHandler ...")

	clientId := r.PostFormValue("id")

	fmt.Fprintf(w, "Client: %s\n", clientId)
}
