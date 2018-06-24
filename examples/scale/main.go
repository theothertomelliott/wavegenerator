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

	var (
		a      = wavegenerator.TriangleTone(220.0)
		b      = a * (1 + 1.0/8.0)
		c      = a * (1 + 2.0/8.0)
		cSharp = a * (1 + 2.5/8.0)
		e      = a * (1 + 4.0/8.0)
		fSharp = a * (1 + 5.5/8.0)
		g      = a * (1 + 7/8.0)
		A      = a * 2
	)

	player.Play(time.Second/4, a)
	player.Play(time.Second/4, b)
	player.Play(time.Second/4, c)
	player.Play(time.Second/4, cSharp)
	player.Play(time.Second/4, e)
	player.Play(time.Second/4, fSharp)
	player.Play(time.Second/4, g)
	player.Play(time.Second/4, A)
}
