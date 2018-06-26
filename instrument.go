package wavegenerator

import (
	"math"
	"time"
)

// NewInstrument creates an Instrument for a specific voice
func NewInstrument(voice Voice, timeUnit time.Duration) *Instrument {
	return &Instrument{
		voice:    voice,
		timeUnit: timeUnit,
		// Note frequencies from: https://pages.mtu.edu/~suits/notefreqs.html
		scale0: map[string]float64{
			"C":  16.35,
			"C#": 17.32,
			"Db": 17.32,
			"D":  18.35,
			"D#": 19.45,
			"Eb": 19.45,
			"E":  20.60,
			"F":  21.83,
			"F#": 23.12,
			"Gb": 23.12,
			"G":  24.50,
			"G#": 25.96,
			"Ab": 25.96,
			"A":  27.50,
			"A#": 29.14,
			"Bb": 29.14,
			"B":  30.87,
		},
	}
}

type Instrument struct {
	voice    Voice
	timeUnit time.Duration
	scale0   map[string]float64
}

func (i *Instrument) Tone(note string, scale uint, duration float64) Tone {
	f0 := i.scale0[note]
	a := math.Pow(2.0, 1.0/12.0)
	n := (12.0 * float64(scale))
	fn := f0 * math.Pow(a, n)
	return &instrumentTone{
		instrument: i,
		frequency:  fn,
		duration:   i.timeUnit * time.Duration(duration),
	}
}

var _ Tone = &instrumentTone{}

type instrumentTone struct {
	instrument *Instrument
	frequency  float64
	duration   time.Duration
}

func (i *instrumentTone) Duration() time.Duration {
	return i.duration
}

func (i *instrumentTone) Frequency() float64 {
	return i.frequency
}

func (i *instrumentTone) Voice() Voice {
	return i.instrument.voice
}
