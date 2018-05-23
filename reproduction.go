// Completely unrealistic reproduction simulator
// Author: Jarod Knoten
//

package main

import (
	"fmt"

	"github.com/jarodknoten/reproduction/producer"
)

func main() {

	// Counters
	elapsedDays := 1

	//Options
	simDays := 10
	//maxChildren := 5
	maxAge := 100
	//maxBirthAge := 50

	for i := 0; i < simDays; i++ {
		var p = producer.New("Test", 1, true, maxAge)
		p.Dump()
		fmt.Printf("%d days have elapsed", elapsedDays)
		elapsedDays++
	}

}
