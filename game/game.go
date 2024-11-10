package game

import (
	"log"
	"time"
)

type Game struct {
	logger  *log.Logger
	status  string
	counter int
}

func Ini(logger *log.Logger) *Game {
	g := &Game{
		logger: logger,
		status: "new",
	}

	return g
}

func (g *Game) Start() error {
	exit := false

	for counter := 1; !exit; counter++ {
		g.logger.Printf("%s-%d\n", g.status, counter)
		g.counter = counter
		time.Sleep(5 * time.Second)

		if counter == 15 {
			exit = true
			g.status = "finish"
		}
	}

	return nil
}
