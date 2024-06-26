package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		var temperature, convertedTemperature float64
		var temperatureStr, fromUnit, toUnit string

		fmt.Print("Enter temperature: ")
		fmt.Scanln(&temperatureStr)

		// Check if inputted temperature is a valid value
		temperature, err := strconv.ParseFloat(temperatureStr, 64)
		if err != nil {
			fmt.Println("Temperature is not a valid value. Please try again.")
			continue // Restart the loop
		}

		for {
			fmt.Print("Enter current unit (C, F, K): ")
			fmt.Scanln(&fromUnit)
			fromUnit = strings.ToUpper(fromUnit)

			// Check if fromUnit is a valid value
			if fromUnit != "C" && fromUnit != "F" && fromUnit != "K" {
				fmt.Println("Temperature unit not available. Please try again.")
				continue // Restart the loop
			}
			break // Exit the inner loop once valid input is received
		}

		for {
			fmt.Print("Enter target unit (C, F, K): ")
			fmt.Scanln(&toUnit)
			toUnit = strings.ToUpper(toUnit)

			// Check if toUnit is a valid value
			if toUnit != "C" && toUnit != "F" && toUnit != "K" {
				fmt.Println("Temperature unit not available. Please try again.")
				continue // Restart the loop
			}
			break // Exit the inner loop once valid input is received
		}

		// Function to set the convertedTemperature based on the inputted fromUnit and toUnit
		switch fromUnit {
		case "C":
			if toUnit == "F" {
				convertedTemperature = temperature*9/5 + 32 // Convert from Celsius to Fahrenheit
			} else if toUnit == "K" {
				convertedTemperature = temperature + 273.15 // Convert from Celsius to Kelvin
			} else {
				fmt.Println("Temperature unit not available!")
				return
			}
			fmt.Printf("Converted Temperature: %.2f %s\n", convertedTemperature, toUnit)

		case "F":
			if toUnit == "C" {
				convertedTemperature = (temperature - 32) * 5 / 9 // Convert from Fahrenheit to Celsius
			} else if toUnit == "K" {
				convertedTemperature = (temperature-32)*5/9 + 273.15 // Convert from Fahrenheit to Kelvin
			} else {
				fmt.Println("Temperature unit not available!")
				return
			}
			fmt.Printf("Converted Temperature: %.2f %s\n", convertedTemperature, toUnit)

		case "K":
			if toUnit == "C" {
				convertedTemperature = temperature - 273.15 // Convert from Kelvin to Celsius
			} else if toUnit == "F" {
				convertedTemperature = (temperature-273.15)*9/5 + 32 // Convert from Kelvin to Fahrenheit
			} else {
				fmt.Println("Temperature unit not available!")
				return
			}
			fmt.Printf("Converted Temperature: %.2f %s\n", convertedTemperature, toUnit)
		}
		break // Exit the outer loop once valid input is received
	}
}
