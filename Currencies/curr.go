package Currencies

func Conv(con string, con2 string, amount float64) float64 {
	valsDO := map[string]float64{
		"EUR": 1.17,
		"GBP": 1.35,
		"JPY": 0.006,
		"CAD": 0.72,
		"CHF": 1.26,
		"CNY": 0.14,
		"INR": 0.01,
		"HKD": 0.12,
		"EGP": 0.02,
		"USD": 1.00,
	}
	vals2 := map[string]float64{
		"EUR": 0.85,
		"GBP": 0.74,
		"JPY": 147.5,
		"CAD": 1.39,
		"CHF": 0.79,
		"CNY": 7.12,
		"INR": 88.7,
		"HKD": 7.78,
		"EGP": 47.7,
		"USD": 1.00,
	}
	out := valsDO[con] * amount * vals2[con2]

	return out
}
