package main

import (
	"time"

	"github.com/hajimehoshi/oto"
)

type Player struct {
	player       *oto.Player
	samplingRate int
	Volume       float64
}

func NewPlayer(samplingRate int) (*Player, error) {
	player, err := oto.NewPlayer(samplingRate, 1, 2, samplingRate/10)
	if err != nil {
		return nil, err
	}
	return &Player{
		player:       player,
		samplingRate: samplingRate,
		Volume:       0.25,
	}, nil
}

type Tone interface {
	Frequency() float64
	Voice() Voice
}

type TriangleTone float64

func (t TriangleTone) Frequency() float64 {
	return float64(t)
}

func (t TriangleTone) Voice() Voice {
	return generateNoise
}

func (p *Player) Play(duration time.Duration, tones ...Tone) {
	var wave []float64
	for _, tone := range tones {
		toneWave := tone.Voice()(p.samplingRate, duration, p.Volume, tone.Frequency())
		for index, v := range toneWave {
			if len(wave) <= index {
				wave = append(wave, v)
			} else {
				wave[index] += v
			}
		}
	}
	for _, val := range wave {
		low, high := toBytes(val)
		_, err := p.player.Write([]byte{low, high})
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

func (p *Player) Close() error {
	return p.player.Close()
}

type Voice func(samplingRate int, duration time.Duration, volume float64, freq float64) []float64
