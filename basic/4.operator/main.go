package main

func main() {
  // 산술 연산
  var a = 1
  a++                 // 2
  var b = (a * 5) / 2 // 5
  println(a, b)

  // 비교 연산
  var c = "A"
  var d = "A"
  println(c == d) // true
  println(c != d) // false
  println(a < b) // true

  // 논리 연산
  var e = true
  var f = false
  println(e && f) // false
  println(e || f) // true
  println(!f) // true

  // 비트 연산
  var g = 1
  var h = 1
  println(g & h) // 1
  println(g | h) // 1
  println(g << h) // 2

  // 할당 연산
  var k int
  k = 100
  k /= 2
  k >>= 2
  k |= 2

  // 포인터 연산
  var l = 10
  var m = &l
  println(l, m, *m) // 10 0xc00003c768 10
}
