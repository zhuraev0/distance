package main

import "testing"

func Test_distance(t *testing.T) {
	tests := []struct {
		name         string
		fuelConsumed int
		fuelFullFuel int
		want         int
	}{
		{"> hundred kilometres", 16, 24, 150},
		{"hundred km", 8, 8, 100},
		{"< hundred km", 9, 2, 22},
	}
	for _, test := range tests {
		got := fuelPerHundredKm(test.fuelConsumed, test.fuelFullFuel)
		if got != test.want {
			t.Error("Test fail:", test.name, "got:", got, "want:", test.want)
		}
	}
}