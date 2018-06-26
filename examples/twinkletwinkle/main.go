package main

import (
	"time"

	"github.com/theothertomelliott/wavegenerator"
)

func main() {
	player, err := wavegenerator.NewPlayer(44100)
	if err != nil {
		panic(err)
	}
	defer player.Close()

	i := wavegenerator.NewInstrument(wavegenerator.Triangle, time.Second/4)

	player.Play(i.Tone("C", 4, 1))
	player.Play(i.Tone("C", 4, 1))
	player.Play(i.Tone("G", 4, 1))
	player.Play(i.Tone("G", 4, 1))
	player.Play(i.Tone("A", 4, 1))
	player.Play(i.Tone("A", 4, 1))
	player.Play(i.Tone("G", 4, 2))

	player.Play(i.Tone("F", 4, 1))
	player.Play(i.Tone("F", 4, 1))
	player.Play(i.Tone("E", 4, 1))
	player.Play(i.Tone("E", 4, 1))
	player.Play(i.Tone("D", 4, 1))
	player.Play(i.Tone("D", 4, 1))
	player.Play(i.Tone("C", 4, 2))

	player.Play(i.Tone("G", 4, 1))
	player.Play(i.Tone("G", 4, 1))
	player.Play(i.Tone("F", 4, 1))
	player.Play(i.Tone("F", 4, 1))
	player.Play(i.Tone("E", 4, 1))
	player.Play(i.Tone("E", 4, 1))
	player.Play(i.Tone("D", 4, 2))

	player.Play(i.Tone("G", 4, 1))
	player.Play(i.Tone("G", 4, 1))
	player.Play(i.Tone("F", 4, 1))
	player.Play(i.Tone("F", 4, 1))
	player.Play(i.Tone("E", 4, 1))
	player.Play(i.Tone("E", 4, 1))
	player.Play(i.Tone("D", 4, 2))

	player.Play(i.Tone("C", 4, 1))
	player.Play(i.Tone("C", 4, 1))
	player.Play(i.Tone("G", 4, 1))
	player.Play(i.Tone("G", 4, 1))
	player.Play(i.Tone("A", 4, 1))
	player.Play(i.Tone("A", 4, 1))
	player.Play(i.Tone("G", 4, 2))

	player.Play(i.Tone("F", 4, 1))
	player.Play(i.Tone("F", 4, 1))
	player.Play(i.Tone("E", 4, 1))
	player.Play(i.Tone("E", 4, 1))
	player.Play(i.Tone("D", 4, 1))
	player.Play(i.Tone("D", 4, 1))
	player.Play(i.Tone("C", 4, 2))
}
