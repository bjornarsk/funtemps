package conv
var Fahrenheit float64
var Celsius float64
var Kelvin float64
/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius2(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return 56.67
}

// De andre konverteringsfunksjonene implementere her
// ...
func FahrenheitToCelsius(value float64) float64 {
    return (value - 32) * (5 / 9)
}

func CelsiusToFahrenheit(value float64) float64 {
    return value(9/5) + 32
}

func CelsiusToKelvin(value float64) float64 {
    return value + 273.15
}

func KelvinToCelsius(value float64) float64 {
    return value - 273.15
}

func KelvinToFahrenheit(value float64) float64 {
    return (value-273.15)(9/5) + 32
}

func FahrenheitToKelvin(value float64) float64 {
    return (value-32)*(5/9) + 273.15
}


