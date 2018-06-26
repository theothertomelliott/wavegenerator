package wavegenerator

import "time"

type Tone interface {
	Duration() time.Duration
	Frequency() float64
	Voice() Voice
}
