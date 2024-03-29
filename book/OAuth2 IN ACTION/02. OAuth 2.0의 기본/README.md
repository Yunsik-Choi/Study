# OAuth 2.0 프로토콜의 개요

OAuth 트랜잭션에는 토큰을 발급하고 발급된 토큰을 사용하는 2개의 중요한 단계가 있다.

표준 OAuth 트랜잭션은 다음과 같은 일련의 이벤트로 구성된다.
- 리소스 소유자는 클라이언트에게 자신을 대신해 작업을 수행하라고 말한다.
- 클라이언트는 인가 서버의 리소스 소유자에게 인가를 요청한다.
- 리소스 소유자는 클라이언트를 인가한다.
- 클라이언트는 인가 서버로부터 토큰을 전달받는다.
- 클라이언트는 보호된 리소스에 접근하기 위해 토큰을 사용한다.

# OAuth 2.0 인가 그랜트
인가 코드 그랜트는 리소스 소유자가 클라이언트에게 접근 권한을 위임했다는 것을 나타내기 위해 임시 자격 증명인 인가 코드를 사용한다.

## 인가 코드 그랜트의 단계
- 리소스 소유자는 클라이언트에게 자신을 대신해 보호된 리소스에 접근해 작업을 수행하라고 말한다.
- 클라이언트는 리소스 소유자의 권한 일부를 자신에게 위임해달라고 요청하기 위해 리소스 소유자를 HTTP 리다이렉트를 이용해 인가 서버의 인가 엔드 포인트로 리다이렉트시킨다.
- 인가 서버는 전달된 파라미터를 분석해 그에 맞는 작업을 수행한다.
- 그런 다음, 인가 서버는 일반적으로 사용자의 인증을 요구한다. 사용자의 인증 정보는 사용자(사용자의 웹 브라우저)와 인가 서버 간에 직접 전달된다.
- 사용자는 클라이언트 애플리케이션을 인가한다. 이 단계에서 리소스 소유자는 클라이언트에게 자신의 권한 일부를 위임할 것인지를 선택한다.
- 인가 서버가 사용자를 클라이언트 애플리케이션으로 리다이렉트시키면서 사용자에게 인가 코드를 전달한다.
- 사용자가 클라이언트에게 인가가 코드를 전달한다. 
- 클라이언트가 인가 코드를 획득하면 그것을 다시 인가 서버의 토큰 엔드 포인트로 전달한다.
- 인가 서버는 전달된 클라이언트의 토큰 요청이 유효하면 토큰을 발급한다.
- 클라이언트는 토큰 응답을 분석해 액세스 토큰 값을 추출하고 그것을 이용해 보호된 리소스에 접근한다.
- 보호된 리소스는 헤더에서 전달된 토큰을 추출하고 그것이 유효한 것인지 확인하고 그것에 연관된 정보 즉 누가 그것을 인가했고 무엇을 위해 인가됐는지 확인한다.
- 그 다음에는 요청에 대한 적절한 응답을 전달한다.

# OAuth의 구성원: 클라이언트, 인가 서버, 리소스 소유자, 그리고 보호된 리소스
OAuth 시스템은 클라이언트, 리소스 소유자, 인가 서버, 보호된 리소스 4개의 구성원으로 이루어진다.
각 구성원은 OAuth 프로토콜의 서로 다른 부분을 담당하며 모두 함께 동작해 OAuth 프로토콜이 동작한다.

**OAuth 클라이언트**는 리소스 소유자를 대신해 보호된 리소스에 접근하고자 하는 소프트웨어다.
인가 서버로부터 액세스 토큰을 획득하고 획득한 액세스 토큰을 이용해 보호된 리소스에 접근하는 역할을 담당한다.

**보호된 리소스**는 HTTP 서버를 통해 접근하며 접근을 위해서는 OAuth 토큰이 필요하다.
보호된 리소스는 전달된 토큰의 유효성을 검사하고 요청을 어떻게 처리할지를 결정한다.

