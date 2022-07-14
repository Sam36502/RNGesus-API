package divinity

func RandomFloat() (float64, error) {
	err := InitialiseRNG()
	if err != nil {
		return 0, err
	}

	// Add prayer influence
	f := RandomFloatWithPrayers(mostBelovedFloat())
	agePrayers()

	return f, nil
}

func RandomInt(min, max int64) (int64, error) {
	flt, err := RandomFloat()
	if err != nil {
		return 0, err
	}
	return int64(flt*float64(max-min+1)) + min, nil
}
