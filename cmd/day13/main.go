package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bowmanrogere/advent-of-code-2020/internal"
)

func main() {
	filename := fmt.Sprintf("%s/Development/advent-of-code-2020/cmd/day13/bus-info.txt", os.Getenv("HOME"))

	busInfo, err := internal.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	puzzle1(busInfo)
}

func puzzle1(busInfo []string) {
	departureTime, _ := strconv.Atoi(busInfo[0])
	bussesString := busInfo[1]

	busses := make([]int, 0)

	for _, b := range strings.Split(bussesString, ",") {
		if b != "x" {
			bus, _ := strconv.Atoi(b)
			busses = append(busses, bus)
		}
	}

	waitTimes := make(map[int]int)
	for _, bus := range busses {
		waitTimes[bus] = bus*(departureTime/bus+1) - departureTime
	}

	leastWait := -1
	bus := -1

	for k, v := range waitTimes {
		if leastWait == -1 || v < leastWait {
			leastWait = v
			bus = k
		}
	}

	println(fmt.Sprintf("Puzzle #1: %d", bus*leastWait))
}
