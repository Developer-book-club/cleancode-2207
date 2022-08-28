### 목차

- 미확인 unChecked 예외를 사용하라 🤔
- 예외에 의미를 제공하라 🧐
- 호출자를 고려해 예외 클래스를 정의하라 🥳
- 정상의 흐름을 정의하라 😙
- Null을 반환하지 마라 😫
- Null을 전달하지 마라 😞
- 결론 🥳



### 미확인 unChecked 예외를 사용하라 🤔

---

- 자바 프로그래머들은 확인된(ckecked) 예외의 장단점을 놓고 논쟁을 벌여왔다!!
  - 당시는 확인된 예외를 멋진 아이디어라 생각했다. 실제로도 확인된 예외는 몇가지 장점을 제공한다. 하지만 지금은 안정적인 소프트웨어를 제작하는 요소로 확인된 예외가 반드시 필요하지는 않다는 사실이 분명해졌다. 
  - 우리는 확인된 오류가 치느는 비용에 상응하는 이익을 제공하는지 (철저히) 따져봐야 한다!!
    - 확인된 예외는 OCP(Open Closed Principle)를 위반한다. 메서드에서 확인된 예외를 던졌는데 catch 블록이 세 단계 위에 있다면 그 사이 메서드 모두가 선언부에 해당 예외를 정의해야 한다.
  - 즉, 하위 단계에서 코드를 변경하면 상위 단계 메서드 선언부를 전부 고쳐야 한다는 말이다(Open Closed Principle) 위반!!🤔
- 아주 중요한 라이브러리를 작성한다면 모든 예외를 잡아야 한다. 하지만 일반적인 애플리케이션은 의존성이라는 비용이 이익보다 크다.



### 예외에 의미를 제공하라 🧐

---

- 예외를 던질 때는 전후 상황을 충분히 덧붙인다. 그러면 오류가 발생한 원인과 위치를 찾기가 쉬워진다.
  - 오류 메시지에 정보를 담아 예외와 함계 던진다. 실패한 연산 이름과 실패 유형도 언급 한다. 🫶
  - 애플리케이션이 로깅(Logging) 기능을 사용한다면 catch 블록에서 오류를 기록하도록 충분한 정보를 넘겨준다.



### 호출자를 고려해 예외 클래스를 정의하라 🥳

---

- 오류를 분류하는 방법은 수없이 많다. 오류가 발생한 위치로 분류가 가능하다.
  - 예를 들어, 오류가 발생한 컴포넌트로 분류한다.  아니면 유형으로도 분류가 가능하다.
    - 디바이스 실패, 네트워크 실패, 프로그래밍 오류 등으로 분류 한다.
- 애플리케이션에서 오류를 정의할 때 프로그래머에게 가장 중요한 관심사는 오류를 잡아내는 방법이 되어야 한다.

```java
ACMPort port = new ACMPort(12);

try {
  port.open();
} catch(DeviceResponseException e) {
  reportPortError(e);
  logger.log("Device response exception", e);
} catch (ATM1212UnlockedException e) {
  reportPortError(e);
  logger.log("Unlock exception", e);
} catch (GMXError e) {
  reportPortError(e);
  logger.log("Device response exception");
} finally {
  ...
}

```

- 위 코드는 중복이 심하지만 그리 놀랍지 않다. 대다수 상황에서 우리가 오류를 처리하는 방식은 오류를 일으킨 원인과 무관하게 비교적 일정하다.

  - 1) 오류를 기록한다.
    2) 프로그램을 계속 수행해도 좋은지 확인한다.

    

```java
LocalPort port = new LocalPort(12);

try {
  port.open()
} catch (PortDeviceFailure e) {
  reportError(e);
  logger.log(e.getMessage(), e)
} finally {
  ...
}


public class LocalPort {
  private ACMEPort innerPort;
  
  public LocalPort(int portNumber) {
    innerPort = new ACMEPort(portNumber);
  }
  
  public void open() {
    try {
      innerPort.open();
    } catch (DeviceResponseException e) {
      throw new PortDeviceFailure(e);
    } catch (ATM1212UnlocheckedException e) {
      throw new PortDeviceFailure(e);
    } catch (GMXError e) {
      throw new PortDeviceFailure(e);
    }
  }
  ...
}

```

- LocalPort 클래스는 단순히 ACMEPort 클래스가 던지는 예외를 잡아 변환하는 감싸기(Wrapper)클래스일 뿐이다.

  - LocalPort 클래스처럼 ACMPort를 감싸는 클래스는 매우 유용하다. 실제로 외부 API를 사용할 때는 감싸기 기법이 최선이다.

    - 외부 API를 감싸면 외부 라이브러리와 프로그램 사이에서 의존성이 크게 줄어든다!! 😍 혹은 나중에 다른 라이브러리로 갈아타도 비용이 적다.
    - 감싸기 클래스에서 외부 API를 호추랗는 대신 테스트 코드를 넣어주는 방법으로 프로그램을 테스트 하기도 쉬워진다.
    - 마지막 장점으로 감싸기 기법을 사용하면 특정 업체가 API 설계한 방식에 발목 잡히지 않는다. 프로그램이 사용하기 편리한 API를 정의 하면 그만이다.

    

    

### 정상의 흐름을 정의하라 😙

---

