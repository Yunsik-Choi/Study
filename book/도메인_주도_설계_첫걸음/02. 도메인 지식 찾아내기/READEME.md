# 비즈니스 문제
우리가 개발하는 소프트웨어 시스템은 비즈니스 문제를 해결하는 솔루션이다. 

비즈니스 문제는 비즈니스 도메인과 하위 도메인의 모든 수준에서 발생할 수 있다. 기업의 목표는 고객의 문제를 해결하는 솔루션을 제공하는 것이다. 

하위 도메인은 세분화된 문제 도메인으로 특정 비즈니스 기능에 대한 솔루션을 제공하는 것이 목적이다.

# 도메인 지식 찾아내기
효과적인 소프트웨어 솔루션을 설계하려면 적어도 기본적인 비즈니스 도메인 지식이 있어야 한다. 이런 지식은 도메인 전문가의 몫이다. 우리는 도메인 전문가가 될 필요도 없고 될 수도 없다. 도메인 전문가를 이해하고 그들이 쓰는 동일한 비즈니스 용어를 사용하는 것이 중요하다.

효과적인 소프트웨어는 도메인 전문가가 문제를 생각하는 방식 즉 멘탈 모델을 모방해야 한다.
비즈니스 문제와 요구사항 이면에 있는 이유에 대한 이해가 없다면 솔루션은 비즈니스 요구사항을 소스코드로 번역한 것에 불과하다.

소프트웨어 프로젝트 성공은 도메인 전문가와 소프트웨어 엔지니어 간의 효과적인 지식 공유에 달렸다.
문제를 해결하려면 문제를 이해해야 한다.
결국 도메인 전문가와 소프트웨어 엔지니어 간의 효과적인 지식 공유에는 효과적인 커뮤니케이션이 필요하다.

# 커뮤니케이션
거의 모든 소프트웨어 프로젝트에는 다양한 역할의 이해관계자의 협업이 필요하다고 할 수 있다.
결과는 이런 모든 참여자가 얼마나 잘 협력할 수 있느냐에 달려있다.
프로젝트와 연관된 모든 것에 대한 합의와 일치하는 프로젝트의 성공에 필수다.

전형적인 소프트웨어 개발 생에주기에서 도메인 지식은 분석 모델로 알려진 엔지니어 친화적인 형태로 변환된다.
분석 모델은 도메인 지식 이면에 존재하는 비즈니스 도메인에 기반하기보다는 시스템 요구사항을 설명한 것에 지나지 않는다.
이런 식의 중재는 지식을 공유하는 데 해를 끼친다.
모든 변환은 정보를 잃게 만든다.
이런 경우 비즈니스 문제 해결에 중요한 도메인 지식은 소프트웨어 엔지니어에게 전달되는 과정에서 손실된다.
이런 현상은 일반적인 소프트웨어 프로젝트에서 변환할 때뿐만 아니라 분석 모델을 소프트웨어 설계 모델로 변환할 때도 발생한다.
문서화된 커뮤니케이션은 최신 정보를 담아내지 못한다.
결국 소스코드가 나중에 프로젝트를 유지관리할 소프트웨어 엔지니어에게까지도 비즈니스 도메인 지식을 전달하는데 사용된다.

이 같은 문제를 해결하기 위해 도메인 주도 설계는 도메인 전문가가 소프트웨어 엔지니어에게 지식을 전달하기 위한 더 나은 방법을 제언하는데 유비쿼터스 언어가 바로 그것이다.

# 유비쿼터스 언어란 무엇인가?
전통적인 소프트웨어 개발 생애주기
- 도메인 지식이 분석 모델로
- 분석 모델이 요구사항으로
- 요구사항은 시스템 설계로
- 시스템 설계는 소스코드로

반면에 도메인 주도 설계에서는 이같이 도메인 지식을 계속해서 변환하는 대신 비즈니스 도메인을 설명하기 위한 단일화된 언어 체계를 세우고자 하는데 이것이 유비쿼터스 언어다.

