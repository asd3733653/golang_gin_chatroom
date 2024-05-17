package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

// MinHeap 實現了最小堆的介面
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// TopN 返回數組中的前 N 個最大值
func TopN(nums []int, N int) []int {
	h := &MinHeap{}

	// 將前 N 個數據添加到最小堆中
	for i := 0; i < N; i++ {
		*h = append(*h, nums[i])
	}

	// 建立最小堆
	heap.Init(h)

	// 對剩餘的數據進行處理
	for i := N; i < len(nums); i++ {
		if nums[i] > (*h)[0] {
			(*h)[0] = nums[i]
			heap.Fix(h, 0)
		}
	}

	// 將最小堆中的元素轉換為結果並返回
	result := make([]int, N)
	for i := 0; i < N; i++ {
		result[i] = heap.Pop(h).(int)
	}
	return result
}

func main() {
	// 生成一億個隨機數
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	const N = 100000000
	nums := make([]int, N)
	for i := 0; i < N; i++ {
		nums[i] = rng.Intn(N) + 1 // 此處假設數據為 1 到 N 的連續數字
	}

	// 取 top 10
	top10 := TopN(nums, 10)
	fmt.Println("Top 10:", top10)
}
