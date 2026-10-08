package main

func ConvertTemperature(temp float64, scale string) float64 {
	var result float64
	if scale == "C" {
		result = temp*9/5 + 32
	}
	if scale == "F" {
		result = (temp - 32) * 5 / 9
	}
	return result
}
