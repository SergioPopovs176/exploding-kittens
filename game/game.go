package game

import (
	"log"
	"os"
	"syscall"
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

		if counter == 10 {
			exit = true
			g.status = "finish"
		}
	}

	g.logger.Println("Sending SIGTERM to self")
	// Отправка SIGTERM самому себе
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		g.logger.Println("Error finding process:", err)
		return err
	}

	// Отправка сигнала SIGTERM текущему процессу
	if err := p.Signal(syscall.SIGTERM); err != nil {
		g.logger.Println("Error sending SIGTERM:", err)
		return err
	}

	return nil
}
