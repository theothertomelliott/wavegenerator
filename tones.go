package wavegenerator

type Tone interface {
	Frequency() float64
	Voice() Voice
}

type TriangleTone float64

func (t TriangleTone) Frequency() float64 {
	return float64(t)
}

func (t TriangleTone) Voice() Voice {
	return generateTriangle
}

type SquareTone float64

func (t SquareTone) Frequency() float64 {
	return float64(t)
}

func (t SquareTone) Voice() Voice {
	return generateSquare
}

type SinTone float64

func (t SinTone) Frequency() float64 {
	return float64(t)
}

func (t SinTone) Voice() Voice {
	return generateSin
}
