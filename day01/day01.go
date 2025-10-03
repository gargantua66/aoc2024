package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input")
	numbers := strings.Split(string(content), "\n")

	map1 := []int{}
	map2 := []int{}

	for _, line := range numbers {
		id := strings.Split(line, "   ")

		map1id, _ := strconv.Atoi(id[0])
		map1 = append(map1, map1id)
		map2id, _ := strconv.Atoi(id[1])
		map2 = append(map2, map2id)
	}

	slices.Sort(map1)
	slices.Sort(map2)

	distance := 0
	for i := 0; i < len(map1); i++ {
		if map1[i] > map2[i] {
			distance = distance + map1[i] - map2[i]
		} else {
			distance = distance + map2[i] - map1[i]
		}

	}
	fmt.Println("Part1: ", distance)

	frequencies := make(map[int]int)
	for _, ditto := range map2 {
		frequencies[ditto] = frequencies[ditto] + 1
	}

	sumnum := 0
	for _, numnum := range map1 {
		sumnum = sumnum + frequencies[numnum]*numnum
	}
	fmt.Println("Part2: ", sumnum)
}
