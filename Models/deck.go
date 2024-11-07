package models

import "fmt"

type Deck struct {
	cards []Card
}

func (d *Deck) shuffle() {
	fmt.Println("Deck was shuffled")
}
