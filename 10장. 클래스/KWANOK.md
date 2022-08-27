# 클래스

## 클래스 체계

1. `static` `public` const
2. `static` `private` variable
3. `instance` `private` variable
4. `public` function
5. `private` function (자신을 호출하는 `public` function 밑에)

보통 위 순서로 클래스 파일에 배치가 됩니다.

그리고 사람들은 위 규칙에 익숙합니다. 

## 캡슐화

변수와 Util 함수는 가능한 공개하지 않는 편이 낫지만 반드시 숨기라는 법은 없습니다.

변수나 Util 함수를 protected로 선언해 테스트코드에 접근을 허용하기도 합니다.

그래도 캡슐화를 풀어주는 결정은 언제나 최후의 수단입니다.

## 클래스는 작아야 합니다.

code line 수를 말하는 것이 아니라 책임이 작아야 합니다.

클래스가 하는 일이 많아지면 안됩니다.

`SuperDashboard` 처럼 클래스의 이름이 모호해지면 클래스에 여러 책음을 떠안겼다는 증거입니다.

그리고 클래스 이름은 `if`, `and`, `or`, `but` 단어를 사용하지 않고서 25단어 내외로 가능해야 합니다.

## 단일 책임 원칙 (Single Responsibility Principle)

클래스나 모듈을 변경할 이유가 하나여야 한다는 원칙입니다.

만약 이 규칙에 위반된다고 생각되는 게 있다면, 클래스를 분리해서 독자적인 클래스로 만드는 게 좋습니다.

우리는 돌아가는 소프트웨어에 초점을 맞춥니다.

소프트웨어가 돌아가기만 하면 만족이 되기 때문에,

그 다음 단계인 깨끗하고 체계적인 소프트웨어를 신경쓰지 않게 됩니다.

그래서 클래스의 책임을 나누지 않고, 다음 문제로 넘어가는 경우가 많습니다.

이렇게 SRP를 반복해서 위반하게 되면 불분명함이 많아져, 더러운 코드가 됩니다.

> 큰 클래스 몇 개가 아니라 여럿으로 이뤄진 시스템이 더 바람직합니다.
> 작은 클래스는 각자 맡은 책임이 하나입니다.
> 변경할 이유도 하나입니다.
> 다른 작은 클래스와 협력해 시스템이 필요한 동작을 수행합니다.
 
## 응집도 (Cohesion)

### 클래스는 인스턴스 변수 수가 작아야 합니다.

클래스 메서드는 클래스 인스턴스 변수를 하나 이상 사용해야 합니다.

일반적으로 메서드가 변수를 더 많이 사용할수록 메서드와 클래스는 응집도가 더 높습니다.

모든 인스턴스 변수를 메서드마다 사용하는 클래스는 응집도가 가장 높습니다.

다음은 응집도가 높은 클래스의 예시입니다.

list를 제외하면 `topOfStack`와 `elements` 를 모두 사용하고 있습니다.

```java
public class Stack {
    private int topOfStack = 0;
    List<Integer> elements = new LinkedList<Integer>();
    
    public int size() {
        return topOfStack;
    }
    
    public void push(int elemnet) {
        topOfStack++;
        elements.add(element);
    }
    
    public int pop() throws PoppedWhenEmpty {
        if (topOfStack == 0) {
            throw new PoppedWhenEmpty();
        }
        int element = elements.get(--topOfStack);
        elements.remove(topOfStack);
        return element;
    }
}
```

함수를 작게, 매개변수 목록을 짧게 전략을 따르다 보면 몇몇 메서드만이 사용하는 인스턴스 변수가 많아집니다.

이는 클래스를 쪼개야 한다는 신호입니다.

응집도가 높아질수록 변수와 메서드를 분리하는 것이 좋습니다.

## 응집도를 유지하면 작은 클래스 여럿이 나옵니다.

큰 함수를 작은 함수로 나누기만 해도 클래스가 늘어납니다.

