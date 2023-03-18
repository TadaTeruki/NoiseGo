package noise

import (
	"math"
	"math/rand"
)

var gradX = [8]float64{1.0, 1.0, -1.0, -1.0, 0.0, 0.0, 1.0, -1.0}
var gradY = [8]float64{1.0, -1.0, 1.0, -1.0, 1.0, -1.0, 0.0, 0.0}

type Noise struct {
	maxValue float64
	hash     []int
}

func New(maxValue uint64) Noise {

	var noise Noise
	noise.maxValue = float64(maxValue)
	noise.hash = make([]int, 256)

	for i := 0; i < len(noise.hash); i++ {
		noise.hash[i] = i
	}

	rand.Shuffle(len(noise.hash), func(i, j int) {
		noise.hash[i], noise.hash[j] = noise.hash[j], noise.hash[i]
	})

	return noise
}

func mod(a, b float64) float64 {
	rp := math.Floor(a / b)
	return a - b*rp
}

func fadeF(t float64) float64 {
	return (t*(t*6-15) + 10) * t * t * t
}

func linearF(x, a, b float64) float64 {
	return x*(b-a) + a
}

func (noise *Noise) gradientF(hash int, ux, uy float64) float64 {
	gradIndex := int(hash) % 8
	return gradX[gradIndex]*ux + gradY[gradIndex]*uy
}

func (noise *Noise) getSquareHash(sqx, sqy float64) int {
	l := len(noise.hash)
	return noise.hash[(noise.hash[int(sqy)%l]+int(sqx))%l]
}

func (noise *Noise) Get(query_x, query_y float64) float64 {

	if noise.maxValue == 0 {
		return 0.0
	}

	var unitsq struct{ sx, sy, ex, ey float64 }
	var hashsq struct{ ts, te, bs, be int }

	query_x = mod(query_x, noise.maxValue)
	query_y = mod(query_y, noise.maxValue)

	unitsq.sx = math.Floor(query_x)
	unitsq.sy = math.Floor(query_y)
	unitsq.ex = mod(unitsq.sx+1.0, noise.maxValue)
	unitsq.ey = mod(unitsq.sy+1.0, noise.maxValue)

	hashsq.ts = noise.getSquareHash(unitsq.sx, unitsq.sy)
	hashsq.te = noise.getSquareHash(unitsq.ex, unitsq.sy)
	hashsq.bs = noise.getSquareHash(unitsq.sx, unitsq.ey)
	hashsq.be = noise.getSquareHash(unitsq.ex, unitsq.ey)

	ux := query_x - unitsq.sx
	uy := query_y - unitsq.sy

	ga := linearF(fadeF(ux), noise.gradientF(hashsq.ts, ux, uy), noise.gradientF(hashsq.te, ux-1.0, uy))
	gb := linearF(fadeF(ux), noise.gradientF(hashsq.bs, ux, uy-1.0), noise.gradientF(hashsq.be, ux-1.0, uy-1.0))
	nz := linearF(fadeF(uy), ga, gb)

	return (nz + 1.0) * 0.5
}

func (noise *Noise) GetOctaved(query_x, query_y float64, octaves int, persistence float64) float64 {

	sum_noise := 0.0
	frequency := 1.0
	ampl := 1.0
	sum_scale := 0.0

	for i := 0; i < octaves; i++ {
		sum_scale += ampl
		sum_noise += noise.Get(query_x*frequency, query_y*frequency) * ampl
		ampl *= persistence
		frequency *= 2
	}

	return sum_noise / sum_scale
}
