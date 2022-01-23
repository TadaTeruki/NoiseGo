package noise

import(
	"math"
	"math/rand"
)

type Noise struct{
	maxSize float64
	hash []int
	gradX, gradY []float64
}

func New(seed int64) Noise{
	var noise Noise
	noise.maxSize = 256.0
	noise.hash = make([]int, int(noise.maxSize))

	noise.gradX = make([]float64, 8)
	noise.gradY = make([]float64, 8)

	gi := 0
	for gy := -1.0; gy <= 1.0; gy++ {
		for gx := -1.0; gx <= 1.0; gx++ {
			if gx == 0 && gy == 0 {continue}
			noise.gradX[gi] = gx
			noise.gradY[gi] = gy
			gi++
		}
	}

	for i := 0; i < len(noise.hash); i++ { noise.hash[i] = i }

	rand.Seed(seed)
	rand.Shuffle(len(noise.hash), func(i, j int) {
		noise.hash[i], noise.hash[j] = noise.hash[j], noise.hash[i]
	})

	return noise
}

func mod(a, b float64) float64{
	rp := math.Floor(a/b)
	return a - b*rp
}

func fadeF(t float64) float64{
	return (t*(t*6-15)+10)*t*t*t
}

func linearF(x, a, b float64) float64{
	return x*(b-a)+a
}

func (noise *Noise) gradientF(hash int, ux, uy float64) float64{
	gradIndex := int(hash)%len(noise.gradX)
	return noise.gradX[gradIndex]*ux + noise.gradY[gradIndex]*uy
}

func (noise *Noise) getSquareHash(sqx, sqy float64) int{
	l := len(noise.hash)
	return noise.hash[(noise.hash[int(sqy)%l] + int(sqx))%l]
}

func (noise *Noise) Get(query_x, query_y float64) float64{

	qx := query_x * noise.maxSize
	qy := query_y * noise.maxSize

	var unitsq struct{sx, sy, ex, ey float64}
	var hashsq struct{ts, te, bs, be int}

	query_x = mod(qx, noise.maxSize)
	query_y = mod(qy, noise.maxSize)

	unitsq.sx = math.Floor(qx)
	unitsq.sy = math.Floor(qy)
	unitsq.ex = mod(unitsq.sx + 1.0, noise.maxSize)
	unitsq.ey = mod(unitsq.sy + 1.0, noise.maxSize)

	hashsq.ts = noise.getSquareHash(unitsq.sx, unitsq.sy)
	hashsq.te = noise.getSquareHash(unitsq.ex, unitsq.sy)
	hashsq.bs = noise.getSquareHash(unitsq.sx, unitsq.ey)
	hashsq.be = noise.getSquareHash(unitsq.ex, unitsq.ey)

	ux := qx - unitsq.sx
	uy := qy - unitsq.sy

	ga := linearF(fadeF(ux), noise.gradientF(hashsq.ts, ux, uy    ), noise.gradientF(hashsq.te, ux-1.0, uy    ))
	gb := linearF(fadeF(ux), noise.gradientF(hashsq.bs, ux, uy-1.0), noise.gradientF(hashsq.be, ux-1.0, uy-1.0))
	nz := linearF(fadeF(uy), ga, gb)
	
	return (nz+1.0)*0.5
}

func (noise *Noise) GetOctaved(query_x, query_y float64, octaves int, persistence float64) float64{
	
	sum := 0.0
	frequency := 1.0
	ampl := 1.0
	scale := 0.0

	for i := 0; i < octaves; i++ {
		scale += ampl
		sum += noise.Get(query_x*frequency, query_y*frequency) * ampl
		ampl *= persistence
		frequency *= 2
	}

	return sum/scale
}