/*
This program demonstrates a simple case using "goto" keyword.
This is a powerful keyword, which should be used carefully.

*/

package main

import "fmt"

func main() {
	stop := 20

	for i := 0; ; i++ {
		for j := 0; j < 5; j++ {
			if stop <= i*j {
				/*
					a simple break here would cause the inner loop to stop,
					but immediately after, the outer loop will restart it ... and this would
					continue forever!
				*/
				fmt.Printf("breaking at (i=%d, j=%d)\n", i, j)
				goto breakFromLoop
			}
		}
	}
breakFromLoop:
}
