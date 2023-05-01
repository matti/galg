package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/montanaflynn/stats"
)

func main() {

	var numbers []float64

	for {
		reader := bufio.NewReader(os.Stdin)
		bytes, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}

		var numeric float64
		if n, err := strconv.ParseFloat(string(bytes), 64); err != nil {
			panic(err)
		} else {
			numeric = n
		}

		numbers = append(numbers, numeric)

		mean, _ := stats.Mean(numbers)
		fmt.Println(math.Floor(mean*10) / 10)
	}
}
