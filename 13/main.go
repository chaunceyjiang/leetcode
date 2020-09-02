package main

import "fmt"

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	romanMap2 := map[string]int{
		"IV": -2,
		"IX": -2,
		"XL": -20,
		"XC": -20,
		"CD": -200,
		"CM": -200,
	}
	count := 0
	l := len(s)
	for i := 0; i < l; i++ {
		t := s[i : i+1]
		count += romanMap[t]
		if i >= 1 {
			tt := s[i-1 : i+1]
			if v, ok := romanMap2[tt]; ok {
				count += v
			}
		}

	}
	return count
}

func main() {
	fmt.Println(romanToInt("CMXCIX"))
}
