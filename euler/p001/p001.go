// https://projecteuler.net/problem=1
//
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we
// get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.

package euler

import (
  "math"
  "fmt"
)

func outBy5(i int) string {
  if i == 0 {
    return "div5 --> [ X ]"
  } else {
    return "div5 --> [   ]"
  }
}

func outBy3(i int) string {
  if i == 0 {
    return "div3 --> [ X ]"
  } else {
    return "div3 --> [   ]"
  }
}

// Extremely straightforward implementation
func Simple(max int, debug bool) int {
  total := 0
  for i := 0; i < max; i++ {
    by5 := int(math.Mod(float64(i), 5))
    by3 := int(math.Mod(float64(i), 3))

    if debug {
      fmt.Printf("i is %d ( %s  %s )", i, outBy5(by5), outBy3(by3))
    }

    if by5 == 0 || by3 == 0 {
      total += i
      if (debug) {
        fmt.Printf("\t[TOTAL] %d (+%d)", total, i)
      }
    }

    if (debug) {
      fmt.Printf("\n")
    }
  }

  return total
}