큰 함수에서 일부를 분리하려고 하는데 여기엔 변수 4개가 필요합니다.

이럴 땐 함수에 인수로 변수 4개를 넘기는 것이 아니라 인스턴스 변수로 승격시키는 것이 좋습니다.

이렇게 하면 응집력을 잃습니다.

나머지는 쓰지 않는 분리되는 함수 하나 때문에 인스턴스 변수를 추가하는 것이기 때문이죠

이런 경우에는 이 함수를 다른 클래스로 뺀다면 해결이 됩니다.

인스턴스 변수 4개에 함수가 하나니까 응집도가 높아지니까요.

## 변경하기 쉬운 클래스

대다수 시스템은 지속적인 변경이 일어납니다.

그리고 변경은 위험합니다.

깨끗한 시스템은 변경에 위험을 낮춰줍니다.

```java
public class Sql{
    public Sql(String table, Column[] columns)
    public String create()
    public String insert(Object[] fields)
    public String selectAll()
    ...
    private String columnList(Column[] columns) 
    private String valuesList(Object[] fields, final Column[] columns)
    업데이트 함수는 없습니다.
}
```

위 코드에서 `update()` 를 추가하려면 Sql 클래스에 손대야 합니다.

Sql 클래스에 손대면 다른 코드에 영향을 줄 잠정적인 위험이 존재합니다.
 
테스트도 다시 해야하죠.

이런 위험을 줄일 수 있는 방법은 Sql 클래스에서 파생하는 클래스로 만드는 것입니다.

`valuesList()` 같은 비공개 메서드는 해당하는 파생 클래스로 옮겼습니다.

모든 파생 클래스가 공통으로 사용하는 비공개 메서드는 `Where`와 `ColumnList` 라는 Util 클래스에 넣을 수 있습니다.

```java
abstract public class Sql {
    public Sql(String table, Column[] columns)
    abstract public String generate();
}

public class CreateSql extends Sql {
    public CreateSql(String table, Column[] columns)
    @Override public String generate()
}

public class SelectSql extends Sql {
    public SelectSql(String table, Column[] columns)
    @Override public String generate()
}

public class InsertSql extends Sql {
    public InsertSql(String table, Column[] columns, Object[] fields)
    @Override public String generate()
    private String valuesList(Object[] fields, final Column[] columns)
}

...

public Class Where {
    public Where(String criteria)
    public String generate()
}

public Class ColumnList {
    public ColumnList(Column[] columns)
    public String generate()
}
```

이렇게 짜면 함수 하나를 수정했다고 다른 함수가 망가질 위험도 사라집니다.

`update()`를 추가한다고 기존 클래스를 변경하지 않아도 됩니다.

그리고 위 코드는 OCP (Open-Closed-Principle) 를 지원합니다.

확장에 개방적이고, 수정에 폐쇄적입니다.

새 기능을 추가하는 건 Sql의 파생 클래스를 생성하면 되지만, 다른 클래스를 닫아놓는 방식으로 수정에는 폐쇄적입니다.

## 변경으로부터 격리

요구사항은 변하기 마련입니다.

우리는 변경으로부터 자유로운 코드를 만들어야 합니다.

### interface

Portfolio 클래스는 외부 API를 사용해 가격을 계산하는데 이러면 테스트짜기 어렵습니다.

데이터가 우리맘대로 결정되지 않기 때문이죠

이런 경우에는 인터페이스를 활용할 수가 있습니다.

가격을 가져오는 개념을 추상화하는 거죠

```java
public interface StockExchange {
    Money currnetPrice(String symbol);
}
```

```java
public Portfolio {
    private StockExchanger exchange;
    public Portfolio(StockExchange exchange) {
        this.exchange = exchange;
    }
}
```

이렇게 결합도가 줄어들면 StockExchange 인터페이스를 구현한 `Mock`을 활용해 테스트하기 좋은 코드가 됩니다.

## 마무리

> 이 책을 읽고 느낀 점은 프로그래머에게 Java는 중요합니다.