- 외부 API를 감싸 독자적인 예외를 던지고, 코드위에 처리기를 정의해 중단된 계산을 처리한다. 대개는 멋진 처리 방식이지만, 때로는 중단이 적합하지 않을 때도 있다.



```java
try {
  MealExpenses expenses = expenseReportDAO.getMeals(employee.getID());
  m_total += expenses.getTotal();
} catch (MealExpensesNotFound e) {
  m_total += getMealPerDiem();
}
```

- 위 코드는 식비를 비용으로 청구하지 않았다면 일일 기본 식비를 총계에 더한다. 그런데 예외가 논리를 따라가기 어렵게 만든다.
  - 특수 상황을 처리할 필요가 없다면 더 좋지 않을까? 그러면 코드가 훨씬 더 간결해지리라.



```java
MealExpenses expenses = expenseReportDAO.getMeals(employee.getID());
m_total += expenses.getTotal();

public class perDiemMealExpenses implements MealExpenses {
  public int getTotal() {
    // 기본값으로 일일 기본 식비를 반환한다.
  }
}
```

- 위처럼 ExpenseReportDAO를 고쳐 언제나 MealExpense 객체를 반환한다. 청구한 식비가 없다면 일일 기본 식비를 반환 하는 MealExPense 객체를 반환 한다.
- 이를 특수 사례 패턴(SPECIAL CASE PATTERN) 이라 부른다. 클래스를 만들거나 객체를 조작해 특수 사례를 처리하는 방식이다.
  - 클라이언트 코드가 예외적인 상황을 처리할 필요가 없어진다. 클래스나 객체가 예외적인 상황을 캡슐화해서 처리하므로



### Null을 반환하지 마라 😫

---

- Null을 반환하는 코드는 일거리를 늘릴 뿐만 아니라 호출자에게 문제를 떠넘긴다. 누구 하나라도 null 확인을 빼먹는다면 애플리케이션이 통제 불능에 빠질지도 모른다.

```java
public void registerItem(Item item) {
  if (item != null) {
    ItemRegistry registry = peristentStore.getItemRegistry();
    if (registry != null) {
      Item existing = registry.getitem(item.getID());
      if (existing,getBillingPeriod().hasRetailOwner()) {
        existing.register(item);
      }
    }
  }
}
```

- 위 코드에서 둘째 행에 nul 확인이 빠졌다는 사실이 있다!! 만약 per-sistenStore(peristentStore)가 null 이라면 실행 시 NullPotinerException이 발생 하여 위쪽 어디선가. NullPointerException을 잡을지도 모르고 아닐지도 모른다. 
- 실상은 null 확인이 너무 많아 문제다. 메서드에서 null을 반환하고픈 유혹이 든다면 그 대신 예외를 던지거나 특수 사례 객체를 반환한다. 
  - 사용하려는 외부 API가 null을 반환 한다면 감싸기 메서드를 구현해 예외를 던지거나 특수 사례 객체를 반환하는 방식을 고려한다.





Null을 전달하지 마라 😞

---

- 메서드에서 null을 반환하는 방식도 나쁘지만 메서드로 null을 전달하는 방식은 더 나쁘다. 정상적인 인수로 null을 기대하는 API가 아니라면 메서드로 null을 전달하는 코드는 최대한 피한다.



```java
public class MetricsCalculator {
  public double xProjection(Point p1, Point p2) {
    return (p2.x - p1.x) * 1.5;
  }
  ...
}

//injection
calculator.xProjection(null, newPoint(12,13));
```

- 위 와 같이 인수로 null을 전달하게 되면 NullPointerException이 발생한다.

###### 새로운 예외 유형 (NullPointException 처리)

```java
public class MetricsCalculator {
 public double xProjection(Point p1, Point p2) {
   if (p1 == null || p2 == null) {
     throw InvalidArgumentException("Invalid argument for MetricsCalculator.xProjection");
   }
   return (p2.x - p1.x) * 1.5;
 }
}
```

- 위 코드는 NullPointerException 오류를 처리 하기에 기존 코드 보다 조금 나을지도모른다 하지만 위 코드는 InvalidArgumentException을 잡아내는 처리기가 필요하다.

###### 새로운 예외 유형 (NullPointException or InvalidArgumentException 처리)

```java
public class MetricsCalculator {
 public double xProjection(Point p1, Point p2) {
   assert p1 != null : "p1 should not be null";
   assert p2 != null : "p2 should not be null";
   return (p2.x - p1.x) * 1.5;
 }
}
```

- 대다수 프로그래밍 언어는 호출자가 실수로 넘기는 null을 적절히 처리하는 방법이 없다.

  - 애초에 null을 넘기지 못하도록 금지하는 정책이 합리적이다!!  즉, 인수로 null 이 넘어오면 코드에 문제가 있다는 말이다. 이런 정책을 따르면 그만큼 부주의한 실수를 저지를 확률도 작아진다.

  

### 결론 🥳

---

- 깨끗한 코드는 읽기도 좋아야 하지만 안정성도 높아야 한다. 이 둘은 상층하는 목표가 아니다.
- 오류 처리를 프로그램 논리와 분리해 독자적인 사안으로 고려하면 튼튼하고 깨끗한 코드를 작성 할 수 있다.
- 오츄 처리를 프로그램 논리와 분리하면 독립적인 추론이 가능해지며 코드 유지보수성도 크게 높아진다.