프로젝트와 관련된 모든 이해관계자는 비즈니스 도메인을 설명할 때 유비쿼터스 언어를 사용해야 한다.
가장 중요한 것은 도메인 전문가가 유비쿼터스 언어를 사용해 비즈니스 도메인을 추론하는 데 편안함을 느껴야 한다는 점이다.

모든 프로젝트 참가자의 공통된 이해는 오직 유비쿼터스 언어와 그 용어의 지속적인 사용을 통해서만 함양될 수 있다.

# 비즈니스 언어
유비쿼터스 언어는 비즈니스 언어다.
그렇기 때문에 기술 용어는 빼고 비즈니스 도메인에 관련된 용어로만 구성해야 한다.
유비쿼터스 언어는 도메인 전문가의 이해와 비즈니스 도메인에 대한 멘탈 모델을 쉽게 이해할 수 있는 관점으로 표현하는 것을 목표로 한다.

## 시나리오
모든 문장을 비즈니스 언어로 작성하고 문장은 비즈니스 도메인을 바라보는 도메인 전문가의 시각을 반영한다.

## 일관성
유비쿼터스 언어는 반드시 정확하고 일관성이 있어야 한다.
가정할 필요가 없어야 하고 비즈니스 도메인의 로직을 명료하게 표현해야 한다.

모호성이 커뮤니케이션을 방해하기 때문에 유비쿼터스 언어의 용어는 오직 하나의 의미를 가져야 한다.

### 모호한 용어
어떤 비즈니스 도메인에서는 정책이라는 용어가 규제 규칙 또는 보험 계약과 같은 여러 의미를 가진다.
정확한 의미는 맥락에 따라 사람 간의 상호작용을 통해서만 알 수 있다.
그러나 소프트웨어는 모호성에 잘 대처하지 못하며 정책이라는 entity를 코드로 모델링하기가 어려울 수 있다.

유비쿼터스 언어는 용어마다 단일 의미를 갖게 하기 때문에 '정책'의 경우 '규제 규칙'과 '보험 계약'의 두 용어를 사용하여 명확한 모델을 만들어야 한다.

### 동의어
유비쿼터스 언어에서 두 용어는 서로 바꿔 사용할 수 없다.
예를 들어 '사용자'라는 용어는 수많은 시스템에서 사용한다.
그러나 도메인 전문가의 언어로 설명하면 '사용자'와 '방문자', '관리자', '계정'등의 다른 용어가 혼용된다.

동의어는 처음에는 해로워 보이지 않는다.
그러나 대부분의 경우 다른 개념을 가진다.

특정 컨텍스트 안에서 각각의 용어를 사용하는 것이 바람직하다.
용어의 차이점을 이해해야 간단하고 명확한 모델을 구축하고 비즈니스 도메인 객체의 구현이 가능하다.

# 비즈니스 도메인 모델
모델링 관점에서 유비쿼터스 언어

## 모델이란 무엇인가?
모델은 실세계의 복제가 아니라 우리가 실제 시스템을 이해하는 데 도움을 주는 인간의 창조물이다.

## 효과적인 모델링
모든 모델에는 목적이 있고 효과적인 모델은 그 목적을 달성하는 데 필요한 세부사항만 포함한다.
유용한 모델은 실세계의 복사본이 아니라 문제를 해결하려는 의도가 있으며, 그 목적에 필요한 정보만 제공해야 한다.

모델은 본질적으로 추상화의 결과다.
추상화의 개념을 불필요한 상세 정보를 생략하여 복잡한 문제를 다룰 수 있게 하고 당면한 문제를 푸는 데 필요한 정보만 남게 한다.
비효과적인 추상화는 필요한 정보를 제거하거나 필요 없는 정보를 포함해 잡음을 유발한다.

추상화의 목적은 모호함이 아니라 절대적으로 정확할 수 있는 새로운 의미론적 수준을 만드는 것이다.

## 비즈니스 도메인 모델링
유비쿼터스 언어를 발전시키는 것은 사실상 비즈니스 도메인 모델을 구축하는 셈이다.
모델은 비즈니스가 기능을 어떻게 구현하느냐에 대한 도메인 전문가의 사고 프로세스인 멘탈 모델을 포착해야 한다.
모델은 관련된 비즈니스 엔티티와 그것의 행동, 인과 관계, 불변성 등을 반영해야 한다.

