package main

import (
	"fmt"

	"github.com/draffensperger/go-interlang/go_to_cxx/swig/point"
)

func main() {
	fmt.Println("Go says: SWIG wrapped so we can calc distance..")

	p1 := point.NewPoint(1.0, 1.0)
	defer point.DeletePoint(p1)

	p2 := point.NewPoint(2.0, 2.0)
	defer point.DeletePoint(p2)

	distance := p1.Distance_to(p2)
	fmt.Printf("Go says: distance is: %v\n", distance)
}
