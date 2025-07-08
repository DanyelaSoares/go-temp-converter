package converter

func CelsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func CelsiusToKelvin(c float64) float64 {
	return c + 273.15
}

func KelvinToCelsius(k float64) float64 {
	return k - 273.15
}