**리소스 소유자**는 클라이언트에게 접근 권한을 위임할 권한을 갖고 있다.
OAuth 시스템의 다른 구성원과 달리 리소스 소유자는 소프트웨어가 아니다.
대부분의 경우, 리소스 소유자는 클라이언트 소프트웨어를 이용하는 사람이다.
리소스 소유자는 웹 브라우저를 이용해 인가 서버와 상호 작용한다.

**OAuth 인가 서버**는 HTTP 서버로서 OAuth 시스템의 핵심적인 구성원이다.
인가 서버는 리소스 소유자와 클라이언트를 인증하고 리소스 소유자가 클라이언트에게 권한을 위임할 수 있는 메커니즘을 제공한다.
그리고 클라이언트에게 토큰을 발급한다.
인가 서버 중에는 인가 결정을 기억하거나 토큰에 대해 자기 검증을 제공하는 경우도 있다.

# OAuth의 구성 요소: 토큰, 범위 그리고 인가 그랜트
### 액세스 토큰
OAuth 액세스 토큰은 간단히 토큰이라고도 하며, 클라이언트에게 권한이 위임됐다는 것을 나타내기 위해 인가 서버가 클라이언트에게 발급한다.
OAuth에서는 토큰의 포맷이나 내용을 정의하지 않지만 항상 클라이언트가 요청한 접근 권한, 리소스 소유자의 클라이언트 인가 그리고 인가 과정에서 클라이언트에게 부여된 권한을 나타낸다.
인가 서버는 토큰을 이해하고 어떤 내용을 의미하는지 알 수 있지만 클라이언트가 볼 때 OAuth 토큰의 형식은 명확하지 않아서 클라이언트는 토큰의 내용이나 의미를 전혀 알지 못한다.
클라이언트의 역할은 인가 서버에 토큰을 요청하고 인가 서버가 발급한 토큰을 이용해 보호된 리소스에 접근하는 것이다.

### 범위
OAuth 범위는 보호된 리소스에 대한 접근 권한을 나타낸다.
OAuth 프로토콜에서 범위는 공백으로 구분된 범위 문자열의 조합으로 표현된다.
따라서 범위를 나타내는 문자열값에는 공백이 포함되면 안 된다.
하지만 범위값의 포맷과 구조가 OAuth 스펙에 정의된 것은 아니다.

범위는 클라이언트에게 부여되는 접근 권한을 제한하기 위한 중요한 메커니즘이다.
범위는 보호된 리소스가 제공하는 API를 기반으로 정의된다.
클라이언트는 특정 범위의 권한을 요청하고 인가 서버는 요청된 권한 범위를 리소스 소유자가 허용할 것인지 아니면 거부할 것인지 선택할 수 있게 해준다.

### 리프레시 토큰
OAuth 리프레시 토큰은 인가 서버가 클라이언트에게 발급하고 클라이언트는 발급 받은 토큰의 내용이 무엇인지 알지 못하거나 개의치 않는다는 점에서 액세스 토큰과 유사하다고 할 수 있다.
차이점이라고 한다면 리프레시 토큰은 보호된 리소스에 전달되지 않는다는 것이다.
대신에 클라이언트는 리소스 소유자와는 관련없이 리프레시 토큰을 이용해 새로운 액세스 토큰을 요청한다.

OAuth에서는 특정 시점에 액세스 토큰을 더 이상 사용할 수 없게 된다.
OAuth 2.0에서는 토큰의 유효 기간이 자동으로 설정되기 때문에 토큰의 유효 기간이 만료되거나 그때 사용자가 없더라도 리소스에 접근할 수 있는 방법이 필요하다.
그래서 만료 시간 없이 지속적으로 사용하는 토큰 대신 리프레시 토큰이 사용된다.
하지만 리프레시 토큰은 리소스에 접근하기 위한 용도로 사용되는 것이 아니라 리소스에 접근하는 데 사용되는 새로운 액세스 토큰을 요청하는 용도로만 사용된다.
그렇게 함으로써 리프레시 토큰과 액세스 토큰의 유출을 개별적으로 제한할 수 있으며 상호 보완적으로 만들 수 있다.

