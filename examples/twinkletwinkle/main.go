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
		c = wavegenerator.TriangleTone(261.63)
		d = c * (1 + 1.0/8.0)
		e = c * (1 + 2.0/8.0)
		f = c * (1 + 2.5/8.0)
		g = c * (1 + 4.0/8.0)
		a = c * (1 + 5.5/8.0)
	)

	player.Play(time.Second/4, c)
	player.Play(time.Second/4, c)
	player.Play(time.Second/4, g)
	player.Play(time.Second/4, g)
	player.Play(time.Second/4, a)
	player.Play(time.Second/4, a)
	player.Play(time.Second/2, g)

	player.Play(time.Second/4, f)
	player.Play(time.Second/4, f)
	player.Play(time.Second/4, e)
	player.Play(time.Second/4, e)
	player.Play(time.Second/4, d)
	player.Play(time.Second/4, d)
	player.Play(time.Second/2, c)

	player.Play(time.Second/4, g)
	player.Play(time.Second/4, g)
	player.Play(time.Second/4, f)
	player.Play(time.Second/4, f)
	player.Play(time.Second/4, e)
	player.Play(time.Second/4, e)
	player.Play(time.Second/2, d)

	player.Play(time.Second/4, g)
	player.Play(time.Second/4, g)
	player.Play(time.Second/4, f)
	player.Play(time.Second/4, f)
	player.Play(time.Second/4, e)
	player.Play(time.Second/4, e)
	player.Play(time.Second/2, d)

	player.Play(time.Second/4, c)
	player.Play(time.Second/4, c)
	player.Play(time.Second/4, g)
	player.Play(time.Second/4, g)
	player.Play(time.Second/4, a)
	player.Play(time.Second/4, a)
	player.Play(time.Second/2, g)

	player.Play(time.Second/4, f)
	player.Play(time.Second/4, f)
	player.Play(time.Second/4, e)
	player.Play(time.Second/4, e)
	player.Play(time.Second/4, d)
	player.Play(time.Second/4, d)
	player.Play(time.Second/2, c)
}
