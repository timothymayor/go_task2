package main

import (
	"fmt"
	"math"
)

//Converter type
type Converter struct{}

// Feet type
type Feet float64

// Centimeter type
type Centimeter float64

//Seconds type
type Seconds float64

//Minutes type
type Minutes float64

//Celsius type
type Celsius float64

//Fahrenheit type
type Fahrenheit float64

//Kilograms type
type Kilograms float64

//Pounds type
type Pounds float64

//Radian type
type Radian float64

//Degree type
type Degree float64

//CentimeterToFeet func
func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	//code here
	return Feet(c / 30.48)
}

//FeetToCentimeter func
func (cvr Converter) FeetToCentimeter(c Feet) Centimeter {
	//code here
	return Centimeter(c * 30.48)
}

//SecondsToMinutes func
func (cvr Converter) SecondsToMinutes(c Seconds) Minutes {
	//code here
	return Minutes(c / 60)
}

//MinutesToSeconds func
func (cvr Converter) MinutesToSeconds(c Minutes) Seconds {
	//code here
	return Seconds(c * 60)
}

//CelsiusToFahrenheit func
func (cvr Converter) CelsiusToFahrenheit(c Celsius) Fahrenheit {
	//code here
	return Fahrenheit(c*9/5 + 32)
}

//FahrenheitToCelsius func
func (cvr Converter) FahrenheitToCelsius(c Fahrenheit) Celsius {
	//code here
	return Celsius((c - 32) * 5 / 9)
}

//KilogramsToPounds func
func (cvr Converter) KilogramsToPounds(c Kilograms) Pounds {
	//code here
	return Pounds(c * 2.2046)
}

//PoundsToKilograms func
func (cvr Converter) PoundsToKilograms(c Pounds) Kilograms {
	//code here
	return Kilograms(c / 2.2046)
}

//RadianToDegree func
func (cvr Converter) RadianToDegree(c Radian) Degree {

	return Degree(c * (180 / math.Pi))
}

//DegreeToRadian func
func (cvr Converter) DegreeToRadian(c Degree) Radian {

	return Radian(c * (math.Pi / 180))
}

func main() {
	cvr := Converter{}

	fmt.Println("100 Centimeter to Feet is : ", cvr.CentimeterToFeet(100), "Feets")
	fmt.Println("3.2808398950131235 Feet To Centimeter is : ", cvr.FeetToCentimeter(3.2808398950131235), "Centimeters \n")

	fmt.Println("60 Seconds to Minutes is : ", cvr.SecondsToMinutes(60), "Minutes")
	fmt.Println("30 Minutes to Seconds is : ", cvr.MinutesToSeconds(30), "Seconds \n")

	fmt.Println("30 Celsius to Fahrenheit is : ", cvr.CelsiusToFahrenheit(30), "Fahrenheit")
	fmt.Println("89 Fahrenheit to Celsius is : ", cvr.FahrenheitToCelsius(89), "Celsius \n")

	fmt.Println("130 Kilograms to Pounds is : ", cvr.KilogramsToPounds(130), "Pounds")
	fmt.Println("250 Pounds to Kilograms is : ", cvr.PoundsToKilograms(250), "Kilograms \n")

	fmt.Println("10 Radian to Degree is : ", cvr.RadianToDegree(10), "Degree")
	fmt.Println("20 Degree to Radian is : ", cvr.DegreeToRadian(20), "Radian")
}
