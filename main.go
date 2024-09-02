package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	Mathematics "Mathematics/funcs"
)

func main() {
	var D []int
	v := os.Args[0:]
	if len(v) != 2 {
		return
	}
	file, err := os.ReadFile(v[1])
	if err != nil {
		fmt.Println(err)
	}
	c := string(file)
	if len(c) < 1 {
		return
	}
	split := strings.Split(c, "\n")
	if split[len(split)-1] == "" {
		split = split[:len(split)-1]
	}
	for i := 0; i < len(split); i++ {
		if split[i] != "" {
			e, err := strconv.Atoi(split[i])
			if err != nil {
				fmt.Println("invalid number!!")
				return
			}
			D = append(D, e)
		} else {
			continue
		}
	}

	fmt.Printf("Average: %d\n", Mathematics.Average(D))
	fmt.Printf("Median: %d \n", Mathematics.Median(D))
	fmt.Printf("Variance: %d \n", Mathematics.Variance(D, Mathematics.Average(D)))
	fmt.Printf("Standard Deviation: %d \n", int(math.Round(math.Sqrt(float64(Mathematics.Variance(D, Mathematics.Average(D)))))))
}
