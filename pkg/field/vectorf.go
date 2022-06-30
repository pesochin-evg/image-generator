package field

import (
	"github.com/ojrac/opensimplex-go"
	"time"
)

type Vector struct {
	X float64
	Y float64
}

type Field struct {
	Seed int64
}

func New() *Field {
	var f Field
	f.Seed =  time.Now().UnixNano()
	return &f
}

func (f *Field) Get(x, y, Height, Width float64) (Vector, float64) {
	noise := opensimplex.NewNormalized(f.Seed)
	// fmt.Printf("%v\n", noise.Eval2(x, y))
	xFloat := (Height / 400 * x / Height)
	yFloat := (Width / 400 * y / Width)
	return Vector{X: noise.Eval2(xFloat, yFloat) - 0.5, Y: noise.Eval2(yFloat, xFloat) - 0.5}, 0.5
}
