# 도메인 복잡성 관리
프로젝트 성공을 보장하려면 소프트웨어 엔지니어부터 도메인 전문가에 이르기까지 모든 이해관계자가 의사소통에 사용할 수 있는 유비쿼터스 언어를 개발하는 것이 중요하다.
우리의 목표는 유비쿼터스 언어를 사용하여 소프트웨어 설계의 의사결정을 내리는 것이기 때문에 언어는 명확하고 일관성이 있어야 한다.
모호성, 암묵적인 가정, 관련 없는 세부사항이 없어야 한다.
그러나 조직 규모에 따라 도메인 전문가의 멘탈 모델은 일관성이 없을 수 있다.
같은 비지니스 도메인에서도 도메인 전문가마다 서로 다른 모델을 사용할 수 있다.

# 일관성 없는 모델
유비쿼터스 언어가 일관성이 있어야 하지만 모호성은 개인 대 개인의 의사소통에 그다지 큰 문제가 되지 않는다.
실제로 다른 부서의 사람들 사이에서 의사소통이 어려울 수 있지만 사람들이 서로 상호작용하면서 정확한 의미를 추론하는 것은 쉽다.

그러나 이러한 다양한 비즈니스 도메인 모델을 소프트웨어로 표현하는 것은 더 어렵다.
소스코드는 모호성에 잘 대처하지 못한다.

이 문제에 대한 전통적인 솔루션은 모든 종류의 문제에 사용할 수 있는 단일 모델을 설계하는 것이다.
이러한 모델은 모든 것에 적합해야 하지만 결국에는 아무 소용이 없다.
무엇을 하든지 항상 복잡성에 직면하게 된다.
즉 관련 없는 세부사항을 필터링하는 복잡성, 필요한 것을 찾는 복잡성, 데이터를 일관된 상태로 유지해야 하는 복잡성에 직면하게 된다.

또 다른 솔루션은 문맥상 정의에 문제가 있는 용어 앞에 접두사를 추가하는 것이다.
그러면 모델을 코드로 만들 수 있다.
그러나 이 접근 방식에는 두 가지 주요 단점이 있다.
1. 인지부하를 유발한다.
각 모델을 언제 사용해야 할까? 충돌하는 모델을 계속해서 구현할수록 실수하기 쉽다.
2. 모델의 구현이 유비쿼터스 언어와 일치하지 않는다.
아무도 대화에서 접두사를 사용하지 않을 것이다.

# 바운디드 컨텍스트
도메인 주도 설계에서 솔루션은 유비쿼터스 언어를 여러 개의 작은 언어로 나눈 다음 각 언어를 적용할 수 있는 명시적인 바운디드 컨텍스트에 할당하면 된다.

## 모델 경계
모델은 복잡한 시스템을 이해하는 데 도움을 주기 위해 구조화한 것이다.
우리가 해결하려는 문제는 모델 본연의 목적이다.
모델은 경계 없이 존재할 수 없다.
경계가 없다면 현실 세계의 복제본처럼 확장될 것이다.
따라서 모델의 경계(바운디드 컨텍스트)를 정의하는 것은 모델링 프로세스의 본질적인 부분이다.

하나의 바운디드 컨텍스트의 유비쿼터스 언어는 다른 바운디드 컨텍스트의 범위에는 완전히 관련이 없다.
바운디드 컨텍스트는 유비쿼터스 언어와 해당 언어가 나타내는 모델의 적용 가능성을 규정한다.
서로 다른 문제 도메인에 따라 고유한 모델을 정의할 수 있다.
바운디드 컨텍스트는 유비쿼터스 언어의 일관성이 유지되는 경계다.
유비쿼터스 언어의 용어, 원칙, 비즈니스 규칙은 해당 바운디드 컨텍스트 내에서만 일관성이 있다.

## 정제된 유비쿼터스 언어
바운디드 컨텍스트를 통해 유비쿼터스 언어의 정의를 완성할 수 있다.
유비쿼터스 언어는 만능이 아니다.

