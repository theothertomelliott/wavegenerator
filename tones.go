package wavegenerator

import "time"

type Tone interface {
	Duration() time.Duration
	Frequency() float64
	Voice() Voice
}

// NewTone returns a tone with the specified properties
func NewTone(duration time.Duration, frequency float64, voice Voice) Tone {
	return &simpleTone{
		duration:  duration,
		frequency: frequency,
		voice:     voice,
	}
}

type simpleTone struct {
	duration  time.Duration
	frequency float64
	voice     Voice
}

func (s *simpleTone) Duration() time.Duration {
	return s.duration
}

func (s *simpleTone) Frequency() float64 {
	return s.frequency
}

func (s *simpleTone) Voice() Voice {
	return s.voice
}
