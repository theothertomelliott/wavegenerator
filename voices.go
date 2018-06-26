package wavegenerator

import (
	"math"
	"math/rand"
	"time"
)

func Sin(samplingRate int, duration time.Duration, volume float64, freq float64) []float64 {
	count := float64(duration) / float64(time.Second) * float64(samplingRate)

	var buffer []float64
	for pos := 0.0; pos < count; pos++ {
		var a = 2 * math.Pi * freq * pos / float64(samplingRate)
		v := math.Sin(a) * volume
		buffer = append(buffer, v)
	}
	return buffer
}

func Square(samplingRate int, duration time.Duration, volume float64, freq float64) []float64 {
	count := float64(duration) / float64(time.Second) * float64(samplingRate)

	var buffer []float64
	for pos := 0.0; pos < count; pos++ {
		var a = 2 * math.Pi * freq * pos / float64(samplingRate)
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

func Sawtooth(samplingRate int, duration time.Duration, volume float64, freq float64) []float64 {
	count := float64(duration) / float64(time.Second) * float64(samplingRate)

	var buffer []float64
	for pos := 0.0; pos < count; pos++ {
		v := float64(int(pos) % int(float64(samplingRate)/freq))
		v = v / freq
		v = v/2 - 1
		buffer = append(buffer, v*volume)
	}
	return buffer
}

func Triangle(samplingRate int, duration time.Duration, volume float64, freq float64) []float64 {
	count := float64(duration) / float64(time.Second) * float64(samplingRate)

	var buffer []float64
	waveSize := int(float64(samplingRate)/freq) / 2
	for pos := 0; pos < int(count); pos++ {
		v := float64(pos % (waveSize))
		v = v / freq
		if int(pos/(waveSize))%2 == 1 {
			v = 1 - v
		}
		v = 2*v - 1

		buffer = append(buffer, v*volume)
	}
	return buffer
}

func Noise(samplingRate int, duration time.Duration, volume float64, freq float64) []float64 {
	count := float64(duration) / float64(time.Second) * float64(samplingRate)

	var buffer []float64
	for pos := 0.0; pos < count; pos++ {
		v := rand.Float64() - 0.5*2
		buffer = append(buffer, v*volume)
	}
	return buffer
}