리프레시 토큰은 또한 클라이언트의 접근 권한을 축소시킬 수도 있다.
만약 클라이언트가 A, B, C 라는 권한 범위를 할당받았는데 A만 있으면 필요한 작업을 수행할 수 있다고 가정하면 리프레시 토큰을 이용해 A 권한 범위만 가진 액세스 토큰을 요청할 수 있다.

만약 리프레시 토큰 자체가 제대로 동작하지 않으면 클라이언트는 리소스 소유자를 매번 귀찮게 할 수밖에 없다.
즉 OAuth 클라이언트에 어떤 장애가 발생하면 OAuth 인가 절차를 다시 수행해야 한다.

### 인가 그랜트
인가 그랜트는 OAuth 프로토콜을 이용해 OAuth 클라이언트에게 보호된 리소스에 대한 접근 권한을 위임해주기 위한 방법이며, 그것이 성공적으로 이뤄지면 클라이언트는 최종적으로 액세스 토큰을 얻게 된다.
인가 그랜트는 OAuth의 전체 프로세스를 의미한다.
즉 클라이언트가 사용자를 인가 엔드포인트로 이동시키고, 인가 코드를 전달받고 마지막으로 인가 코드를 토큰과 교환하는 과정 전체가 인가 그랜트인 것이다.
다시 말하면 인가 그랜트는 토큰을 획득하기 위한 방법이다.

![img.png](img.png)

# OAuth 구성원과 구성 요소 간의 상호 작용 : 백 채널, 프런트 채널, 그리고 엔드 포인트
OAuth는 HTTP 기반의 프로토콜이지만 대부분의 HTTP 기반 프로토콜과 달리, OAuth 통신은 단순한 HTTP 요청과 응답으로 이뤄지지 않는다.

## 백 채널 통신
OAuth 프로세스의 많은 부분에서 일반적인 HTTP 요청과 응답으로 구성 요소 간의 통신이 이뤄진다.
그런 통신은 일반적으로 리소스 소유자와 사용자 에이전트 범위 밖에서 이뤄지기 때문에 백 채널 통신이라고 한다.
백 채널은 브라우저를 제외한 구성 요소 간의 직접적인 HTTP 연결을 이용한다.

백 채널 통신은 모두 일반적인 HTTP 메커니즘으로 이뤄진다.
즉 헤더, 질의 파라미터, 질의 메서드, 그리고 질의 요청 내용에 중요한 트랜잭션 정보가 포함된다.
일반적인 경우보다 좀 더 많은 HTTP 스택이 이용되는 특징이 있으며, 대부분의 간단한 웹 API에서는 클라이언트 개발자가 그 응답에 좀 더 주의를 기울일 수 있다.

인가 서버는 클라이언트가 액세스 토큰과 리프레시 토큰을 요청하기 위해 사용하는 토큰 엔드 포인트를 제공한다.
클라이언트는 토큰 엔드 포인트를 직접 호출하며, 인가 서버가 분석하고 처리할 수 있는 인코딩된 폼의 파라미터 형태로 요청 내용을 전달한다.
그러면 인가 서버는 JSON 객체 형식의 토큰을 응답으로 전달한다.

클라이언트가 보호된 리소스에 접근할 때도 백 채널 통신으로 HTTP를 직접 호출한다.
이때 전달되는 내용은 전적으로 보호된 리소스에 따라 결정된다.

백 채널 통신에서 클라이언트는 OAuth 토큰을 전달하고 보호된 리소스는 자신에게 전달된 토큰과 그것이 갖고 있는 권한을 이해할 수 있어야 한다.

## 프런트 채널 통신
일반적인 HTTP 통신에서 HTTP 클라이언트는 헤더와 질의 파라미터, 질의 내용 그리고 서버에 직접 전달할 정보를 포함해 요청을 보낸다.
서버는 전달된 정보의 내용을 보고 해당 요청에 대해 어떻게 응답할 것인지 결정한다.
서버가 전달하는 HTTP 응답에도 헤더와 응답 내용 그리고 그 외 정보가 포함된다.
하지만 OAuth 에서는 두 구성 요소가 직접 HTTP 요청을 보내고 응답을 받지 않는 경우가 있다.
즉 클라이언트가 인가 서버의 인가 엔드포인트와 상호 작용하는 것이 바로 그런 경우다.
프런트 채널 통신은 중간의 웹 브라우저를 통해 두 시스템간이 간접적으로 HTTP 통신을 하는 방법이다.

