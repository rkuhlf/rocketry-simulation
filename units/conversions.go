package units

func AtmToPa(atm float64) float64 {
	return atm * 101325.0
}

func PaToAtm(pa float64) float64 {
	return pa / 101325.0
}

func FahrenheitToKelvin(f float64) float64 {
	return (f-32)*5/9 + 273.15
}

func MetersToFeet(meters float64) float64 {
	return meters * 3.2808399
}

func FeetToMeters(feet float64) float64 {
	return feet / 3.2808399
}

func InchesToMeters(inches float64) float64 {
	return inches / 12 / 3.2808399
}

func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func KelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}
