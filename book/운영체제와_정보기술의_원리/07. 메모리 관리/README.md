# 1. 주소 바인딩

프로그램이 실행을 위해 메모리에 적재되면 그 프로세스를 위한 독자적인 주소 공간이 생성된다.
이 주소를 논리적 주소 혹은 가상 주소라고 부른다.
CPU는 이와 같이 프로세스마다 독립적으로 갖는 논리적 주소에 근거해 명령을 실행한다.
논리적 주소는 각 프로세스마다 독립적으로 할당되며 0번지부터 시작된다.
반면 물리적 주소는 물리적 메모리에 실제로 올라가는 위치를 말한다.

프로세스가 실행되기 위해서는 해당 프로그램이 물리적 메모리에 올라가 있어야 한다.
또한 CPU가 기계어 명령을 수행하기 위해 논리적 주소를 통해 메모리 참조를 하게 되면 해당 논리적 주소가 물리적 메모리의 어느 위치에 매핑되는지 확인해야 한다. 이렇게 프로세스의 논리적 주소를 물리적 메모리 주소로
연결시켜주는 작업을 주소 바인딩이라고 한다.

### 주소 바인딩의 방식

- 컴파일 타임 바인딩

물리적 메모리 주소가 프로그램을 컴파일할 때 결정되는 주소 바인딩 방식이다. 이 바인딩 기법은 컴파일을 하는 시점에 해당 프로그램이 물리적 메모리의 몇 번지에 위치할 것인지를 결정한다. 이와 같은 주소 바인딩
방식에서 프로그램이 올라가 있는 물리적 메모리의 위치를 변경하고 싶다면 컴파일을 다시 하는 수고가 필요하다.

- 로드 타임 바인딩

프로그램의 실행이 시작될 때에 물리적 메모리 주소가 결정되는 주소 바인딩 방식이다. 이 바인딩 방식에는 로더의 책임하에 물리적 메모리 주소가 부여되며 프로그램이 종료될 때까지 물리적 메모리상의 위치가 고정된다.
로더란 사용자 프로그램을 메모리에 적재시키는 프로그램을 말한다. 로드 타임 바인딩은 컴파일러가 재배치 가능 코드를 생성한 경우에 가능한 주소 바인딩 방식이다.

- 실행시간 바인딩

프로그램이 실행을 시작한 후에도 그 프로그램이 위치한 물리적 메모리상의 주소가 변경될 수 있는 바인딩 방식이다. 이 방식에는 CPU가 주소를 참조할 때마다 해당 데이터가 물리적 메모리의 어느 위치에 존재하는지 주소
매핑 테이블을 이용해 바인딩을 점검해야 한다. 또한 다른 방식들과 달리 실행시간 바인딩 방식이 가능하기 위해서는 기준 레지스터와 한계 레지스터를 포함해 MMU (Memory Management Unit) 라는
하드웨어적인 지원이 뒷받침되어야 한다. MMU는 논리적 주소를 물리적 주소로 매핑해주는 하드웨어 장치이다. 한계 레지스터는 프로세스가 자신의 주소 공간을 넘어서는 메모리 참조를 하려고 하는지 체크하는 용도로 사용되며
다른 프로세스의 주소 영역에 접근하려고 시도하면 트랩을 발생시켜 해당 프로세스를 강제종료시킨다.

# 2. 메모리 관리와 관련된 용어

## 1) 동적로딩

동적로딩은 여러 프로그램이 동시에 메모리에 올라가서 수행되는 다중 프로그래밍 환경에서 메모리 사용의 효율성을 높이기 위해 사용하는 기법 중 하나이다. 동적로딩에서는 프로세스가 시작될 때 그 프로세스의 주소 공간
전체를 메모리에 다 올려놓는 것이 아니라 해당 부분이 불릴 때 그 부분만을 메모리에 적재하는 방식을 사용한다. 즉 프로세스 내에서 실행에 필요한 부분이 실제로 불릴 때마다 메모리에 적재하는 것을 말한다.

## 2) 동적연결

