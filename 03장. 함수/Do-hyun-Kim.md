### 목차

- 작게 만들어라 🧸
- 함수 이름 🧏‍♂️



### 작게 만들어라 🧸

---

- 블록과 들여쓰기

> - 중첩 구조가 생길만큼 함수가 켜져서는 안 된다. if 문, else 문 while문을 주의 하여 사용해야 한다.

- 한 가지만 해라!
  - 함수 이름 아래에서 추상화 수준이 하나인 단계만 수행한다면 그 함수는 한가지 작업만 한다.
  - 단순히 다른 표현이 아니라 의미 있는 이름으로 다른 함수를 추출할 수 있다면 그 함수는 여러 작업을 하는 셈이다.

> 함수는 한가지를 해야한다. (Single Responsibility Principle)
>
> Example) 페이지가 테스트 페이지인지 판단한다 => 그렇다면 설정 페이지와 해체 페이지를 넣는다 => 페이지를 HTML 렌더링 한다

- 함수 당 추상화 수준은 하나로!

  - 함수가 확실히 '한 가지' 작업만 하려면 함수 내 모든 문장의 추상화 수준이 동일해야 한다.
  - 근본 개념과 세부사항을 뒤섞기 시작하면, 함수에 세부사항을 점점 더 추가한다.

- 위에서 아래로 코드 읽기: 내려가기 규칙!

  - 코드는 위에서 아래로 이야기 처럼 읽혀야 좋다.
  - 한 함수 다음에는 추상화 수준이 한 단계 낮은 함수가 온다. 즉 위에서 아래로 프로그램을 읽으면 함수 추상화 수준이 한 번에 한 단계씩 낮아진다.

- Switch 문

  - 다향성을 이용한다.
  - 상속 관계를 숨긴 후에는 절대로 다른 코드에 노출하지 않도록 한다.

  >
  >
  >Protocol 인터페이스를 이용하여 Single Responsibility Principle 과 Open Closed Principle을  막는다.

- 서술적인 이름을 사용하라!

  - 함수가 작고 단순할수록 서술적인 이름을 고르기도 쉬워진다.
  - 이름이 길어도 서술적인 이름이 짧고 어려운 이름보다 좋다.
  - 이름을 넣어 코드를 읽어보면 더 좋다. 
  - 이름을 붙일 때는 일관성이 있어야 한다. 모듈 내에서 함수 이름은 같은 문구, 명사, 동사를 사용한다.

  > includeSetupAndTeardownPages, includeSetupPages, includeSuitesSetupPage, includeSetupPage 등이 좋은 예다.



### 함수 이름 🧏‍♂️

---

- 함수에서 이상적인 인수 개수는 0개(무항)이다
  - 4개 이상(다항)은 특별한 이유가 필요 하다.
- 함수로 Boolean 값을 넘기는 것은 좋지 않다.

> 왜냐 하면 Return 값을 통해 참이면 이걸 하고 거짓 이면 저걸 한다는 의미가 된다!! 하지만 함수는 한가지 일만 해야한다!!

- 때로는 인수 개수가 가변적인 함수도 필요하다. 대표적으로 String.format 메서드가 좋은 예다.

```swift
convenience init(
    format: String,
    arguments argList: CVaListPointer
)
```

> 위 처럼 Swift의 String Format은 두개의 인자 값을 받고 있다.

- 동사와 키워드

  - 함수의 의도나 인수의 순서와 의도를 제대로 표현하려면 좋은 함수 이름이 필수다.
  - 단항 함수는 함수와 인수가 동사/명사 쌍을 이뤄야 한다.

  > Example) Write(name)을 보며 곧바로 이해하며, 이름이 무엇이든 쓴다는 뜻이다. 이처럼 좀 더 나은 이름은 WriteField(name)이다. 그러면 이름(name)이 필드(fIeld)라는 사실이 분명히 드러난다.

  - 함수 이름에 인수 이름을 넣어 인수 순서를 기억할 필요 없도록 한다.

- 부수 효과를 일으키지 마라!

  - 함수 이름에 의미하는 것처럼 한가지 일만 하는 것이 좋다 만약 시간적인 결합이 필요하다면 함수 이름에 분명히 명시한다.

- 출력 인수

  - 일반적으로 출력 인수는 피해야 한다. 함수에서 상태를 변경해야 한다면 함수가 속한 객체 상태를 변경하는 방식을 택한다.

- 명령과 조회를 분리하라!

  - 함수는 뭔가를 수행하거나 뭔가에 답하거나 둘 중 하나만 해야 한다. 둘다 하면 안된다.

  > 함수는 객체 상태를 변경하거나 아니면 객체 정보를 반환하거나 둘중 하나다. 둘다 하면 혼란을 초래한다.

- 오류 코드보다 예외를 사용하라!

  - 명령 함수에서 오류 코드를 반환하는 방식은 명령/조회 분리 규칙을 미묘하게 위반한다.

<span style="color: red; font-weight:bold; "> WRONG </span>

```swift
if (deletePage(page) = E_OK) {
    if (registry.deleteReference(page.name) = E_OK) {
      if (configKeys.deleteKey(page.name.makeKey()) = E_OK){ 
          logger.log('page deleted');
      } else {
          logger.log('configKey not deleted');
      }
    } else {
    	logger.log('deleteReference from registry failed');
    }
} else {
	logger.log('delete failed'); 
    return E_ERR0R;
}
```

<span style="color: green; font-weight:bold; "> Right </span>

```swift
try {
  deletePage(page);
  registry.deleteRefe rence(page.name); 
  configKeys.deleteKey(page.name.makeKey());
}catch (Exception e) { 
	logger.log(e.getMessage());
}
```

- try/Catch 블록은 원래 추하다.  코드 구조에 혼란을 일으키며, 정상 동작과 오류 처리 동작을 뒤섞는다. 그러므로 try/Catch 블록을 별도 함수로 뽑아내는 편이 좋다.

- 오류 처리도 한가지 작업이다.
  - 함수는 '한 가지' 작업만 해야 한다. 오류 처리도 '한 가지' 작업만 속한다. 그로므로 오류를 처리하는 함수는 오류만 처리해야 한다.
- 반복하지 마라!
  - 중복 코드를 제거하므로써 코드의 가독성을 높힌다.
  - 객체 지향 프로그래밍에서는 중복 코드를 부모 클래스로 몰아 중복을 없앤다.
- 함수를 어떻게 짜죠?
  - 소프트웨어를 짜는 행위는 여느 글짓기와 비슷하다. 논문이나 기사를 작성할 때는 먼저 생각을 기록한 후 읽기 좋게 다듬는다.
  - 원하는 대로 읽힐 때까지 말을 다듬고 문장을 고치고 문단을 정리한다.