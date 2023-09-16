package util

import (
	"math/rand"
)

func RandomSliceOfInt (maxIndex int) []int{
	randNum := []int{}

	for i := 0; i < maxIndex ; i++ {
		randNum = append(randNum, rand.Int())
	}
	return randNum
}