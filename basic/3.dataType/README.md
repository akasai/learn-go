# 데이터 타입

 ## 데이터 타입의 종류

1. 불린
    
    `bool`

2. 문자열

    `string` (immutable)
 
3. 정수형

    `int`, `int8`, `int16`, `int32`, `int64`,
 
    `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
 
4. Float 및 복소수
    
    `float32`, `float64`, `complex64`, `complex128`
 
5. 기타

    `byte`: `unit8`과 동일하며 바이트코드에 사용
 
    `rune`: `int32`와 동일하며 유니코드 코드포인트에 사용

### 문자열

문자열 리터럴은 \` \`(Back Quote), " "(따옴표)를 사용.

1. \`\` 
   
   Back Quote로 둘러 쌓인 문자열은 `Raw String Literal`이라고 한다.

   `escape문자열`은 동작하지 않으며 순수한 형태의 string을 사용한다.

   **주로 복수라인의 문자열**을 선언할 때 사용한다.
 
2. " "

   따옴표로 둘러 쌓인 문자열은 `Interpreted String Literal`이라고 한다.

   일반적인 문자열 선언과 같다.

### 데이터 타입 변환 (Type Conversion)

다른 데이터 타입으로의 변환은 `T(v)`라고 표현하며 T는 변환하고자 하는 타입, v는 변환될 값을 의미한다.

```go
a := 1
b := float32(1) // T(v)
```

문자열을 바이트배열로 변경할 땐 `[]byte("ABC")`로 표현할 수 있다.

> **주의할 점**
> 
> Go에서 타입간 변환은 명시적으로 지정해야한다.
> 
> int에서 uint로 변환 시, 암묵적(implicit) 변환은 일어나지 않는다.
> 
> 명시적 지정이 없다면 **런타임에러가 발생한다.**
> 
> **즉, 암묵적 형변환을 지원하지 않는다.**