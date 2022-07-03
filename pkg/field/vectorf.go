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
	Seed int64
}

func New(seed int64) *Field {
	var f Field
	f.Seed = seed
	return &f
}

func (f *Field) Get(x, y, Height, Width float64) (Vector, float64) {
	noise := opensimplex.NewNormalized(f.Seed)
	xFloat := (Height / 400 * x / Height)
	yFloat := (Width / 400 * y / Width)
	return Vector{X: noise.Eval2(xFloat, yFloat) - 0.5, Y: noise.Eval2(yFloat, xFloat) - 0.5}, 0.5
}
