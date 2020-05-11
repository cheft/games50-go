package particles

import (
	"image/color"
	"time"
)

type ParticleEmitterConfig struct {
	MaxParticles int
	Lifetime     Range
	Acceleration Acceleration
	Spawn        Spawn
	Colours      []color.Color
	EmitterLife  time.Duration
}

type Range struct {
	Min float64
	Max float64
}

type Acceleration struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

type Position struct {
	X float64
	Y float64
}

func (e *ParticleEmitter) SetColours(colours []color.Color) {
	e.Config.Colours = colours
}
