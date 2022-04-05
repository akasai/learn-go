package main

import "fmt"

func main() {
  var l float32 = 11.
  var i, j, k int = 1, 2, 3
  fmt.Println(l, i, j, k)

  var a int
  var b bool
  var c string
  fmt.Println(a, b, c) // zero value: 0, false, ""

  var d = 1
  var e = "HI"
  fmt.Println(d, e) // 타입 추론: 1, HI

  f := 1
  g := "HELLO"
  fmt.Println(f, g)

	const aa int = 1
	fmt.Println(aa)

	const bb = 11
	fmt.Println(bb)

	const (
		cc = 22
		dd = 33
		ee = 44
	)
	fmt.Println(cc, dd, ee)

  const (
    first = iota
    second
    third
  )
  fmt.Println(first, second, third) // 0 1 2
}
