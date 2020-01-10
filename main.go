package main

func main() {

}

func fuelPerHundredKm(fuelConsumed int, fuelFullFuel int) int {
	const distancePerHundredKm int = 100
	calculate := distancePerHundredKm * fuelFullFuel / fuelConsumed
	return calculate
}