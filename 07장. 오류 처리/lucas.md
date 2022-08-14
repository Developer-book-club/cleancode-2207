# 07장. 오류 처리

## ✏️ 정리본

- 오류 처리는 프로그램에 반드시 필요한 요소이다. 잘못된 가능성은 언제나 실패하고 이에 대한 에러 핸들링이 있어야 나중에 그것을 바로 잡을 수 있기 때문이 아닐까 생각한다.
- 오류 처리 코드로 인해 프로그램 논리를 이해하기 어려워진다면 깨끗한 코드라 부르기 어렵다고 볼 수 있다.

### 오류 코드보다 예외를 사용하라

- 오류 플래그를 실정하거나 호출자에게 오류 코드를 반환하는 방법보다는, 오류가 발생하면 예외를 던지는 편이 나음

### Try-Catch-Finally 문부터 작성하라

- 예외가 발생하는 코드를 짤 때는 try-catch-finally 문으로 시작하는 편이 낫다.그러면 try 블록에서 무슨 일이 생기든지 호출자가 기대하는상태를 정의하기 쉬워짐
- 스터디원이 보여준 자료로는 성능상의 이슈가 있다고 하였는데, 리팩토링 관점에서 접근하면 try문 에서 문제가 발생하거나 혹은 예상치 못한 예외가 발생할 시 catch문에서 잡아 낼 수 있기 때문이다.
- 강제로 예외를 일으키는 테스트 케이스를 작성한 후 테스트를 통과하게 코드를 작성하는 방법을 권장함. 그러면 자연스럽게 try 블록의 트랜잭션 범위부터 구현하게 되므로 범위 내에서 트랜잭션 본질을 유지하기 쉬워짐

### 미확인 예외를 사용하라

- 일반적인 어플리케이션의 경우는 의존성의 비용이 너무 크기 때문에 확인된 에러처리를 하지 않는 편이 낫다. 확인된 에러 처리를 하게 되면, 최하위 함수를 변경하게되어 새로운 오류가 생긴다고 가정을 하면, 선언부에 throws절을 추가해야 한다.

### 예외에 의미를 제공하라

- 오류 메시지에 정보를 담아 예외와 함께 던진다. 실패한 연산 이름과 실패 유형도 언급한다. 이를 통해 실패한 코드의 의도를 파악할 수 있고 역추적이 가능하다고 한다.
- 애플리케이션이 로깅 기능을 사용한다면 catch 블록에서 오류를 기록하도록 충분한 정보를 넘겨줌!

### 호출자를 고려해 예외 클래스를 정의하라

- **오류를 잡아내느 방법**을 잘 고민해봐야 한다. 그에 관한 방법으로는 오류를 발생한 위치나 오류가 발생한 컴포넌트로 뷴류 등으로 가능하다.
- 외부 라이브러리에서 던지는 예외를 잡아 변환하는 감싸기 클래스

```
LocalPort port = new LocalPort(12);
try {
  port.open();
} catch (PortDeviceFailure e) {
  reportError(e);
  logger.log(e.getMessage(), e);
}

public class LocalPort {
  private ACMDPort innerPort;

  public LocalPort(int portNumber) {
    innerPort = new ACMEPort(portNumber);
  }

  public void open() {
    try {
      innerPort.open();
    } catch (DeviceResponseException e) { } ...
  }
}
```

- 실제로 외부 API를 사용할 때는 감싸기 기법이 최선
  - 외부 라이브러리와 프로그램 사이에서 의존성이 크게 줄어듬
  - 다른 라이브러리로 갈아타도 비용이 적음
  - 감싸기 클래스에서 외부 API를 호출하는 대신 테스트 코드를 넣어주는 방법으로 프로그램 테스트하기도 쉬워짐

### 정상 흐름을 정의하라

- 비지니스 논리와 요류 처리가 잘 분리된 코드
- 특수 사례 패턴: 클래스를 만들거나 객체를 조작해 특수 사례를 처리하는 방식 클라이언트 코드가 예외적인 상황을 처리할 필요가 없어짐. 클래스나 객체가 예외적인 상황을 캡슐화해서 처리하기 때문!!

```
// Soso
try {
  MealExpenses expenses = expenseReportDAO.getMeals(employee.getID());
  m_total += expenses.getTotal();
} catch(MealExpensesNotFound e) {
  m_total += getMealPerDiem();
}

// Good
MealExpenses expenses = expenseReportDAO.getMeals(employee.getID());
  m_total += expenses.getTotal();
  ...

//version3 특수 사례 패턴 -> 객체를 조작
public class PerDiemMealExpenses implements MealExpenses {
  public int getTotal() {
    // return the per diem default
  }
}
```

### null을 반환하지 마라

- null을 반환한느 코드는 일거리를 늘릴 뿐만 아니라 호출자에게 문제를 떠넘기는 행위
- null을 반환하고 싶다면 차라리 예외를 던지거나 특수 사례 객체를 반환해라
  - 외부 API가 null을 반환한다면 감싸기 메서드를 구현해 예외를 던지거나 특수 사례 객체를 반환하는 방식을 고려함

### null을 전달하지 마라

- null을 넘기지 못하도록 금지하는 정책이 합리적임. 즉, 인수로 null이 넘어오면 코드에 문제가 있다는 이야기
- 깨끗한 코드는 읽기도 좋아야 하지만 안정성도 높아야 함.
- 오류 처리를 프로그램 논리와 분리하면 독립적으로 추론이 가능해지며 코드 유지보수성도 크게 높아진다
