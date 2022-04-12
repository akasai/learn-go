package main

func main() {
	ifCondition(1)
  optionalStatement(1)
  switchCondition(1)
  switchCondition(2)
  switchCondition1(1)
  typeCheck(true)
  useFallthrough(1)
}

/**
기본 if 문
 */
func ifCondition(val int) {
	if val == 1 {
		println("1")
	} else if val == 2 {
		println("2")
	} else {
		println("nothing")
	}
}

/**
optional statement
 */
func optionalStatement(val int) {
	if i := 1; val+i == 2 {
		println(`2`)
	}
	// i++ // error: i is scoped value
}

/**
switch statement
 */
func switchCondition(val int) {
	switch i := 1; i + val {
	case 1:
		println(`1`)
	case 2, 3:
		println(`2, 3`)
	default:
		println(`default`)
	}
}

/**
switch 용법 1
 */
func switchCondition1(val int) {
  switch {
  case val == 1:
    println(`1`)
  case val == 2:
    println(`2`)
  default:
    println(`default`)
  }
}

/**
switch 용법 2
 */
func typeCheck(val interface{}) {
  switch val.(type) {
  case int:
    println(`int`)
  case bool:
    println(`bool`)
  }
}

/**
fallthrough
 */
func useFallthrough(val int) {
  switch val {
  case 1:
    println(`test1`)
    fallthrough
  case 2:
    println(`test2`)
    fallthrough
  default:
    println(`default`)
  }
}