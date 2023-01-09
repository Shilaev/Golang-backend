package mymethods

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Rankin float64

// CtoF Fahrenheit
// fahrenheit = 9/5 * tc + 32
func CtoF(celsius Celsius) Fahrenheit {
	return Fahrenheit(9.0/5.0*celsius + 32.0)
}

// FtoC Celsius
// celsius = 5/9 * (tf - 32)
func FtoC(fahrenheit Fahrenheit) Celsius {
	return Celsius(5.0 / 9.0 * (fahrenheit - 32.0))
}

// CtoK Kelvin
// kelvin = tc + 273.15
func CtoK(celsius Celsius) Kelvin {
	return Kelvin(celsius + 273.15)
}

// KtoC Celsius
// celsius = tk - 273.15
func KtoC(kelvin Kelvin) Celsius {
	return Celsius(kelvin - 273.15)
}
