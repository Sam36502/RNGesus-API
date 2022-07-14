package divinity

import "math/rand"

func RandomFloat() (float64, error) {
	err := InitialiseRNG()
	if err != nil {
		return 0, err
	}
	return rand.Float64(), nil
}

func RandomInt(min, max int64) (int64, error) {
	flt, err := RandomFloat()
	if err != nil {
		return 0, err
	}
	return int64(flt*float64(max-min+1)) + min, nil
}
