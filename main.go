package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/hajimehoshi/oto"
)

func main() {
	player, err := oto.NewPlayer(samplingRate, 1, 2, samplingRate/10)
	if err != nil {
		panic(err)
	}

	wave := generateTriangle(samplingRate, 0.5, 440)
	for _, val := range wave {
		low, high := toBytes(val)
		_, err := player.Write([]byte{low, high})
		if err != nil {
			panic(err)
		}
	}
	wave = generateSquare(samplingRate, 0.5, 440)
	for _, val := range wave {
		low, high := toBytes(val)
		_, err := player.Write([]byte{low, high})
		if err != nil {
			panic(err)
		}
	}
	wave = generateSin(samplingRate, 0.5, 440)
	for _, val := range wave {
		low, high := toBytes(val)
		_, err := player.Write([]byte{low, high})
		if err != nil {
			panic(err)
		}
	}
	wave = generateNoise(samplingRate, 0.5, 440)
	for _, val := range wave {
		low, high := toBytes(val)
		_, err := player.Write([]byte{low, high})
		if err != nil {
			panic(err)
		}
	}

	player.Close()

}

func toBytes(val float64) (byte, byte) {
	if val < -1 {
		val = -1
	}
	if val > +1 {
		val = +1
	}
	valInt16 := int16(val * (1<<15 - 1))
	low := byte(valInt16)
	high := byte(valInt16 >> 8)
	return low, high
}

const samplingRate = 44100

func generateSin(count int, volume float64, freq float64) []float64 {
	var buffer []float64
	for pos := 0; pos < count; pos++ {
		var a = 2 * math.Pi * freq * float64(pos) / samplingRate
		v := math.Sin(a) * volume
		buffer = append(buffer, v)
	}
	return buffer
}

func generateSquare(count int, volume float64, freq float64) []float64 {
	var buffer []float64
	for pos := 0; pos < count; pos++ {
		var a = 2 * math.Pi * freq * float64(pos) / samplingRate
		v := math.Sin(a)
		if v >= 0.5 {
			v = 1.0
		}
		if v > -0.5 && v < 0.5 {
			v = 0.0
		}
		if v <= -0.5 {
			v = -1.0
		}
		buffer = append(buffer, v*volume)
	}
	return buffer
}

func generateTriangle(count int, volume float64, freq float64) []float64 {
	var buffer []float64
	for pos := 0; pos < count; pos++ {
		p := freq / 8
		a := float64(pos%int(2*p)) - p
		v := (2 / p) * (p - math.Abs(a))
		v = v - 1

		if pos <= 440 {
			fmt.Printf("%.2f ", v)
		}

		buffer = append(buffer, v*volume)
	}
	fmt.Println()
	return buffer
}

func generateNoise(count int, volume float64, freq float64) []float64 {
	var buffer []float64
	for pos := 0; pos < count; pos++ {
		v := rand.Float64() - 0.5*2
		buffer = append(buffer, v*volume)
	}
	return buffer
}