대신 유비쿼터스 언어는 바운디드 컨텍스트 경계 안에서만 보편적으로 적용된다.
유비쿼터스 언어는 바운디드 컨텍스트에 포함된 모델을 설명하는 데만 집중한다.
모델은 우리가 해결해야 하는 문제 없이는 존재할 수 없기 때문에 유비쿼터스 언어는 명시적으로 적용 가능한 컨텍스트 없이 정의하거나 사용할 수 없다.

## 바운디드 컨텍스트의 범위
서로 다른 도메인 전문가들은 동일한 비즈니스 엔티티에 대해 상충되는 멘탈 모델을 가지고 있다.
비즈니스 도메인을 모델링하기 위해 우리는 모델을 분할하고 각 세분화된 모델에 적용 가능한 컨텍스트를 엄격하게 정의해야 했다.

유비쿼터스 언어의 일관성은 해당 언어의 가장 넓은 경계를 식별하는 데 도움이 될 뿐이다.
일관성이 없는 모델과 용어가 있기 때문에 더 이상 커질 수 없다.
그러나 모델을 더 작은 바운디드 컨텍스트로 분해할 수 있다.

유비쿼터스 언어의 범위(바운디드 컨텍스트)를 정의하는 것은 전략적인 설계 의사결정이다.
경계는 비즈니스 도메인의 고유한 컨텍스트에 따라 넓힐 수 있고, 비즈니스 도메인을 더 작은 문제 도메인으로 세분화하여 좁힐 수도 있다.

바운디드 컨텍스트 크기 자체는 의사결정 요소가 아니다.
모델의 크기에 정답은 없다.
다만 모델 자체로 유용해야 한다.
유비쿼터스 언어의 경계가 넓을수록 일관성을 유지하기가 더 어렵다.
큰 유비쿼터스 언어를 더 작고 관리하기 쉬운 문제 도메인으로 나누는 것은 도움이 될 수 있지만, 바운디드 컨텍스트를 작게 만들기 위해 노력하는 것은 역효과를 낼 수 있다.
더 작게 만들수록 설계를 통합하는 오버헤드가 커진다.

따라서 바운디드 컨텍스트의 크기에 대한 결정은 문제 도메인이 무엇이냐에 따라 달라진다.
때로는 넓은 경계를 사용하는 것이 더 명확하고 어떤 경우에는 분해하는 것이 더 합리적이다.

## 바운디드 컨텍스트 대 하위 도메인

### 하위 도메인
비즈니스 전략을 이해하려면 비즈니스 도메인을 분석해야 한다.
분석 단계에는 다양한 하위 도메인(핵심, 지원, 일반)을 식별하는 작업이 포함된다.

하위 도메인은 상호 관련된 유스케이스 집합과 유사하다.
유스케이스에서는 비즈니스 도메인과 시스템 요구사항에 따라 정의된다.
소프트웨어 엔지니어로서 우리는 요구사항을 정의하지 않는다.
그것은 비즈니스가 담당한다.
대신 소프트웨어 엔지니어는 하위 도메인을 식별하기 위해 도메인을 분석한다.

### 바운디드 컨텍스트
바운디드 컨텍스트는 소프트웨어 엔지니어에 의해 설계된다.
모델의 경계를 선택하는 것은 전략적 설계의 의사결정이다.
비즈니스 도메인을 더 작고 관리 가능한 문제 도메인으로 어떻게 나눌지 정한다.

### 하위 도메인과 바운디드 컨텍스트 사이의 상호작용
비현실적이지만 이론적으로는 단일 모델이 전체 비즈니스 도메인에 적용될 수 있다.
모델이 충돌하면 도메인 전문가의 멘탈 모델을 따라 시스템 바운디드 컨텍스트로 분해할 수 있다.
모델이 여전히 크고 유지보수하기 어려운 경우 더 작은 바운디드 컨텍스트로 분해할 수 있다.

