# 창발성

### 창발

> 남이 모르거나 하지 아니한 것을 처음으로 또는 새롭게 밝혀내거나 이루는 일.

## 소프트웨어 품질을 크게 높여주는 네 가지 규칙

- 모든 테스트를 실행한다
- 중복을 없앤다
- 프로그래머 의도를 표현한다
- 클래스와 메서드 수를 최소로 줄인다

## 규칙 1: 모든 테스트를 실행한다

우리는 설계는 의도한 대로 돌아가는 시스템을 내놓아야 합니다.

시스템이 의도한 대로 돌아가는지 검증할 수단이 없다면 노력에 대한 가치는 인정받기 어렵습니다.

**SRP**를 준수하는 클래스는 테스트가 훨씬 더 쉽습니다.

결합도가 높으면 테스트 케이스를 작성하기 어렵습니다.

테스트를 많이 작성할수록 DIP같은 원칙을 적용하고

DI, 인터페이스 추상화등을 사용해 결합도를 낮출 수 있습니다.

이렇게 하면 설계 품질은 높아집니다.

## 규칙 2: 중복을 없애라

### 똑같은 코드는 중복입니다

중복되는 코드가 있다면 따로 추출해서 하나의 함수나 메서드로 만들어서 사용하면 중복을 줄일 수 있습니다.

## 규칙 3 : 프로그래머 의도를 표현하라

자신이 이해하는 코드를 짜기는 쉽습니다.

하지만 코드를 유지보수하는 사람이 코드를 짠 사람의 입장은 다릅니다.

그래서 우리는 유지보수할 사람이 읽기 쉬운 코드를 작성하려고 노력해야합니다.

- 좋은 이름을 선택한다.
- 함수와 클래스 크기를 가능한 줄인다.
    - 작은 클래스나 함수 이름은 짓기도 쉽고, 구현하기도 쉽습니다.
- 표준 명칭을 사용한다.
    - 예를 들어 디자인 패턴은 의사소통과 표현력 강화가 주 목적입니다.
    - COMMAND나 VISITOR와 같은 표현을 사용하면 다른 개발자가 의도를 잘 파악할 수 있습니다.
- 단위 테스트 케이스를 꼼꼼히 작성한다.
    - 테스트 케이스는 예제로 보여주는 문서입니다.

## 규칙 4 : 클래스와 메서드 수를 최소로 줄입니다.

중복을 제거하고, 의도를 표현하고, **SRP**를 준수하다보면 득보다 실이 많아집니다.

클래스와 메서드 크기를 줄이려고 그냥 분리하는 경우도 있습니다.

이런 규칙에 너무 의존하지 말고 실용적인 방식을 택하는 게 좋습니다.
