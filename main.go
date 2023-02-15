package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/bjornarsk/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var out string
var funfacts string
var fahr float64
var celsius float64
var kelvin float64

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i garder celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader kelvin")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func formatNumber(num float64) string {
	formatted := strconv.FormatFloat(num, 'f', 2, 64)
	formatted = strings.TrimRight(formatted, "0")
	formatted = strings.TrimRight(formatted, ".")

	if (num > 1000) || (num < -1000) {
		// Add commas as separators between large numbers
		formatted = addCommas(formatted)
		// Add a space between the number and the comma
		formatted = strings.Replace(formatted, ",", " ", -1)
	}
	return formatted
}

func addCommas(str string) string {
	n := len(str)
	if n <= 3 {
		return str
	}
	return addCommas(str[:n-3]) + "," + str[n-3:]
}

func main() {

	flag.Parse()

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.
	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.
	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())
	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	fmt.Println("Dette er første linje")

	// Eksempel på enkel logikk, Fahrenheit til Celsius
	if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		celsius = conv.FahrenheitToCelsius(fahr)
		// skal returnere °C
		fmt.Println(fahr, "°F er", (formatNumber(celsius)), "°C")
	}

	// Celsius til Fahrenheit
	if out == "F" && isFlagPassed("C") {
		fahr = conv.CelsiusToFahrenheit(celsius)
		fmt.Println(celsius, "°C er", fahr, "°F")
	}

	// Fahrenheit til Kelvin
	if out == "K" && isFlagPassed("F") {
		kelvin = conv.FahrenheitToKelvin(fahr)
		fmt.Println(fahr, "°F er", kelvin, "°K")
	}

	// Celsius til Kelvin
	if out == "K" && isFlagPassed("C") {
		kelvin = conv.CelsiusToKelvin(celsius)
		fmt.Println(celsius, "°C er", kelvin, "°K")
	}

	// Kelvin til Celsius
	if out == "C" && isFlagPassed("K") {
		celsius = conv.KelvinToCelsius(kelvin)
		fmt.Println(kelvin, "°K er", celsius, "°C")
	}

	// Kelvin til Fahrenheit
	if out == "F" && isFlagPassed("K") {
		fahr = conv.KelvinToFahrenheit(kelvin)
		fmt.Println(kelvin, "°K er", fahr, "°F")
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