하위 도메인은 발견하고 바운디드 컨텍스트는 설계한다는 점이다.
하위 도메인은 비즈니스 전략에 의해 정의된다.
그러나 소프트웨어 엔지니어는 특정 프로젝트의 컨텍스트와 제약 조건을 해결하기 위해 소프트웨어 솔루션과 바운디드 컨텍스트를 설계할 수 있다.

모델은 특정 문제를 해결하기 위한 것이다.
어떤 경우는 같은 개념의 여러 모델을 동시에 사용하여 다른 문제를 해결하는 것이 도움이 될 수 있다.
문제에 따라 동일한 하위 도메인에 대한 다른 모델을 사용하여 해결하는 것이 합리적일 수 있다.
바운디드 컨텍스트와 하위 도메인의 설계를 일대일 관계로 제한하면 이런 식의 유연성이 떨어지고 바운디드 컨텍스트 안에서 하나의 하위 도메인 모델만을 사용하게 된다.

## 경계
바운디드 컨텍스트 패턴은 물리적 경계와 소유권 경계를 규정하기 위한 도메인 주도 설계 도구다.

### 물리적 경계
바운디드 컨텍스트는 모델 경계뿐만 아니라 이를 구현하는 시스템의 물리적 경계 역할도 한다.
각 바운디드 컨텍스트는 개별 서비스/프로젝트로 구현돼야 한다.
즉 구현, 진화, 버전 관리를 각각의 다른 바운디드 컨텍스트와 독립적으로 해야 한다.

바운디드 컨텍스트 간의 명확한 물리적 경계를 통해 각 바운디드 컨텍스트를 요구사항에 가장 적합한 기술 스택으로 구현할 수 있다.

바운디드 컨텍스트는 여러 하위 도메인을 포함할 수 있다.
이러한 경우 바운디드 컨텍스트는 물리적 경계고 하위 도메인은 논리적 경계다.
논리적 경계는 프로그래밍 언어의 종류에 따라 네임스페이스나 모듈, 패키지 같은 다른 이름을 갖는다.

### 소유권 경계
바운디드 컨텍스트는 한 팀에서만 구현, 발전 유지 관리해야 한다.
이러한 분리는 팀이 서로 모델에 대해 만들 수 있는 암묵적인 가정을 제거한다.

대신 팀은 서로 다른 바운디드 컨텍스트로 분리된 모델과 시스템을 명시적으로 연동하기 위한 통신 프로토콜을 정의해야 한다.

팀과 바운디드 먼텍스트 간의 관게는 단방향이다.
바운디드 컨텍스트는 한 팀만 소유해야 한다.
그러나 단일 팀이 여러 개의 바운디드 컨텍스트를 소유할 수도 있다.

# 결론
도메인 전문가의 멘탈 모델에 내재된 충돌을 발견할 때마다 유비쿼터스 언어를 여러 개 바운디드 컨텍스트로 분해해야 한다.
유비쿼터스 언어는 바운디드 컨텍스트의 범위 내에서 일관성이 있어야 한다.
그러나 서로 다른 바운디드 컨텍스트에서는 동일한 용어라도 다른 의미를 가질 수 있다.

하위 도메인이 발견되면 바운디드 컨텍스트도 설계한다.
도메인을 바운디드 컨텍스트로 나누는 것은 전략적 설계의 의사결정이다.

바운디드 컨텍스트와 유비쿼터스 언어는 한 팀에서 만들고 유지보수할 수 있다.
두 팀이 동일한 바운디드 컨텍스트에서 작업을 공유할 수 없다.
그러나 한 팀이 여러 개의 바운디드 컨텍스트에서 일하는 것은 가능하다.

바운디드 컨텍스트 시스템을 서비스, 하위 시스템 등의 물리적 구성요소로 분해한다.
각 바운디드 컨텍스트의 수명주기는 서로 독립적이다.
각 바운디드 컨텍스트는 시스템의 나머지 부분과 독립적으로 발전할 수 있다.
그러나 바운디드 컨텍스트는 시스템을 구성하기 위해 함께 작동해야 한다.
일부 변경은 의도치 않게 다른 바운디드 컨텍스트에 영향을 줄 수 있다.
