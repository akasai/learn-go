package main

func main() {
	a := `테스트입니다. \n 안녕하세요?`
	b := "테스트입니다. \n 안녕하세요?"

	println(a)
	// 테스트입니다. \n 안녕하세요?
	println(b)
	// 테스트입니다.
	//  안녕하세요?

	i := 100
	u := uint(i)
	f := float32(i)

	println(u, f) // 100 +1.000000e+002

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2) // [3/32]0xc00003c740 ABC
}
