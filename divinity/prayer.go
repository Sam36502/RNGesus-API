package divinity

import "math/rand"

const (
	TIME_TO_FORGET = 12
)

type FloatPrayer struct {
	Count uint64
	Age   int
}

var floatPrayers = make(map[float64]FloatPrayer)

func PrayForFloat(f float64) {
	p, exists := floatPrayers[f]
	if exists {
		p.Count++
		p.Age = 0
	} else {
		p = FloatPrayer{
			Count: 1,
			Age:   0,
		}
	}
}

// Goes through each prayer increasing how long
// it's been waiting. Eventually God forgets it.
func agePrayers() {
	for n, p := range floatPrayers {
		p.Age++
		floatPrayers[n] = p
	}
}

// Returns which float has the most prayers in
// favour of it and how many prayers it's got
func mostBelovedFloat() (float64, uint64) {
	var favourite float64
	var favPrayers uint64
	var n uint64 = 0
	for f, p := range floatPrayers {
		if p.Count > n {
			n = p.Count
			favourite = f
			favPrayers = n
		}
	}
	return favourite, favPrayers
}

// Takes a float and how many prayers it has
// and returns close values relative to the
// number of prayers
func RandomFloatWithPrayers(f float64, p uint64) float64 {
	if uint64(rand.Intn(1000)) < p {
		return f
	}

	if uint64(rand.Intn(100)) < p {
		return f + (randomSign() * rand.Float64() * 0.1)
	}

	if uint64(rand.Intn(10)) < p {
		return f + (randomSign() * rand.Float64() * 0.01)
	}

	return rand.Float64()
}

// Picks either 1 or -1
func randomSign() float64 {
	return float64(rand.Intn(2))*2 - 1
}
