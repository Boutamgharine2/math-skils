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
	split := strings.Split(c, "\n")
	for i := 0; i < len(split); i++ {
		e, err := strconv.Atoi(split[i])
		if err != nil {
			fmt.Println(err)
		}
		D = append(D, e)
	}

	fmt.Printf("Average: %d\n", Mathematics.Average(D))
	fmt.Printf("Median: %d \n", Mathematics.Median(D))
	fmt.Printf("Variance: %d \n", Mathematics.Variance(D, Mathematics.Average(D)))
	fmt.Printf("Standard Deviation: %d \n",int(math.Sqrt(float64(Mathematics.Variance(D, Mathematics.Average(D))))))
}
