package main

import "fmt"
import "github.com/morimolymoly/slicer"

func main() {
	numbers := make([]int, 400)
	for i := 0; i < 400; i++ {
		numbers[i] = i
	}
	index := slicer.GetSliceIndex(400, 100)
	for _, i := range index {
		fmt.Printf("start:%d, end:%d\n", i.Start, i.End)
		for _, j := range numbers[i.Start:i.End] {
			fmt.Println(j)
		}
	}
}