유비쿼터스 언어는 도메인의 모든 가능한 상세 정보를 포함하는 게 아니다.
모델은 소프트웨어가 해결하고자 하는 특정 문제를 해결하는 데 필요한 만큼의 비즈니스 도메인 관점을 포함하면 된다.

엔지니어링 팀과 도메인 전문가의 효과적인 커뮤니케이션은 필수다.
비즈니스 도메인이 복잡할수록 커뮤니케이션의 중요성은 커진다.
비즈니스 도메인이 복잡할수록 비즈니스 로직을 모델링하고 코드로 구현하기가 더 어렵다.
비즈니스 도메인의 이해를 확인할 신뢰성 있는 유일한 방법은 도메인 전문가가 이해할 수 있는 비즈니스 언어로 대화하는 것이다.

## 지속적인 노력
유비쿼터스 언어를 정형화 하려면 언어의 소유자인 도메인 전문가와의 상호작용이 필요하다.
오직 실제 도메인 전문가와의 상호작용만이 비즈니스 도메인에 대한 부정확함이나 잘못된 가정, 또는 전체적인 이해 오류를 발견할 수 있다.

모든 이해관계자는 모든 프로젝트와 관련된 커뮤니케이션에 유비쿼터스 언어를 지속적으로 사용해서 지식 공유를 확산하고 비즈니스 도메인에 대한 공유된 이해를 강화해야 한다.
요구사항, 테스트, 문서화, 소스코드 자체 등 프로젝트 전반에 걸쳐 이 언어를 지속해서 사용해야 한다.

가장 중요한 점은 유비쿼터스 언어를 발전시키는 것은 진행형이라는 것이다.
지속해서 검증하고 발전시켜야 한다.
모두가 이 언어를 일상적으로 사용하면 비즈니스 도메인에 대한 통찰을 얻을 것이다.
이런 획기적인 발전이 일어나면 새롭게 얻은 도메인 지식에 맞춰 유비쿼터스 언어 또한 발전해야 한다.

## 도구
유비쿼터스 언어를 수집하고 관리하는 과정을 돕는 도구와 기술이 있다.
예를 들어 위키는 유비쿼터스 언어를 수집하고 관리하는 용어집으로 사용될 수 있다.
이런 용어집은 비즈니스 도메인의 용어에 대한 정보를 얻을 수 있는 거점 역할을 하므로 새로운 팀원이 쉽게 적응하게 해준다.

또한 다 함께 용어집을 유지보수하는 것이 중요하다.
유비쿼터스 언어가 변경되면 모든 팀원이 수정할 수 있게 독려해야 한다.
이는 팀 리더나 아키텍트만 용어집의 유지보수를 담당하는 중앙 집중식 방식과 정반대다.

프로젝트 용어집을 유지보수하는 것의 장점이 명백함에 불구하고도 본질적인 한계가 있다.
용어집은 엔티티의 이름, 과정, 역할 등의 명사에만 효과적이다.
명사가 중요하긴 하지만 행동을 포착하는 것이 중요하다.
행동은 단순히 명사와 관련된 동사의 목록이 아니라 규칙, 가정, 그리고 불변성을 가진 실제 비즈니스 로직이다.
이런 개념은 용어집으로 문서화하기 훨씬 어렵다.
그러므로 용어집은 유스케이스 또는 거킨 테스트 처럼 행동을 포착하는데 적합한 다른 도구와 함께 사용하는 것이 가장 좋다.

거킨 언어로 작성된 자동화 테스트는 유비쿼터스 언어를 포착하기에 좋은 언어일 뿐 아니라 도메인 전문가와 소프트웨어 엔지니어의 간극을 메우는 보도 도구로서의 역할을 할 수 있다.
도메인 전문가는 테스트를 읽고 시스템의 기대 행동을 검증할 수 있다.

