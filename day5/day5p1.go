package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var reSeeds = regexp.MustCompile(`seeds: ([ 0-9]*)$`)
var reMapDeclaration = regexp.MustCompile("^([-a-z]+) map:")
var reMapInputs = regexp.MustCompile(`^([0-9]+)\s([0-9]+)\s([0-9]+)$`)

type rng struct {
	start  int
	length int
}

func newRng(start int, length int) *rng {
	return &rng{start, length}
}

type srcDstMap struct {
	src     string
	dst     string
	nextMap *srcDstMap
	srcRngs []*rng
	dstRngs []*rng
}

func newSrcDstMap(src string, dst string, m *srcDstMap) *srcDstMap {
	return &srcDstMap{src, dst, m, []*rng{{}}, []*rng{{}}}
}

func (t *srcDstMap) addRanges(srcStart int, dstStart int, length int) {
	t.srcRngs = append(t.srcRngs, newRng(srcStart, length))
	t.dstRngs = append(t.dstRngs, newRng(dstStart, length))
}

func (t *srcDstMap) getValue(n int) int {
	for i, rng := range t.srcRngs {
		if n >= rng.start && n < rng.start+rng.length {
			return t.dstRngs[i].start + (n - rng.start)
		}
	}
	return n
}

func (t *srcDstMap) getNextValue(n int) int {
	val := t.getValue(n)
	if t.nextMap == nil {
		return val
	}
	return t.nextMap.getNextValue(val)
}

func getSeedLocation(n int) int {
	return seedToSoil.getNextValue(n)
}

var humidityToLocation = newSrcDstMap("humidity", "location", nil)
var temperatureToHumidity = newSrcDstMap("temperature", "humidity", humidityToLocation)
var lightToTemperature = newSrcDstMap("light", "temperature", temperatureToHumidity)
var waterToLight = newSrcDstMap("water", "light", lightToTemperature)
var fertilizerToWater = newSrcDstMap("fertilizer", "water", waterToLight)
var soilToFertilizer = newSrcDstMap("soil", "fertilizer", fertilizerToWater)
var seedToSoil = newSrcDstMap("seed", "soil", soilToFertilizer)

var typeMap = map[string]*srcDstMap{
	"seed-to-soil":            seedToSoil,
	"soil-to-fertilizer":      soilToFertilizer,
	"fertilizer-to-water":     fertilizerToWater,
	"water-to-light":          waterToLight,
	"light-to-temperature":    lightToTemperature,
	"temperature-to-humidity": temperatureToHumidity,
	"humidity-to-location":    humidityToLocation,
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var currentMap *srcDstMap
	scanner := bufio.NewScanner(file)
	var inputSeedValues []int
	for scanner.Scan() {
		line := scanner.Text()

		seedsMatch := reSeeds.FindStringSubmatch(line)
		if seedsMatch != nil {
			seeds := strings.Split(seedsMatch[1], " ")
			for _, seedStr := range seeds {
				seed, err := strconv.Atoi(seedStr)
				if err != nil {
					log.Fatal(err)
				}
				inputSeedValues = append(inputSeedValues, seed)
			}
		}

		mapMatch := reMapDeclaration.FindStringSubmatch(line)
		if mapMatch != nil {
			currentMap = typeMap[mapMatch[1]]
			continue
		}

		mapInputsMatch := reMapInputs.FindStringSubmatch(line)
		if mapInputsMatch != nil {
			dstRangeStart, err := strconv.Atoi(mapInputsMatch[1])
			if err != nil {
				log.Fatal(err)
			}
			srcRangeStart, err := strconv.Atoi(mapInputsMatch[2])
			if err != nil {
				log.Fatal(err)
			}
			rangeLength, err := strconv.Atoi(mapInputsMatch[3])
			if err != nil {
				log.Fatal(err)
			}
			currentMap.addRanges(srcRangeStart, dstRangeStart, rangeLength)
		}
	}

	var locationValues []int
	for _, seed := range inputSeedValues {
		locationValues = append(locationValues, getSeedLocation(seed))
	}
	slices.Sort(locationValues)
	fmt.Println(locationValues[0])
}
