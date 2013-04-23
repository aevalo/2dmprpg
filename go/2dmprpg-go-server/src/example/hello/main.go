// Hello is a trivial example of a main package.
package main

import (
  "example/newmath"
  "fmt"
)

func main() {
  foobar()
  fmt.Printf("Sqrt(2) = %v\n", newmath.Sqrt(2))
}
