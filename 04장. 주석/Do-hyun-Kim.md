### 목차

- 주석 🖊
- 주석은 나쁜 코드를 보안하지 못한다 👎
- 코드로 의도를 표현해라! 🗣
- 좋은 주석 😽
- 법적인 주석 👮‍♀️
- 정보를 제공하는 주석 💻
- TODO 주석 🗓
- 나쁜 주석 😵
- 같은 이야기를 중복하는 주석 😵‍💫
- 닫는 괄호에 다는 주석 👄
- 주석으로 처리한 코드 😡



### 주석 🖊

---

- 잘 달린 주석은 그 어떤 정보보다 유용하다.
- 주석은 언제나 실패를 의미한다.  주석 없이는 자신을 표현할 방법을 찾지 못해 할 수없이 주석을 사용한다.
- 주석은 오래될수록 코드에 멀어진다.
- 우리는(간혹 필요할지라도) 주석을 가능한 줄이도록 꾸준히 노력해야 한다.

### 주석은 나쁜 코드를 보안하지 못한다 👎

---

- 코드에 주석을 추가하는 일반적인 이유는 코드의 품질이 나쁘기 때문이다.

### 코드로 의도를 표현해라! 🗣

---

- 확실히 코드만으로 의도를 설명하기 어려운 경우가 존재한다. 

```swift
//직원에게 복지 혜택을 받을 자격이 있는지 검사한다.
if ((employee.flags & HOURLY_FLAG) && employee.age > 65))

// 코드로 명확하게 의도를 표현 할 수 있다.
if (employee.isEligibleForFullBenefits())
```



### 좋은 주석 😽

---

- 어떤 주석은 필요하거나 유익하다. 하지만 정말로 좋은 주석은, 주석을 달지 않을 방법을 찾아낸 주석이라는 것이다!



### 법적인 주석 👮‍♀️

----

- 회사가 정립한 구현 표준에 맞춰 법적인 이유로 특정 주석을 넣으라고 명시한다.

> Example) 각 소스 파일 첫머리에 주석으로 들어가는 저작권 정보와 소유권 정보는 필요하고도 타당하다.

```swift
// Copyright (C) 2003,2004,2005 by Object Mentor, Inc All rights reserved.
// GNU General Public License 버전 2 따르는 조건으로 배포한다.
```

### 정보를 제공하는 주석 💻

---

- 때로는 기본적인 정보를 주석으로 제공하면 편리하다.
- 위와 같은 주석이 유용 할지라도 가능하다면 함수 이름에 정보를 담는 편이 더 좋다.

> Example) 함수 이름을 responderBeingTested로 바꾸면 주석이 필요 없어진다.

```swift
//테스트 중인 Responder 인스턴스를 반환한다.
func responderBeingTestedInstance() -> Responder
```



- 다른 예시 (Date Format)
  - 이왕이면 시각과 날짜를 변환하는 클래스를 만들어 코드를 옮겨주면 더 좋고 깔끔하겠다. 그러면 주석이 필요 없다.

```java
// kk:mm:ss EEE, MMM dd, yyyy 형식 이다.
Pattern timeMatcher = Pattern.compile("\\d*:\\d*:\\d* \\w*, \\w* \\d*, \\d*");
```



### TODO 주석 🗓

---

- 때로는 '앞으로 할 일'을 //TODO 주석으로 남겨두면 편하다.
- 그래도 TODO 주석을 떡칠한 코드는 바람직하지 않다. 그러므로 주기적으로 TODO 주석을 점검해 없애도 괜찮은 주석은 없애라고 권한다.

```swift
//TODO: 현재 팰요하지 않다.
//체크아웃 모델을 도입하면 함수가 필요 없다.
private func makeVersion() throws -> Void {
  
}
```



### 나쁜 주석 😵

---

- 대다수 주석이 이 범주에 속한다.
- 주석을 달기로 결정했다면 충분한 시간을 들여 최고의 주석을 달도록 노력한다.
- 함수의 목적대로 작고 캡슐화가 잘 된 함수에서는 불필요한 주석은 혼란을 일으킬수 있으며 바이트만 낭비한다.

```swift
public func loadProperties() throws -> Void {
  try {
    
  } catch {
    // 속성 파일이 없다면 기본값을 모두 메모리에 읽어 들였다는 의미다.
  }
}
```



### 같은 이야기를 중복하는 주석 😵‍💫

- 같은 코드 내용을 그대로 중복한다.

  > 자칫하면 코드보다 주석을 읽는 시간이 더 오래 걸린다.

- 주석이 코드보다 더 많은 정보를 제공하지 못한다. 코드보다 읽기가 쉽지도 못한다.

```swift
// this.closed 가 true일 때 반환되는 유틸리티 메서드다.
// 타임아웃에 도달하면 예외를 던진다.

public synchronized void waitForClosed(final long timeoutMillis) {
  if(!closed) {
    wait(timeoutMillis);
    if(!closed) {
      throw new Exception("MockResponseSender Could not be closed");
    }
  }
}

```



### 닫는 괄호에 다는 주석 👄

----

-  중첩이 심하고 장황한 함수라면 의미가 있을지 모르지만 작고 캡슐화된 함수에는 잡음일 뿐이다.

> 닫는 괄호에 주석을 달아야겠다는 생각이 든다면 대신에 함수를 줄이려 시도하자.



### 주석으로 처리한 코드 😡

- 주석으로 처리한 코드만큼 밉살스러운 관행도 드물다. 다음과 같은 코드는 작성하지 마라!!
  - 주석으로 처리된 코드는 사람들이 지우기를 주저한다. 

```swift
//func inputStreamResponse() -> Void {
// 
//}
```

