# 함수 정의

``` go
package main

import "fmt"

func Add(a int, b int) int {
    return a + b
}

func main() {
    c := Add(3, 6)
    fmt.Println(c) // 9
}
```

- 반환하는 값이 있으면 타입을 적고 아니면 비워둡니다.
- 중괄호로 함수 코드 블록을 표시합니다. Go 언어에는 함수 코드 블록의 시작을 알리는 중괄호가 함수를 정의하는 라인과 항상 같은 줄에 있어야 합니다.

## 멀티 함수 반환

``` go
package main

import "fmt"

func Divide(a, b int) (int, bool) {
    if b == 0 {
        return 0, false
    }
    return a / b, true
}
```

함수는 값을 여러 개 반환할 수 있습니다.
반환값이 여럿일 때는 반환 타입들을 소괄호로 묶어서 표현합니다.

## 변수명을 지정해 반환하기

``` go
package main

import "fmt"

func Divide(a, b int) (result int, success bool) {
    if b == 0 {
        result = 0
        success = false
        return
    }
    result = a / b
    success = true
    return
}
```

함수 선언부에 반환 타입을 적을 때 변수명까지 지정해주면 return 문으로 해당 변수를 명시적으로 반환하지 않아도 값을 반환할 수 있습니다.