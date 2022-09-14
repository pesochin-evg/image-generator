package field

import (
	"github.com/ojrac/opensimplex-go"
)

type Vector struct {
	X float64
	Y float64
}

func (v *Vector) Swap() {
	v.X, v.Y = v.Y, v.X
}

type Field struct {
	Seed  int64
	freq  float64
	noise opensimplex.Noise
}

func New(seed int64) *Field {
	var f Field
	f.Seed = seed
	f.freq = float64(seed % 2)
	f.freq += 250
	f.noise = opensimplex.NewNormalized(f.Seed)
	return &f
}

func (f *Field) Get(x, y, Height, Width float64) (Vector, float64) {
	xFloat := (Height / f.freq * x / Height)
	yFloat := (Width / f.freq * y / Width)
	return Vector{X: f.noise.Eval2(xFloat, yFloat) - 0.5, Y: f.noise.Eval2(yFloat, xFloat) - 0.5}, 0.5
}
