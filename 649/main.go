package main

import "fmt"

// FIXME 这个应该可以优化，不用每次都有去寻找离增加最近的R或者D
//func predictPartyVictory(senate string) string {
//	flag := make([]bool, len(senate))
//	r_n, d_n := 0, 0
//
//	for i := 0; i < len(senate); i++ {
//		if senate[i:i+1] == "R" {
//			r_n++
//		} else {
//			d_n++
//		}
//	}
//	var i = 0
//	var p string
//	for r_n !=0 && d_n!=0 {
//		if flag[i]{
//			i++
//			continue
//		}
//		i%=len(senate)
//		if !flag[i]{
//			p =senate[i:i+1]
//		}
//		for j:=i+1;j<len(senate)+i;j++{
//			j%=len(senate)
//			if senate[j:j+1] != p && !flag[j] {
//				if  p == "R"{
//					d_n--
//					flag[j] = true
//				}else {
//					r_n--
//					flag[j] = true
//				}
//				i++
//				break
//			}
//		}
//
//	}
//
//	if r_n == 0{
//		return "Dire"
//	}else {
//		return "Radiant"
//	}
//}


// 执行用时为 0 ms 的范例
func predictPartyVictory(senate string) string {
	l := len(senate)
	q1, q2 := []int{}, []int{}
	for i := 0; i < l; i++ {
		if senate[i] == 'R' {
			q1 = append(q1, i)
		} else {
			q2 = append(q2, i)
		}
	}
	for 0 != len(q1) && 0 != len(q2) {
		i, j := q1[0], q2[0]
		q1, q2 = q1[1:], q2[1:] // 切分数组
		if i < j {  // 如果R的索引小于D 的索引，则表示R在D的前面，那么则该R会保留下来，同时它的所以会在下次需要循环才会被用到，所以交换到到后面，同时索引加len(senate)
			q1 = append(q1, i+l)
		} else {
			q2 = append(q2, j+l)
		}
	}
	if len(q1) > len(q2) {
		return "Radiant"
	}
	return "Dire"
}

func main() {
	fmt.Println(predictPartyVictory("RDD"))
}
