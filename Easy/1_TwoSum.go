package main

import "log"

func main() {
	input := []int{0, 4, 3, 0}
	log.Println(twoSum(input, 0))
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, 0)
	for ndx, val := range nums {
		numMap[ndx] = val
	}

	var i = 0
	var j = 0
	var sum = 0
	var fin bool

	for fin {
		one := numMap[i]
		j = i + 1

		log.Println(i, one)

		for j < len(numMap) {
			two := numMap[j]
			sum = one + two

			log.Println(j, two, sum)
			if sum == target {
				fin = true
				break
			}

			j++
		}
		if sum == target {
			fin = true
		}

		i++
	}

	return []int{i - 1, j}
}
