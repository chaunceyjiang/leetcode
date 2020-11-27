package main

func fourSumCount(A []int, B []int, C []int, D []int) int {
	m := make(map[int]int)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			m[A[i]+B[j]]++
		}
	}
	total := 0
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			if _, ok := m[-(C[i] + D[j])]; ok {
				total++
			}
		}
	}
	return total
}

func main() {

}
