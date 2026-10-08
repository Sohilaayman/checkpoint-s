package main

func ConvertSeconds(totalSeconds int) []int {
	hours := totalSeconds / 3600
	remainding := totalSeconds % 3600

	minutes := remainding / 60
	seconds := remainding % 60

	return []int{hours, minutes, seconds}
}
