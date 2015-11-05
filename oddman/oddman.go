package main

import (
	"bytes"
	"fmt"
)

func main() {
	var nrOfCases int
	var nrOfGuests int
	var a int
	var buffer bytes.Buffer
	fmt.Scanf("%d", &nrOfCases)
	for i := 1; i <= nrOfCases; i++ {
		odd := 0
		fmt.Scanf("%d", &nrOfGuests)
		for j := 0; j < nrOfGuests; j++ {
			fmt.Scanf("%d", &a)
			odd = odd ^ a
		}
		buffer.WriteString(fmt.Sprintf("Case #%d: %d\n", i, odd))
	}
	fmt.Print(buffer.String())
}
