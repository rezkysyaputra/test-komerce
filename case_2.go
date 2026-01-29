package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	TransportFamilyPSBB(5, []int{1, 2, 4, 3, 3})
	TransportFamilyPSBB(8, []int{2, 3, 4, 4, 2, 1, 3, 1})
	TransportFamilyPSBB(5, []int{1, 5})
}

func TransportFamilyPSBB(families int, space []int) {
	if families != len(space) {
		fmt.Println("Input must be equal with count of family")
		return
	}
	familyLeft := make(map[int]bool)
	totalBus := 0

	sort.Slice(space, func(i, j int) bool {
		return space[i] > space[j]
	})

	for i, v := range space {
		if familyLeft[i] == true {
			continue
		}

		busCapacity := 4
		busCapacity -= v
		familyLeft[i] = true

		for j := i + 1; j < len(space); j++ {
			if !familyLeft[j] && space[j] <= busCapacity {
				busCapacity -= space[j]
				familyLeft[j] = true
			}
		}

		totalBus++
	}

	fmt.Println("Minimum bus required is: " + strconv.Itoa(totalBus))
}