연결이란 프로그래머가 작성한 소스코드를 컴파일하여 생성된 목적 파일과 이미 컴파일된 라이브러리 파일들을 묶어 하나의 실행파일을 생성하는 과정을 말한다. 동적연결은 컴파일을 통해 생성된 목적 파일과 라이브러리 파일
사이의 연결을 프로그램의 실행 시점까지 지연시키는 기법이다. 동적연결에서는 라이브러리가 실행 시점에 연결된다. 즉 실행 파일에 라이브러리 코드가 포함되지 않으며 프로그램이 실행되면서 라이브러리 함수를 호출할 때가
되어서야 라이브러리에 대한 연결이 이루어진다. 동적연결을 가능하게 하기 위해 실행파일의 라이브러리 호출 부분에 해당 라이브러리의 위치를 찾기 위한 스텁이라는 작은 코드를 둔다. 라이브러리 호출 시 스텁을 통해 해당
라이브러리가 메모리에 이미 존재하는지 살펴보고 그럴 경우 그 주소의 메모리 위치에서 직접 참조하며 그렇지 않을 경우 디스크에서 동적 라이브러리 파일을 찾아 메모리로 적재한 후 수행하게 된다. 동적연결에서는 다수의
프로그램이 공통으로 사용하는 라이브러리를 메모리에 한 번만 적재하므로 메모리 사용의 효율성을 높일 수 있다.

## 3) 중첩

중첩이란 프로세스의 주소 공간을 분할해 실제 필요한 부분만을 메모리에 적재하는 기법을 말한다. 중첩은 초창기 컴퓨터 시스템에서 물리적 메모리 크기 제약으로 인해 하나의 프로세스조차도 메모리에 한꺼번에 올릴 수 없을
때 프로세스의 주소 공간을 분할해서 당장 필요한 일부분을 메모리에 올려 실행하고 해당 부분에 대한 실행이 끝난 후에 나머지 부분을 올려 실행하는 기법을 뜻한다.

## 4) 스와핑

스와핑이란 메모리에 올라온 프로세스의 주소 공간 전체를 디스크의 스왑 영역에 일시적으로 내려놓는 것을 말한다. 이때 스왑 영역은 백킹스토어라고도 부르며 디스크 내에 파일 시스템과는 별도로 존재하는 일정 영역을
말한다. 스왑 영역은 프로세스가 수행 중인 동안에만 디스크에 일시적으로 저장하는 공간이므로 저장 기간이 상대적으로 짧은 저장공간이라고 할 수 있다.

스왑 영역은 다수의 사용자 프로세스를 담을 수 있을 만큼 충분히 큰 저장공간이어야 하고 어느 정도의 접근 속도가 보장되어야 한다. 한 가지 주의할 점은 스와핑이라는 개념이 프로세스가 종료되어 그 주소 공간을 디스크로
내쫓는 것이 아니라 특정한 이유로 수행 중인 프로세스의 주소 공간을 일시적으로 메모리에서 디스크로 내려놓는 것을 의미한다는 점이다. 특별히 스와핑에서 일어나는 작업의 방향에 따라 디스크에서 메모리로 올리는 작업을
스왑 인 메모리에서 디스크로 내리는 작업을 스왑 아웃이라고 부른다.

스와핑이 일어나는 과정은 일반적으로 스와퍼라고 불리는 중기 스케줄러에 의해 스왑 아웃시킬 프로세스를 선정한다. 스왑 아웃 대상으로 선정된 프로세스에 대해서는 현재 메모리에 올라가 있는 주소 공간의 내용을 통째로
디스크 스왑 영역에 스왑 아웃시키게 된다. 스와핑의 가장 중요한 역할은 메모리에 존재하는 프로세스의 수를 조절하는 것이다. 너무 많은 프로그램이 메모리에 동시에 올라오게 되면 프로세스당 할당되는 메모리의 양이
지나치게 적어져 시스템 전체의 성능이 크게 떨어진다. 스와핑은 이러한 문제를 해결하기 위해 몇몇 프로그램을 통째로 디스크의 스왑영역으로 내쫓음으로써 메모리에 남아 있는 프로그램들에게 실행에 필요한 적절한 메모리
공간을 보장한다.

# 3. 물리적 메모리의 할당 방식

물리적 메모리는 운영체제 상주 영역과 사용자 프로세스 영역으로 나뉘어 사용된다.
운영체제 상주 영역은 인터럽트 백터와 함께 물리적 메모리의 낮은 주소 영역을 사용하며 운영체제 커널이 이곳에 위치하게 된다.
한편 사용자 프로세스 영역은 물리적 메모리의 높은 주소 영역을 사용하며 여러 사용자 프로세스들이 이곳에 적재되어 실행된다.

