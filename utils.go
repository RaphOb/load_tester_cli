package main

import "time"

func countStatus(status []float32) (int, int) {
	success, failure := 0, 0
	for _, code := range status {
		if int(code) >= 200 && int(code) < 300 {
			success++
		} else {
			failure++
		}
	}
	return success, failure
}

func waitForRate(rate int, freq string) time.Duration {
	var interval time.Duration
	if rate <= 0 {
		return 0
	}
	switch freq {
	case "ms":
		interval = time.Millisecond / time.Duration(rate)
	case "s":
		interval = time.Second / time.Duration(rate)
	case "min":
		interval = time.Minute / time.Duration(rate)
	default:
		interval = 0
	}
	return interval
}
