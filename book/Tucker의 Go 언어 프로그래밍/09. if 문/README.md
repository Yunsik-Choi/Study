# if 문 사용법

``` go
package main

import "fmt"

func main() {
    temp := 33

    if temp > 28 {
        fmt.Println("에어컨을 켠다")
    } else if temp <= 3 {
        fmt.Println("히터를 켠다")
    } else {
        fmt.Println("대기한다")
    }
}
```

## 쇼트서킷
&& 연산은 좌변이 false 이면 우변을 검사하지 않고 false 처리를 합니다.
|| 연산 역시 좌변이 true이면 우변은 검사하지 않고 true 처리를 합니다.
이를 쇼트서킷이라고 합니다.

``` go
package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
    cnt++
    return cnt
}

func main() {
    if false && IncreaseAndReturn() < 5 { // 함수가 호출되지 않습니다.
        fmt.Println("1 증가")
    }

    if true || IncreaseAndReturn() < 5 { // 함수가 호출되지 않습니다.
        fmt.Println("2 증가")
    }

    fmt.Println(cnt) // 0
}
```

## if 초기문; 조건문

if 문 조건을 검사힉 전에 초기문을 넣을 수 있습니다.
초기문은 검사에 사용할 변수를 초기화할때 주로 사용합니다.

``` go
if 초기문; 조건문 {
    문장
}
```

``` go
package main

import "fmt"

func getMyAge() (int, bool) {
    return 33, true
}

func main() {
    if age, ok := getMyAge(); ok && age < 20 {
        fmt.Println(age)
    } else if ok && age < 30 { // age 값에 접근 가능
        fmt.Println(age)
    } else if ok {
        fmt.Println(age)
    } else {
        fmt.Println("Error")
    }

    fmt.Println(age) // Error - age는 소멸되었음
}
```

초기문에서 선언한 변수의 범위는 if문 안으로 한정된다는 사실에 주의해야 합니다.
일반적으로 변수는 변수가 선언된 중괄호를 벗어나면 소멸되지만 if문 초기문에 선언된 변수는 if문이 종료되기 전까지 유지됩니다.