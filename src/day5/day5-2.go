package day5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day5p2() {
	fileSeeds, _ := os.Open("./input/day-5-seeds.txt")
	scannerSeeds := bufio.NewScanner(fileSeeds)
	scannerSeeds.Scan()
	seedsBytes := scannerSeeds.Bytes()
	textSeeds := make([]byte, len(seedsBytes))

	copy(textSeeds, seedsBytes)

	fileMaps, _ := os.Open("./input/day-5-maps.txt")
	scannerMaps := bufio.NewScanner(fileMaps)
	textMaps := make([][]byte, 0)

	for scannerMaps.Scan() {
		src := scannerMaps.Bytes()
		dst := make([]byte, len(src))
		copy(dst, src)
		textMaps = append(textMaps, dst)
	}

	seed2soilMap := createMaps(textMaps, 1, 48)
	soil2fertMap := createMaps(textMaps, 51, 83)
	fert2waterMap := createMaps(textMaps, 86, 95)
	water2lightMap := createMaps(textMaps, 98, 109)
	light2tempMap := createMaps(textMaps, 112, 134)
	temp2humMap := createMaps(textMaps, 137, 174)
	hum2locMap := createMaps(textMaps, 177, 210)

	seedStrings := string(seedsBytes)
	seedsSplits := strings.Split(seedStrings, " ")

	minSeed := math.MaxInt64
	for i := 0; i < len(seedsSplits)-1; i += 2 {
		seed, _ := strconv.Atoi(seedsSplits[i])
		seedRange, _ := strconv.Atoi(seedsSplits[i+1])
		for j := 0; j < seedRange; j++ {
			currSeed := seed + j
			soil := useMaps(seed2soilMap, currSeed)
			fert := useMaps(soil2fertMap, soil)
			water := useMaps(fert2waterMap, fert)
			light := useMaps(water2lightMap, water)
			temp := useMaps(light2tempMap, light)
			hum := useMaps(temp2humMap, temp)
			loc := useMaps(hum2locMap, hum)

			minSeed = min(minSeed, loc)
		}

	}

	fmt.Println(minSeed)

}

func getSeedsRanges(line []byte) []int {
	seedsBytes := string(line)
	seedsStrings := strings.Split(seedsBytes, " ")
	seeds := make([]int, 0)
	for i := 0; i < len(seedsStrings)-1; i += 2 {
		seed, _ := strconv.Atoi(seedsStrings[i])
		seedRange, _ := strconv.Atoi(seedsStrings[i+1])
		seeds = append(seeds, seed)
		seeds = append(seeds, seed+seedRange)
	}
	return seeds
}