사용자 프로세스 영역의 관리 방법은 프로세스를 메모리에 올리는 방식에 따라 연속할당 방식과 불연속할당 방식으로 나누어 볼 수 있다.

## 1) 연속할당 방식

연속할당 방식은 프로세스를 메모리에 올릴 때 그 주소 공간을 여러 개로 분할하지 않고 물리적 메모리의 한 곳에 연속적으로 적재하는 방식이다. 이 방식은 물리적 메모리를 고정된 크기의 분할로 미리 나누어 놓는지 그렇지
않은지에 따라 고정분할 방식과 가변분할 방식으로 나뉜다.

### 고정분할 방식

고정분할 방식은 물리적 메모리를 주어진 개수만큼의 영구적인 분할로 미리 나누어두고 각 분할에 하나의 프로세스를 적재해 실행시킬 수 있게 한다. 이때 분할의 크기는 모두 동일하게 할 수도 있고 서로 다르게 할 수도
있다. 하지만 두 방식 모두 하나의 분할에는 하나의 프로그램만을 적재할 수 있다. 따라서 동시에 메모리에 올릴 수 있는 프로그램의 수가 고정되어 있으며 수행 가능한 프로그램의 최대 크기 또한 제한된다는 점에서
가변분할 방식에 비해 융통성이 떨어진다.

또한 고정분할 방식에는 외부조각과 내부조각이 발생할 수 있다. 외부조각이란 프로그램의 크기보다 분할의 크기가 작은 경우 해당 분할이 비어 있는데도 불구하고 프로그램을 적재하지 못하기 때문에 발생하는 메모리 공간을
의미한다. 외부조각은 특정 프로그램에 배당된 공간이 아니기 때문에 만약 이 외부조각의 크기보다 작은 크기의 프로그램이 도착한다면 그 프로그램을 외부조각에 적재시킬 수 있다. 반면 내부조각은 프로그램의 크기보다 분할의
크기가 큰 경우 해당 분할에 프로그램을 적재하고 남는 메모리 공간을 뜻한다. 즉 하나의 분할 내부에서 발생하는 사용되지 않는 메모리 조각을 말한다. 내부조각은 특정 프로그램에 이미 배당된 공간으로 볼 수 있으므로
내부조각에 수용할 수 있는 작은 프로그램이 있다 해도 공간을 활용할 수 없어 메모리 낭비가 된다.

### 가변분할 방식

가변분할 방식은 메모리에 적재되는 프로그램의 크기에 따라 분할의 크기, 개수가 동적으로 변하는 방식을 말한다. 가변분할 방식에서 분할의 크기를 프로그램의 크기보다 일부러 크게 할당하지는 않기 때문에 내부조각은
발생하지 않는다. 하지만 이미 메모리에 존재하는 프로그램이 종료될 경우 중간에 빈 공간이 발생하게 되며 이 공간이 새롭게 시작되는 프로그램의 크기보다 작을 경우 외부조각이 발생할 가능성이 있다.

가변분할 방식에서 주요하게 다루는 쟁점 중 하나는 주소 공간의 크기가 n인 프로세스를 메모리에 올릴 때 물리적 메모리 내 가용 공간 중 어떤 위치에 올릴 것인지를 결정하는 문제이다. 이를 동적 메모리 할당 문제라고
부른다. 연속할당 기법에서는 새로운 프로세스를 메모리에 올리기 위해 프로세스의 주소 공간 전체를 담을 수 있는 가용 공간을 찾아야 한다. 물리적 메모리 내에는 다양한 크기의 가용 공간들이 흩어져 있으며 이렇게
존재하는 가용 공간들을 효율적으로 관리하기 위해 운영체제는 이미 사용 중인 메모리 공간과 사용하고 있지 않은 가용 공간에 대한 정보를 각각 유지하고 있다.

동적 메모리 할당 문제를 해결하는 대표적인 방법

1. 최초적합 방법

크기가 n 이상인 가용 공간 중 가장 먼저 찾아지는 곳에 프로세스를 할당하는 방법이다. 이 방법은 가용 공간을 모두 탐색하는 방법이 아니므로 시간적인 측면에서 효율적이다.

2. 최적적합 방법

