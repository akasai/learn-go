# 조건문

다른 언어가 지원하는 조건문은 동일하게 지원한다.

## if 문

go에서는 if문의 조건은 괄호를 생략해도 되며, 중괄호(brace: `{`)를 if문과 같은 라인에 두어야한다.

**조건은 반드시 `boolean`이어야 한다.**

```go
if k == 1 {
  println(`one`)
}
```

else if 문 역시 동일하게 적용한다.

```go
if k == 1 {
  println(`one`)
} else if k == 2 {
  println(`two`)
} else {
  println(`other`)
}
```

### optional statement

if 문의 조건절에서 간단한 문장과 함께 사용할 수 있다.

단, `if문 scope`안에서만 사용가능하다.

| if 문 뿐만 아니라 다양한 조건문에서 활용 가능하다.

```go
if val := 2; val < max {
  println(`success`)
}
val++ // error
```

## switch 문

다른 언어와 동일한 문법 형태를 갖는다.

if문 처럼 expression의 괄호는 생략 가능하며, 복수의 case의 경우 `,(콤마)`로 나열한다.

```go
switch val {
  case 1:
    println(`one`)
  case 2:
    println(`two`)
  case 3, 4:
    println(`three, four`)
}

// val
switch val := 1; val {
  // ...
}
```

switch문에는 go만의 특별한 용법이 있다. 

### switch 뒤에 expression이 없을 수 있음, case문에 expression을 쓸 수 있음

기준을 정의하지 않고 switch문을 사용할 수 있다.

```go
func grade(score int) {
  switch {
    case score > 90:
      println(`A`)
    case score > 90:
      println(`B`)
    default:
      println(`out`)
  }
}
```

### No default fall through

일반적으로 switch문은 case블럭에서 break문을 통해 빠져나올 수 있다.

만약 break문이 없다면 이후의 case블럭을 모두 수행한다.

go는 break문과 관계없이 항상 switch문을 빠져나온다. (go컴파일러가 자동으로 break문을 추가하기 때문)

만약 break와 관계없이 case블럭들을 추가적으로 수행하려면 fallthrough문을 명시해주면 된다.

```go
switch val {
    case 1:
      println(`1`)
      fallthrough
    case 2:
      println(`2`)
      fallthrough
    case 3:
      println(`3`)
      fallthrough
    default:
      println(`default`)
}
```
```go
// result
1
2
3
default
```

### Type switch

변수의 타입을 체크할 수 있다.

```go
switch v.(type) {
  case int:
    println(`int`)
  case bool:
    println(`bool`)
  // ...
}
```