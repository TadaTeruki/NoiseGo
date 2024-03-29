# NoiseGo

Implementation of [**Improved Perlin Noise**](https://doi.org/10.1145/566654.566636) in Golang. <br>
<br>

## Preview

<img src="https://user-images.githubusercontent.com/69315285/150670088-e690b5f4-b15f-4950-b959-a143277994f1.png" height="300"> <img src="https://user-images.githubusercontent.com/69315285/150670142-3a5530cf-30f3-4dd7-9d34-36e9aec0c988.png" height="300">


## Usage

```go

import(
    "github.com/TadaTeruki/NoiseGo/v2/noise"
)

func main(){
//  noise.New(seed int64, maxValue uint64) noise.Noise             
//   - Initialize noise generator 
//      maxValue : maximum coordinate value
    nz := noise.New(255)
    
    
// noise.Get(x, y float64) float64
//  - Get noise value with coordinate(x, y)
    r1 := nz.Get(200.2, 81.4)
    
// noise.GetOctaved(x, y float64, octaves int, persistence float64) float64
//  - Get octaved noise value with coordinate(x, y) and parameters(octaves, persistence)
    r2 := nz.GetOctaved(200.2, 81.4, 10, 0.5)
 
}
```

|maxValue = 30|maxValue = 3|
|---|---|
|<img src="https://user-images.githubusercontent.com/69315285/150671923-ce22fbfd-6397-456f-bedc-d9823f9a6bf9.png" height="180">|<img src="https://user-images.githubusercontent.com/69315285/150671993-3256f67b-be14-4a1c-bc5f-9dffcce59e50.png" height="180">|

## Installation


```
 $ go get -u github.com/TadaTeruki/NoiseGo/v2
```
