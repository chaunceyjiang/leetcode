package main

import "fmt"

func largestPerimeter(A []int) int {
	heapSort(A)
	for i:=len(A)-1;i>1;i--{
		if A[i-1]+A[i-2]>A[i]{
			return A[i-1]+A[i-2]+A[i]
		}
	}
	return 0
}

/*
	大顶堆： nums[i] > nums[2*i+1] && nums[2*i+2]
*/
func heapSort(nums []int) {

	// 1.构建 大顶堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		// len(nums)/2-1 表示第一个非叶子节点
		adjustHeap(nums, i, len(nums)) // 将每一个非叶子节点，都进行调整，让这个数组都满足整堆特性。
	}
	// 利用堆 进行排序，最后得到一个排好序的数组
	for j := len(nums) - 1; j > 0; j-- {
		// 将堆顶元素与数组最后一个元素交换
		nums[0], nums[j] = nums[j], nums[0]
		//移动后，重新对进行调整，以保持堆特效,重新调整时，排除了最后一个元素
		adjustHeap(nums, 0, j)
	}
}

func adjustHeap(nums []int, i, length int) {
	temp := nums[i] // 保存当前节点数值
	// i 当前节点  length 当前堆数组长度，长度会随着堆的排序，越来越小
	// 自顶向下 调整
	for k := i*2 + 1; k < length; k = 2*k + 1 {
		// k 标示 i 的左节点
		// k+1 标示 i 的右节点
		// 先在左右节点比较
		if k+1 < length && nums[k] < nums[k+1] {
			// 如果左节点小于右节点,则将k 调整为右节点，表示右节点大
			k = k + 1
		}
		// 子节点中最大的节点，再跟父节点比较

		if nums[k] > temp {
			// 则将最大子节点的值，调整到父节点
			nums[i] = nums[k]
			i = k // 将当前指针，移动到最大子节点
		} else {
			// 表示当前节点做为父节点已经是最大的了，意味着最小下面，一定没有更大的数值了，当前调整结束，退出循环
			break
		}
	}
	// 将值放到最终的位置
	nums[i] = temp
}

func main() {
	fmt.Println(largestPerimeter([]int{3,6,2,3}))
}
