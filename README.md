# NoiseGo - Simple 2D noise generator for go

Implementation of [**Improved Perlin Noise**](https://doi.org/10.1145/566654.566636) for Golang.

## Preview

<img src="https://user-images.githubusercontent.com/69315285/150670088-e690b5f4-b15f-4950-b959-a143277994f1.png" height="300"> <img src="https://user-images.githubusercontent.com/69315285/150670142-3a5530cf-30f3-4dd7-9d34-36e9aec0c988.png" height="300">

## Usage

```go

import(
    noise "github.com/TadaTeruki/NoiseGo"
)

func main(){
//  noise.New(seed int64, maxValue uint64) noise.Noise             
//   - Initialize noise generator                     
//  < arguments >                                     
//      seed : Seed value of noise            
//      maxValue : Maximum coordinate value (NOT float64)
//                 -When you queried a noise value with coodinate
//                  which is greater than maxValue or less than 0.0,
//                  the return value will be looped.
    nz := noise.New(100, 256)
    
    
// noise.Get(x, y float64) float64
//  - Get noise value with coordinate(x, y)
    r1 := nz.Get(200.2, 81.4)
    
// noise.GetOctaved(x, y float64, octaves int, persistence float64) float64
//  - Get octaved noise value with coordinate(x, y) and parameters(octaves, persistence)
    r2 := nz.GetOctaved(200.2, 81.4, 10, 0.5)
 
}
```

## References

 - **K. Perlin : Improving noise** <br>
  ACM Transactions on Graphics <br>
  Volume 21, Issue 3, July 2002, pp 681–682<br>
  https://doi.org/10.1145/566654.566636

## LICENSE

MIT License<br>
Copyright (c) 2022 Tada Teruki (多田 瑛貴)
