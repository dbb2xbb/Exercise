package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	RAND_NUM = 1
)

func RandomIndex(randNum int64) []int64 {
	var (
		arr  []int64
		i, j int64
		r    *rand.Rand
	)
	if randNum < 0 || randNum > 0xFFFFFFFF {
		return nil
	}
	arr = make([]int64, randNum)
	for i = 0; i < randNum; i++ {
		arr[i] = i
	}
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i = 0; i < randNum; i++ {
		j = i + r.Int63n(randNum-i)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

type RandomFunc func(num int64) []int64

func CalcRunningTime(randomFunc RandomFunc, params interface{}) {
	var (
		now     time.Time
		escaped time.Duration
		result  []int64
	)
	now = time.Now()
	result = randomFunc(params.(int64))
	escaped = time.Since(now)
	fmt.Println("run time:", escaped)
	fmt.Println("result:", result)
}

func main() {
	CalcRunningTime(RandomIndex, int64(RAND_NUM))
}