거킨 기반의 테스트 스위트를 관리하는 것은 때때로 어려운 일이다.
특히 프로젝트 초기 단계에 더욱 그렇다.
그러나 복잡한 비즈니스 도메인인 경우 확실히 가치가 있다.

이런 도구가 유용하긴 하지만 일상적인 상호작용에서 실제로 유비쿼터스 언어를 사용하는 것보다는 못하다.
도구를 사용하여 유비쿼터스 언어를 관리하더라도 문서화가 실제로 언어를 사용하는 것을 대체할 것으로 기대하지는 말자. 
프로세스나 도구보다 개인과 상호작용이 우선이다.

## 도전과제
유비쿼터스 언어를 발전시키는 것이 간단한 과정 같지만 실제로 그렇지 않다.
도메인 지식을 수집하는 신뢰할 만한 유일한 방법은 도메인 전문가와 대화를 하는 것이다.
대부분의 경우 가장 중요한 지식은 암묵지다.
이것은 문서화되거나 코드로 작성되지 않고 도메인 전문가의 정신에만 존재한다.
여기에 접근하는 유일한 방법은 질문하는 것이다.

이 기법에 경험이 쌓이면 이 과정이 단순히 존재하는 지식을 발견하는 것뿐만 아니라 도메인 전문가와 협력해서 모델을 함께 만들어나가는 것이 자주 포함된다는 사실을 알게 된다.
비즈니스 도메인의 특성에 대해 질문하면 종종 숨어있던 충돌과 공백을 찾아내 명확하게 할 수 있다.
이는 특히 핵심 하위 도메인의 경우에 일반적이다.
이런 과정은 도메인 전문가가 자신의 영역을 더 잘 이해하도록 돕는 상호보완적인 배움의 과정이기도 하다.

프로젝트에 도메인 주도 설계를 적용하려 할 때 이미 비즈니스 도메인을 표현하는 언어가 있을 수 있다.
하지만 그 언어가 DDD의 원칙을 따르지 않아 효과적으로 반영하지 않을 수 있다.
그런 경우 문서화나 소스코드와 같이 제어하기 쉬운 부분부터 올바른 언어를 사용하도록 하자.

회사에서 영어를 사용하지 않는 나라에서 어떤 언어를 사용해야 하냐? 
비즈니스 도메인 엔티티의 이름은 영어 명사로 해야 한다.
그러면 코드에서도 쉽게 동일한 용어를 사용하게 될 것이다.

# 결론
효과적인 커뮤니케이션과 지식 공유는 성공적인 소프트웨어 프로젝트에 필수다.
소프트웨어 엔지니어가 소프트웨어 솔루션을 설계하고 개발하기 위해서는 반드시 비즈니스 도메인을 이해해야 한다.

도메인 주도 설계의 유비쿼터스 언어는 도메인 전문가와 소프트웨어 엔지니어의 지식 간극을 메워주는 효과적인 도구다.
대화, 문서화, 테스트, 다이어그램, 소스코드 등 프로젝트 전반에 걸쳐 모든 이해관계자가 공유된 언어를 사용함으로써 커뮤니케이션과 지식 공유를 강화한다.

효과적인 커뮤니케이션을 위해서는 유비쿼터스 언어에서 반드시 모호성과 암묵적 가정을 제거해야 한다.
언어의 모든 용어는 일관성이 있어야 하며 모호하지 않고 동의어가 없어야 한다.

유비쿼터스 언어를 육성하는 것은 지속적인 과정이다.
프로젝트가 발전함에 따라 더 많은 도메인 지식이 발견된다.
이런 통찰이 유비쿼터스 언어에 반영되는 것이 중요하다.

위키 기반 용어집과 거킨 테스트 같은 도구는 유비쿼터스 언어를 문서화하고 유지보수하는 과정을 상당히 쉽게 해준다.
그러나 효과적인 유비쿼터스 언어의 전제 조건은 언어를 사용해야 한다는 것이다.
즉 유비쿼터스 언어를 모든 프로젝트 관련 커뮤니케이션에서 일관되게 사용해야 한다.
