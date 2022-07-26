# 시스템

## 시스템 제작과 시스템 사용을 분리해야 합니다.

> 소프트웨어 시스템은 준비과정과 런타임 로직을 분리해야 합니다.

시작 단계는 모든 애플이케이션이 풀어야 할 관심사입니다.

아래 예시처럼, 대다수 애플리케이션은 시작 단계라는 관심사를 분리하지 않습니다.

```java
public Service getService() {
    if (service == null) {
        service = new MyServiceImpl(...);
    }
    return service;
}
```

위 방식은 초기화 지연 (Lazy Initialization) 혹은 계산 지연 (Lazy Evaluation) 이라는 기법입니다.

장점은 
- 실제로 필요할 때 까지 객체를 생성하지 않습니다.
- 그래서 불필요한 부하가 걸리지 않습니다.
- 어플리케이션을 시작하는 시간이 빨라집니다.
- null 포인터를 반환하지 않습니다.

하지만 단점은 
- `getService()` 메서드가 `MyServiceImpl` 과 생성자 인수에 명시적으로 의존합니다.
- 단위 테스트에서 getService 메서드를 호출하기 때문에 `MyServiceImpl` 가 무거운 경우 문제입니다.
- 일반 런타임 로직에 객체 생성 로직을 섞어서 service가 null인 경로와 null이 아님 경로등 모든 경로를 테스트 해야합니다.

그리고 결정적으로 `MyServiceImpl` 가 모든 상황에 적합한 객체인지 모릅니다.

체계적이고 탄탄한 시스템을 만들고 싶다면 손쉬운 기법으로 모듈성을 깨선 안됩니다.

## Main 분리

main 함수에서 시스템에 필요한 객체를 생성하고 그 다음에 애플리케이션에 넘깁니다.

애플리케이션은 main이나 객체가 생성되는 과정을 전혀 모릅니다.

단지 모든 객체가 적절히 준비됐다고 가정합니다.

## 팩토리

객체가 생성되는 시점이 어플리케이션에서 결정할 필요도 생깁니다.

주문처리 시스템에서, 주문 처리를 하려면 주문 정보를 담은 객체를 생성해야 하는데

이 때 ABSTRACT FACTORY 패턴을 사용합니다.

그렇게 되면 생성하는 시점은 애플리케이션에서 결정하지만 생성하는 코드는 애플리케이션은 모릅니다.

## 의존성 주입

사용과 제작을 분리하는 강력한 메커니즘 하나가 **의존성 주입 (Dependency Injection)** 입니다.

... 먼말인지 모르겠다

## 확장

처음부터 올바르게 시스템은 만들 수 있다는 믿음은 미신입니다.

대신 우리는 주어진 사용자 스토리에 맞춰 시스템을 구현해야 합니다.

### 횡단 관심사

관심사는 애플리케이션의 자연스러운 경계를 넘나드는 경향이 있습니다.

이런 관심사 또한 모듈화할 수 있습니다.

AOP에서 관점이라는 모듈 구성 개념은

"특정 관심사를 지원하려면 시스템에서 특정 지점들이 동작하는 방식을 일관성 있게 바꿔야 한다." 라고 명시합니다.

다음은 자바에서 사용하는 관점 혹은 관점과 유사한 메커니즘 새 가지입니다.

### 자바 프록시

단순한 상황에 적합합니다.

프록시를 사용하면 깨끗한 코드를 작성하기 어렵습니다.

AOP 해법에 필요한 시스템 단위로 실행 지점을 명시하는 메커니즘도 제공하지 않습니다.

### 순수 자바 AOP 프레임워크

TODO: 자바를 잘 몰라서 뭔말인지 모르겠어서 나중에 추가하겠습니다... 

### AspectJ 관점

> 언어 차원에서 관점을 모듈화 구성으로 지원하는 자바 언어확장

AspectJ는 관점을 분리하는 좋은 도구입니다.

하지만 사용법을 익혀야 합니다.

AspectJ 애너테이션 폼은 새로운 도구와 새로운 언어라는 부담을 어느 정도 완화합니다. 

## 테스트 주도 시스템 아키텍처 구축

관점으로 관심사를 분리하는 방식은 좋습니다.

어플리케이션 도메인 논리를 POJO로 작성할 수 있다면

코드 수준에서 아키텍처 관심사를 분리할 수 있다면 테스트 주도 아키텍처 구축이 가능해집니다.

### 의사 결정을 최적화해야 합니다

가장 적합한 사람에게 맡기면 좋습니다.

마지막 순간까지 결정을 미뤄서 최선의 결정을 내려야 합니다.

성급한 결정은 불충분한 지식으로 내린 결정입니다.

### 명백한 가치가 있을 때 표준을 현명하게 사용해야 합니다.

EJB2는 단지 표준이라는 이유만으로 많은 팀이 사용했습니다.

가볍고 간단한 설계로 충분했을 프로젝트에도 표준이라서 채택했습니다.

이건 좋지 않습니다.

목적에 맞는 선택을 해야 합니다.

### 시스템은 도메인 특화 언어가 필요합니다.

**DSL(Domain-Specific Language) 은 간단한 스크립트 언어나 표준 언어로 구현한 API를 의미합니다.**

도메인 전문가가 사용하는 언어로 도메인 논리를 구현하면 도메인을 잘못 구현할 가능성이 줄어듭니다.

DSL을 사용하면 고차원 정책부터 저차원 정책까지 모든 추상화 수준과 도메인을 POJO로 표현할 수 있습니다.

