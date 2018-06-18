package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/oto"
)

func main() {
	const samplingRate = 44100
	player, err := oto.NewPlayer(samplingRate, 1, 2, samplingRate/10)
	if err != nil {
		panic(err)
	}

	testWaves(player, samplingRate, time.Second/4, 220)
	testWaves(player, samplingRate, time.Second/4, 220*1.3)
	testWaves(player, samplingRate, time.Second/4, 220*1.5)
	testWaves(player, samplingRate, time.Second/4, 440)

	player.Close()

}

func testWaves(player *oto.Player, samplingRate int, duration time.Duration, freq float64) {

	var wave []float64
	wave = generateSawtooth(samplingRate, duration, 0.5, freq)
	for _, val := range wave {
		low, high := toBytes(val)
		_, err := player.Write([]byte{low, high})
		if err != nil {
			panic(err)
		}
	}
	wave = generateTriangle(samplingRate, duration, 0.5, freq)
	for _, val := range wave {
		low, high := toBytes(val)
		_, err := player.Write([]byte{low, high})
		if err != nil {
			panic(err)
		}
	}
	wave = generateSquare(samplingRate, duration, 0.5, freq)
	for _, val := range wave {
		low, high := toBytes(val)
		_, err := player.Write([]byte{low, high})
		if err != nil {
			panic(err)
		}
	}
	wave = generateSin(samplingRate, duration, 0.5, freq)
	for _, val := range wave {
		low, high := toBytes(val)
		_, err := player.Write([]byte{low, high})
		if err != nil {
			panic(err)
		}
	}
	wave = generateNoise(samplingRate, duration, 0.5, freq)
	for _, val := range wave {
		low, high := toBytes(val)
		_, err := player.Write([]byte{low, high})
		if err != nil {
			panic(err)
		}
	}
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

func generateSin(samplingRate int, duration time.Duration, volume float64, freq float64) []float64 {
	count := float64(duration) / float64(time.Second) * float64(samplingRate)

	var buffer []float64
	for pos := 0; pos < int(count); pos++ {
		var a = 2 * math.Pi * freq * float64(pos) / float64(samplingRate)
		v := math.Sin(a) * volume
		buffer = append(buffer, v)
	}
	return buffer
}

func generateSquare(samplingRate int, duration time.Duration, volume float64, freq float64) []float64 {
	count := float64(duration) / float64(time.Second) * float64(samplingRate)

	var buffer []float64
	for pos := 0; pos < int(count); pos++ {
		var a = 2 * math.Pi * freq * float64(pos) / float64(samplingRate)
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

func generateSawtooth(samplingRate int, duration time.Duration, volume float64, freq float64) []float64 {
	count := float64(duration) / float64(time.Second) * float64(samplingRate)

	var buffer []float64
	for pos := 0; pos < int(count); pos++ {
		v := float64(pos % int(float64(samplingRate)/freq))
		v = v / freq
		v = v/2 - 1
		buffer = append(buffer, v*volume)
	}
	return buffer
}

func generateTriangle(samplingRate int, duration time.Duration, volume float64, freq float64) []float64 {
	count := float64(duration) / float64(time.Second) * float64(samplingRate)

	var buffer []float64
	waveSize := int(float64(samplingRate)/freq) / 2
	for pos := 0; pos < int(count); pos++ {
		v := float64(pos % (waveSize))
		v = v / freq
		v = v/2 - 1

		if int(pos/(waveSize))%2 == 1 {
			v = 1 - v
		}

		buffer = append(buffer, v*volume)
	}
	return buffer
}

func generateNoise(samplingRate int, duration time.Duration, volume float64, freq float64) []float64 {
	count := float64(duration) / float64(time.Second) * float64(samplingRate)

	var buffer []float64
	for pos := 0; pos < int(count); pos++ {
		v := rand.Float64() - 0.5*2
		buffer = append(buffer, v*volume)
	}
	return buffer
}
