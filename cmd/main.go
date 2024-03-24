package main

import (
	"flag"
	"log"

	"github.com/cheina97/gomaze/pkg/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	w := flag.Int("w", 20, "width of the maze (cells number)")
	h := flag.Int("h", 20, "height of the maze (cells number)")
	flag.Parse()
	game, err := game.NewGame(*w, *h)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
