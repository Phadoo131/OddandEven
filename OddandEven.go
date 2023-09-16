package OddandEven

import (
	"log"
)

func Even (num []int) []int{
	if num == nil {
		log.Fatal("There are no numbers in the slice")
		return nil
	}

	newNum := []int{}

	for _, i := range num {
		if i%2 == 0 {
			newNum = append(newNum, i)
		}
	}
	return newNum
}

func Odd (num []int) []int{
	if num == nil {
		log.Fatal("There are no numbers in the slice")
		return nil
	}

	newNum := []int{}

	for _, i := range num {
		if i%2 == 1 {
			newNum = append(newNum, i)
		}
	}
	return newNum

}

func Calculate(num []int) int {
	answer := 0

	for _, i := range num{
		answer += i
	}
	return answer
}