크기가 n 이상인 가장 작은 가용 공간을 찾아 그곳에 새로운 프로그램을 할당하는 방법이다. 이 방법은 가용 공간들의 리스트가 크기순으로 정렬되어 있지 않은 경우에 모든 가용 공간 리스트를 탐색해야 하므로 시간적
오버헤드가 발생하고 다수의 매우 작은 가용 공간들이 생성될 수 있다는 단점이 존재하지만 공간적인 측면에서 효율적이다.

3. 최악적합 방법

가용 공간 중에서 가장 크기가 큰 곳에 새로운 프로그램을 할당하는 방법이다. 이 방법은 최적적합 방법과 마찬가지로 모든 가용 공간 리스트를 탐색해야 하는 오버헤드가 발생한다. 또한 상대적으로 더 큰 프로그램을 담을
수 있는 가용 공간을 빨리 소진한다는 문제점이 있다.

한편 가변분할 방식에서 발생하는 외부조각 문제를 해결하기 위한 방법으로 컴팩션이라는 것이 있다. 컴팩션은 물리적 메모리중에서 프로세스에 의해 사용 중인 메모리 영역을 한쪽으로 몰고 가용 공간들을 다른 한쪽으로 모아서
하나의 큰 가용 공간을 만드는 방법이다. 이는 비용이 매우 많이 드는 작업이므로 중간에 일부 가용 공간이 발생하더라도 가급적 적은 수의 메모리 이동으로 효율적인 컴팩션을 수행하는 방법이 필요한데 매우 복잡한
문제이다. 또한 컴펙션은 효율적인 메모리 공간의 사용을 위해 수행 중인 프로세스의 물리적 메모리 위치를 옮겨야 하므로 프로그램의 실행 도중에 프로세스의 주소가 동적으로 바뀔 수 있는 실행시간 바인딩 방식이 지원되는
환경에서만 수행될 수 있다.

## 2) 불연속할당 기법

불연속할당 기법이란 하나의 프로세스가 물리적 메모리의 여러 위치에 분산되어 올라갈 수 있는 메모리 할당 기법을 말한다. 불연속할당 방식에는 하나의 프로그램을 반할하는 기준에 따라 동일한 크기로 나누어 메모리에 올리는
페이징 기법과 크기는 일정하지 않지만 의미 단위로 나누어 메모리에 올리는 세그먼테이션 기법 그리고 세그먼테이션을 기본으로 하되 이를 다시 동일 크기의 페이지로 나누어 메모리에 올리는 페이지드 세그먼테이션 기법 등이
있다.

# 4. 페이징 기법

페이징 기법이란 프로세스의 주소 공간을 동일한 크기의 페이지 단위로 나누어 물리적 메모리의 서로 다른 위치에 페이지들을 저장하는 방식을 말한다. 페에징 기법에서는 각 프로세스의 주소 공간 전체를 물리적 메모리에 한꺼번에 올릴 필요가 없으며 일부는 백킹스토어에 일부는 물리적 메모리에 혼재시키는 것이 가능하다. 페이징 기법에서는 물리적 메모리를 페이지와 동일한 크기의 프레임으로 미리 나누어둔다. 이는 메모리에 올리는 단위가 동일한 크기의 페이지 단위이므로 메모리를 같은 크기로 미리 분할해두더라도 빈 프레임이 있으면 어떤 위치이든 사용될 수 있기 때문이다. 따라서 페이징 기법은 연속할당에서 발생했던 동적 메모리 할당 문제가 발생하지 않는다는 장점을 가진다.

페이징 기법에서는 주소 변환 절차가 연속할당 방식에 비해 다소 복잡하다. 이는 하나의 프로세스라 하더라도 페이지 단위로 물리적 메모리에 올리는 위치가 상이하므로 논리적 주소를 물리적 주소로 변환하는 작업이 페이지 단위로 이루어져야 하기 때문이다. 따라서 페이징 기법에서는 모든 프로세스가 각각의 주소 변환을 위한 페이지 테이블을 가지며 이 테이블은 프로세스가 가질 수 있는 페이지 개수만큼 주소 변환 엔트리를 가지고 있게 된다.

페이징 기법에서는 프로세스의 주소 공간과 물리적 메모리가 모두 같은 크기 페이지 단위로 나누어지기 때문에 빈 공간은 어느 곳이든 활용할 수 있다. 따라서 메모리상의 가용 공간의 크기가 작아서 빈 공간임에도 활용되지 못하는 외부조각 문제가 발생하지 않는다. 그러나 프로그램의 크기가 항상 페이지 크기의 배수가 된다는 보장이 없기 때문에 프로세스의 주소 공간 중 제일 마지막에 위치한 페이지에서는 내부조각이 발생할 가능성이 있다.

