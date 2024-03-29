package field

import (
	"github.com/ojrac/opensimplex-go"
)

const (
	freqStart = 250 // minimal frequency value
	freqDiff  = 50  // max value, that add to freqStart
	maxLength = 0.7 // Length of the longest vector
)

// Vector type that stores X,Y coordinates
// of the geometrical vector
type Vector struct {
	X float64
	Y float64
}

// Field where each point is a vector with values X,Y that is generated by noise
// function
type Field struct {
	Seed  int64
	freq  float64
	noise opensimplex.Noise
}

// GenerateField generates new field object and 2D noise
// based on given seed variable
func GenerateField(seed int64) *Field {
	var f Field
	f.Seed = seed
	f.freq = freqStart + float64(seed%freqDiff)
	f.noise = opensimplex.NewNormalized(f.Seed)
	return &f
}

// GenerateFieldFreq Generates new field object and 2D noise based on given seed
// and specified frequency variable
func GenerateFieldFreq(seed int64, freq float64) *Field {
	var f Field
	f.Seed = seed
	f.freq = freq
	f.noise = opensimplex.NewNormalized(f.Seed)
	return &f
}

// Get returns vector for given point and maximal length of it
func (f *Field) Get(x, y int) (Vector, float64) {
	fX := float64(x) / f.freq
	fY := float64(y) / f.freq
	return Vector{X: f.noise.Eval2(fX, fY) - 0.5, Y: f.noise.Eval2(fX+1000, fY+1000) - 0.5}, maxLength
}