이 기술은 웹 브라우저의 양쪽 세션을 분리하며, 그로 인해 서로 다른 보안 도메인 간의 작업을 가능하게 해준다.
그렇다면 두 소프트웨어가 서로 간의 대화 없이 어떻게 통신할 수 있을까?
프런트 채널 통신은 웹 브라우저가 방문해야 하는 URL을 파라미터를 전달함으로써 이뤄진다.
그런 URL을 전달받은 쪽에서는 해당 URL을 파싱하고 함께 전달된 파라미터와 동일한 방법으로 웹 브라우저를 해당 URL로 리다이렉트 한다.

그렇게 중간의 웹 브라우저를 통해 두 통신 주체는 서로 간접적으로 통신한다.
각 프런트 채널 요청과 응답은 실제로 HTTP 요청과 응답 트랜잭션 쌍이라고 할 수 있다.

예를 들면 인가 코드 그랜트에서 클라이언트는 사용자를 인가 엔드 포인트로 리다이렉트해야 하지만, 요청의 특정 부분은 인가 서버와도 통신해야 한다.
이를 위해 클라이언트는 웹 브라우저에게 HTTP 리다이렉트를 전달해야 한다.

인가 서버는 다른 HTTP 요청과 마찬가지로 인입되는 URL을 파싱해 클라이언트가 파라미터로 전달한 정보를 추출한다.
그때 인가 서버는 리소스 소유자와 상호 작용해 웹 브라우저와 일련의 HTTP 트랜잭션을 통해 사용자를 인증하고 클라이언트에 대한 권한 위임을 요청한다.
그리고 인가 서버는 클라이언트에게 인가 코드를 전달하기 위해 이 경우에도 웹 브라우저에게 HTTP 리다이렉트를 전달한다.
하지만 이번에는 리다이렉트되는 목적지 주소가 클라이언트의 redirect_url이 된다.
또한 전달하는 리다이렉트 URL에 질의 파라미터를 포함시킨다.

웹 브라우저가 인가 서버가 전달한 주소로 리다이렉트하면 클라이언트 애플리케이션에는 HTTP 요청이 전달된다.
그러면 클라이언트는 전달된URL 파라미터를 파싱한다.
이런 식으로 클라이언트와 인가 서버는 서로 직접 대화하지 않고 중개인을 통해 서로 메시지를 전달한다.

프런트 채널을 통해 전달되는 모든 정보는 웹 브라우저가 접근 가능하다.
따라서 그 내용을 읽거나 질의가 전송되기 전에 질의 내용이 조작될 잠재적인 가능성이 존재한다.
이 점 때문에 OAuth 프로토콜은 프런트 채널을 통해 전달되는 정보의 종류를 제한하고 프런트 채널에서 사용된 어떤 정보도 단독으로 권한 위임 작업을 위해 사용될 수 없도록 한다.
인가 코드는 웹 브라우저가 직접 사용할 수 없지만, 백 채널에서 클라이언트의 자격 증명과 함께 제시돼야 한다.
OpenID 커넥트와 같은 몇몇 프로토콜은 클라이언트나 인가 서버가 프런트 채널 메시지를 서명함으로써 보다 향상된 보안성을 제공한다.

# 요약
OAuth는 많은 유동적인 부분을 갖고 있는 프로토콜이지만, 권한 위임하기 위한 완전한 방법을 간단히 제공할 수 있도록 만들어졌다.
- OAuth는 토콘을 획득하고 그것을 사용하는 것에 관한 것이다.
- OAuth 시스템의 구성 요소는 프로세스의 다른 부분을 중요하게 여긴다.
- OAuth 구성 요소는 서로 간의 통신을 위해 직접적(백 채널), 간접적(프런트 채널)인 HTTP 통신을 수행한다.