## 1) 주소 변환 기법
페이징 기법에서는 CPU가 사용하는 논리적 주소를 페이지 번호(p)와 페이지 오프셋(d)으로 나누어 주소 변환에 사용한다.
페이지 번호는 각 페이지별 주소 변환 정보를 담고 있는 페이지 테이블 접근 시 인덱스로 사용되 해당 인덱스의 항목에는 그 페이지의 물리적 메모리상의 기준 주소 즉 시작 위치가 저장된다.
따라서 특정 프로세스의 p번째 페이지가 위치한 물리적 메모리의 시작 위치를 알고 싶다면 해당 프로세스의 페이지 테이블에서 p번째 항목을 찾아보면 된다.
한편 페이지 오프셋은 하나의 페이지 내에서 변위를 알려준다.
따라서 기준 주소값에 변위를 더함으로써 요청된 논리적 주소에 대응하는 물리적 주소를 얻을 수 있다.

## 2) 페이징 테이블의 구현
페이지 테이블은 페이징 기법에서 주소 변환을 하기 위한 자료구조로 물리적 메모리에 위치하게 된다.
현재 CPU에서 실행 중인 프로세스의 페이지 테이블에 접근하기 위해 운영체제는 2개의 레지스터를 사용하는데 이들은 각각 페이지 테이블 기준 레지스터와 페이지 테이블 길이 레지스터로 불린다.
페이지 테이블 기준 레지스터는 메모리 내에서의 페이지 테이블의 시작 위치를 가리키며 페이지 테이블 길이 레지스터는 페이지 테이블의 크기를 보관한다.
페이징 기법에서의 메모리 접근 연산은 주소 변환을 위해 페이지 테이블에 접근하는 것과 변환된 주소에서 실제 데이터에 접근하는 것 이렇게 두 번의 메모리 접근을 필요로 한다.
즉 메모리에 한 번 접근하기 위해서는 매번 메모리에 두 번 접근해야 하는 오버헤드가 뒤따르게 되는 것이다. 이와 같은 페이지 테이블 접근 오버헤드를 줄이고 메모리의 접근 속도를 향상시키기 위해 TLB라고 불리는 고속의 주소 변환용 하드웨어 캐시가 사용되기도 한다.

## 3) 계층적 페이징


## 4) 역페이지 테이블


## 5) 공유 페이지
공유 코드는 메모리 공간의 효율적인 사용을 위해 여러 프로세스에 의해 공통으로 사용될 수 있도록 작성된 코드를 말한다.
공유 코드는 재진입 가능 코드 또는 순수 코드라고도 불리며 읽기전용의 특성을 가지고 있다.
공유 페이지란 공유 코드를 담고 있는 페이지를 말한다.
공유 페이지는 여러 프로세스에 의해 공유되는 페이지이므로 물리적 메모리에 하나만 적재되어 메모리를 좀 더 효율적으로 사용할 수 있게 한다.

## 6) 메모리 보호


# 5. 세그멘테이션
세그멘테이션은 프로세스의 주소 공간을 의미 단위의 세그먼트로 나누어 물리적 메모리에 올리는 방법이다.
하나의 프로세스를 구성하는 주소 공간은 일반적으로 코드, 데이터, 스택 등의 의미 있는 단위들로 구성된다.
세그먼트는 이와 같이 주소 공간을 기능 단위 또는 의미 단위로 나눈 것을 뜻한다.
프로세스의 주소 공간 전체를 크게는 하나의 세그먼트로 볼 수도 있으며 일반적으로는 코드, 데이터, 스택 등의 기능 단위로 세그먼트를 정의한다.
많게는 프로그램을 구성하는 함수 하나하나를 각각 세그먼트로 정의할 수도 있다.
한 가지 주의할 점은 세그먼트가 이와 같이 특정 크기 단위로 나눈 것이 아니라 의미를 가질 수 있는 논리적 단위로 나눈 것이기 때문에 그 크기가 균일하지 않다는 점이다.
세그멘테이션 기법에서는 논리적 주소가 <세그먼트 번호, 오프셋> 으로 나뉘어 사용된다.
세그먼트 번호는 해당 논리적 주소가 프로세스 공간 내에서 몇 번째 세그먼트에 속하는지를 나타내며 오프셋은 그 세그먼트 내에서 얼마만큼 떨어져 있는지에 대한 정보를 나타낸다.

