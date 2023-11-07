여러 필드를 묶어서 하나의 구조체를 만듭니다.
배열이 같은 타입의 값들을 변수 하나로 묶어줬던 것과 달리 구조체는 다른 타입의 값들을 변수 하나로 묶어주는 기능입니다.

``` go
type 타입명 struct {
    필드명 타입
    ...
    필드명 타입
}
```

- type 키워드를 적어서 새로운 사용자 정의 타입을 정의할 것을 알립니다.
- 그런 뒤 타입명을 적습니다. 타입명의 첫 번째 글자가 대문자이면 패키지 외부로 공개되는 타입입니다.
- 타입 종류인 struct를 적습니다.


``` go
package main

import "fmt"

type House struct {
    Address string
    Size int
    Price float64
    Type string
}

func main() {
    var house House
    house.Address = "경기도 용인시 ..."
    house.Size = 28
    house.Price = 9.8
    house.Type = "아파트"

    fmt.Println(house.Address) // 경기도 용인시 ...
    fmt.Println(house.Size) // 28
    fmt.Println(house.Price) // 9.8
    fmt.Println(house.Type) // 아파트
}
```

## 구조체 변수 초기화

초깃값을 생략하면 모든 필드가 기본값으로 초기화됩니다.

``` go
var house House
```

모든 필드값을 중괄호 사이에 넣어서 초기화합니다.

``` go
var house House = House {"서울시 강동구", 28, 9.80, "아파트"}
```

``` go
var house House = House {
    "서울시 강동구",
    28,
    9.80,
    "아파트", // 여러 줄로 초기화할 때는 제일 마지막 값에 뒤에 꼭 쉼표를 달아주세요.
}
```

일부 필드값만 초기화할 때는 "필드명: 필드값" 형식으로 초기화합니다.
초기화되지 않은 나머지 변수에는 기본값이 할당됩니다.

``` go
var house House = House{ Size: 28, Type: "아파트"}
```

## 구조체를 포함하는 구조체
구조체의 필드로 다른 구조체를 포함할 수 있습니다.
일반적은 내장 타입처럼 포함하는 방법과 포함된 필드 방식이 있습니다.

``` go
package main

import "fmt"

type Use struct {
    Name string
    ID string
    Age int
}

type VIPUser struct {
    UserInfo User
    VIPLevel int
    Price int
}

func main() {
    user := User{"송하나", "hana", 23}
    vip := VIPUser{
        User{"화랑", "hwarang", 40}
        3,
        250
    }

    fmt.Println(vip.UserInfo.Name) // 화랑
    fmt.Println(vip.UserInfo.ID) // hwarang
    fmt.Println(vip.UserInfo.Age) // 40
}
```

## 포함된 필드 방식
vip에서 Name이나 ID와 같이 UserInfo 안에 속한 필드에 접근하려면 vip.UserInfo.Name과 같이 두 단계를 걸쳐 접근해야 합니다.
구조체에서 다른 구조체를 필드로 포함할 때 필드명을 생략하면 .을 한 번만 찍어 접근할 수 있습니다.

``` go
package main

import "fmt"

type Use struct {
    Name string
    ID string
    Age int
}

type VIPUser struct {
    User
    VIPLevel int
    Price int
}

func main() {
    user := User{"송하나", "hana", 23}
    vip := VIPUser{
        User{"화랑", "hwarang", 40}
        3,
        250
    }

    fmt.Println(vip.Name) // 화랑
    fmt.Println(vip.ID) // hwarang
    fmt.Println(vip.Age) // 40
}
```

이름이 겹칠 경우 현재 변수 타입에 해당하는 구조체의 필드에 접근합니다.