세그멘테이션 기법에서는 주소 변환을 위해 세그멘테이션 테이블을 사용한다.
이떄 세그멘트 테이블의 각 항목은 기준점과 한계점을 가지고 있다.
기준점은 물리적 메모리에서 그 세그먼트의 시작 위치를 나태내고 한계점은 그 세그먼트의 길이를 나타낸다.
페이징 기법에서는 모든 페이지의 길이가 동일하므로 페이지 테이블의 항목에 기준점이라 할 수 있는 페이지 프레임 위치만 유지하고 있으면 된다.
하지만 세그먼테이션 기법에서는 세그먼트 길이가 균일하지 않으므로 세그먼트의 위치 정보 뿐 아니라 길이 정보를 함께 보관하고 있는 것이다.

세그먼테이션 기법에서는 세그먼트 테이블 기준 레지스터와 세그먼트 테이블 길이 레지스터를 사용하게 된다.
세그먼트 테이블 기준 레지스터는 현재 CPU에서 실행 중인 프로세스의 세그먼트 테이블이 메모리의 어느 위치에 있는지 그 시작 주소를 담고 있으며 세그먼트 테이블 길이 레지스터는 그 프로세스의 주소 공간이 총 몇 개의 세그먼트로 구성되는지 즉 세그먼트 개수를 나타낸다.

세그먼테이션 기법에서는 논리적 주소를 물리적 주소로 변환하기 전에 두 가지 사항을 먼저 확인한다.
첫 번째 요청된 세그먼트 번호가 STLR에 저장된 값보다 작은 값인가 하는 점이다.
만약 그렇지 않다면 이는 존재하지 않는 세그먼트에 대한 접근 시도이므로 예외상황을 발생시켜 메모리 접근을 봉쇄해야 할 것이다.
두 번째는 논리적 주소의 오프셋값이 그 세그먼트 길이보다 작은 값인가 하는 점이다.
이는 세그먼트 테이블의 해당 항목에 있는 한계점과 요청된 논리적 주소의 오프셋 값을 비교해 확인하게 된다.
만약 세그먼트 길이를 넘어서는 오프셋 위치에 대한 접근 시도라면 역시 예외상황을 발생시켜 해당 메모리 위치에 대한 접근을 봉쇄하게 된다.
이 두 가지 사항을 모두 만족하는 경우에 한해서 유효한 메모리 접근 요청으로 판단해 주소 변환 작업이 이루어진다.

세그먼테이션 기법에서는 프로그램을 의미 단위로 나누기 때문에 세그먼트의 길이가 균일하지 않다.
따라서 물리적 메모리 관리에서 외부 조각이 발생하게 되며 세그먼트를 어느 가용 공간에 할당할 것인지를 결정하는 문제가 발생한다.
세그먼트를 가용 공간에 할당하는 방식에는 최초적합 방식과 최적적합 방식이 있다.

# 6. 페이지드 세그멘테이션
페이징 기법과 세그먼테이션 기법의 장점만을 취하는 주소 변환 기법으로 페이지드 세그멘테이션 기법을 생각해볼 수 있다.
이 기법은 세그먼테이션 기법과 마찬가지로 프로그램을 의미 단위의 세그먼트로 나눈다.
단 세그먼트가 임의의 길이를 가질 수 있는 것이 아니라 반드시 동일한 크기 페이지들의 집합으로 구성되어야 한다.
그리고 물리적 메모리에 적재하는 단위는 페이지 단위로 한다.
즉 페이지드 세그먼테이션 기법에서는 하나의 세그먼트 크기를 페이지 크기의 배수가 되도록 함으로써 세그먼테이션 기법에서 발생하는 외부조각의 문제점을 해결하며 동시에 세그먼트 단위로 프로세스 간의 공유나 프로세스 내의 접근 권한 보호가 이루어지도록 함으로써 페이징 기법의 약점을 해소한다.

페이지드 세그먼테이션 기법에서는 주소 변환을 위해 외부의 세그먼트 테이블과 내부의 페이지 테이블 이렇게 두 단계의 테이블을 이용한다.
하나의 세그먼트가 여러 개의 페이지로 구성되므로 각 세그먼트마다 페이지 테이블을 가지게 된다.
