# 2장 의미있는 이름

### 의도를 분명히 밝혀라  
- 변수, 함수, 클래스 명이 주석 없이도 명확하게 의도를 드려내야 한다.
- 예시 1
  - Bad
    - int d; // 경과 시간(단위: 날짜)
  - Good
    - int elapsedTimeInDays;
    - int daysSinceCreation;
    - int daysSinceModification;
    - int fileAgeInDays;
- 예시 2

```java
// Bad
public List<int[]> getThem() {
    List<int[]> list1 = new ArrayList<int[]>();
    for (int[] x : theList) {
        if (x[0] == 4) {
            list1.add(x);
        }
    }
    return list1;
}
```
 
```java
// Good
public List<int[]> getFlaggedCells() {
    List<int[]> flaggedCells = new ArrayList<int[]>();
    for (int[] cell : gameBoard) {
        if (cell[STATUS_VALUE] == FLAGGED) {
            flaggedCells.add(cell);
        }
    }
    return flaggedCells;
}
```

### 그릇된 정보를 피하라
- 코드의 의미를 명확히 하기 위해선 그릇된 단서를 남기면 안된다
- 널리 쓰이는 의미가 있는 단어를 다른 의미로 사용해도 X
- 서로 흡사한 이름을 사용하지 않도록 주의한다.
- 개발자에게는 특수한 의미를 가지는 단어(List 등)는 실제 컨테이너가 List가 아닌 이상 accountList와 같이 변수명에 붙이지 말자. 차라리 accountGroup, bunchOfAccounts, accounts등으로 명명하자(Item: Object, Items: List)
- 비슷해 보이는 명명에 주의하자.
- 유사한 개념은 유사한 표기법을 사용한다.


### 의미 있게 구분하라(불용어-noise word-를 쓰지 말자)  
- 연속된 숫자 덧붙이기 X (a1, a2, a3,... 등)
- 의미가 불분명한 불용어(noise word) 추가 X (Product, Info, data,... 등)
- 이름이 달라야 한다면 의미도 달라야 함
- 읽는 사람이 차이를 알도록 지어야 함

- 나의 경우, 페이지를 a10, a20, a30으로 나눌경우 REST API 함수명을
- getA10List(return값이 배열일때), getA11(리턴값이 객체일때)
- postA10, putA10 이렇게 나눔

- 예시  
 - `Name` VS `NameString`
 - `getActiveAccount()` VS `getActiveAccounts()` VS `getActiveAccountInfo()` (이들이 혼재할 경우 서로의 역할을 정확히 구분하기 어렵다.)
 - `money` VS `moneyAmount`
 - `message` VS `theMessage`

### 발음하기 쉬운 이름을 사용하라  

```java
// Bad
class DtaRcrd102 {
    private Date genymdhms;
    private Date modymdhms;
    private final String pszqint = "102";
    /* ... */
};
```
 
```java
// Good
class Customer {
    private Date generationTimestamp;
    private Date modificationTimestamp;
    private final String recordId = "102";
    /* ... */
};
```

### 검색하기 쉬운 이름을 사용하라  
- grep(유닉스를 위해 만들어진 텍스트 검색 기능을 가진 명령어) ex) MAX_CLASSES_PER_STUDENT
- 긴 이름 사용 사용
- 검색하기 쉬운 이름 사용
- 이름 길이 ∝ 범위 크기 (저자는 간단한 메서드에서 로컬 변수 정도만 한 문자 사용)

### 인코딩을 피하라
- 인코딩 : 코드화, 암호화. 사용자가 입력한 문자나 기호들을 컴퓨터가 이용할 수 있는 신호로 만드는 것. 
- 헝가리안 표기법
  - 변수 및 함수 앞에 데이터 타입을 명시하는 규칙
  - 변수, 함수, 클래스 이름이나 타입을 바꾸기 어려워지며 읽기도 어려워짐
- 맴버 변수 접두어 
  - 클래스와 함수는 접두어가 필요 없을 정도로 작아야 함.  
- 인터페이스와 구현
  - 인터페이스 클래스 이름과 구현 클래스 이름 중 하나를 인코딩 해야 한다면 구현 클래스 이름 택.
 
| Do / Don't | Interface class | Concrete(Implementation) class |
| ---------- | --------------- | ------------------------------ |
| Don't      | IShapeFactory   | ShapeFactory                   |
| Do         | ShapeFactory    | ShapeFactoryImp                |
| Do         | ShapeFactory    | CShapeFactory                  |

### 자신의 기억력을 자랑하지 마라   
- 똑똑한 프로그래머와 전문가 프로그래머 사이에서 나타나는 차이점 하나만 들자면, 전문가 프로그래머는 명료함이 최고라는 사실을 이해한다. 
### 클래스 이름  
- 클래스 이름과 객체 이름 ⇒ 명사나 명사구 
- ex) Customer, WikiPage, Account, AddressParser 등  
- Manager, Processor, Data, Info와 같은 단어는 피하자  
- 동사 X
### 메서드 이름  
- 메서드 이름 ⇒ 동사
- ex) postPayment, deletePage, save 등 
- 접근자, 변경자, 조건자는 get, set, is로 시작하자.  (추가: should, has 등도 가능)
- 생성자를 오버로드할 경우 정적 팩토리 메서드를 사용하고 해당 생성자를 private으로 선언한다.

```java 
// 첫번째 보다 두 번째 방법이 더 좋다.  
Complex fulcrumPoint = new Complex(23.0);  
Complex fulcrumPoint = Complex.FromRealNumber(23.0);  
```

### 기발한 이름은 피하라  
- 특정 문화에서만 사용되는 재미있는 이름보다 의도를 분명히 표현하는 이름을 사용하라  
  - HolyHandGrenade → DeleteItems  
  - whack() → kill()  

### 한 개념에 한 단어를 사용하라  
- 추상적인 개념 하나에 단어 하나를 사용하자.  
  - fetch, retrieve, get  
  - controller, manager, driver  
- 일관성 있는 어휘를 사용
- 똑같은 메서드는 같은 단어 사용

### 말장난을 하지 마라(위 내용에 이어)  
- 한 단어를 두 가지 목적으로 사용하지 말자. 아래와 같은 경우에는 ii를 append 혹은 insert로 바꾸는게 옳겠다.

```java
public static String add(String message, String messageToAppend)  
public List<Element> add(Element element)  
```

### 해법 영역(Solution Domain) 용어를 사용하자  
- 전산 용어, 알고리즘 이름, 패턴 이름, 수학 용어 등을 사용해도 괜찮다.
- 모든 이름을 문제 영역에서 가져오는 것은 현명하지 못함
- 기술 이름에는 기술 이름이 가장 적합
- `JobQueue`, `AccountVisitor(Visitor pattern)`

### 문제 영역(Problem Domain) 용어를 사용하자  
- 적절한 '프로그래머 용어' 가 없다면 문제 영역에서 이름을 가져온다. (유지보수시 의미파악가능)
- 문제 영역 개념과 관련이 깊은 코드라면 문제 영역에서 이름을 가져와야 한다. 
- 우수한 프로그래머와 설계자라면 해법 영역과 문제 영역을 구분할 줄 알아야 한다. 

### 의미 있는 맥락을 추가하라  
- 클래스, 함수, namespace등으로 감싸서 맥락(Context)을 표현하라  
- 그래도 불분명하다면 접두어를 사용하자.  

```java
// Bad
private void printGuessStatistics(char candidate, int count) {
    String number;
    String verb;
    String pluralModifier;
    if (count == 0) {  
        number = "no";  
        verb = "are";  
        pluralModifier = "s";  
    }  else if (count == 1) {
        number = "1";  
        verb = "is";  
        pluralModifier = "";  
    }  else {
        number = Integer.toString(count);  
        verb = "are";  
        pluralModifier = "s";  
    }
    String guessMessage = String.format("There %s %s %s%s", verb, number, candidate, pluralModifier );

    print(guessMessage);
}
```

```java
// Good
public class GuessStatisticsMessage {
    private String number;
    private String verb;
    private String pluralModifier;

    public String make(char candidate, int count) {
        createPluralDependentMessageParts(count);
        return String.format("There %s %s %s%s", verb, number, candidate, pluralModifier );
    }

    private void createPluralDependentMessageParts(int count) {
        if (count == 0) {
            thereAreNoLetters();
        } else if (count == 1) {
            thereIsOneLetter();
        } else {
            thereAreManyLetters(count);
        }
    }

    private void thereAreManyLetters(int count) {
        number = Integer.toString(count);
        verb = "are";
        pluralModifier = "s";
    }

    private void thereIsOneLetter() {
        number = "1";
        verb = "is";
        pluralModifier = "";
    }

    private void thereAreNoLetters() {
        number = "no";
        verb = "are";
        pluralModifier = "s";
    }
}
```

### 불필요한 맥락을 없애라  
- `Gas Station Delux` 이라는 어플리케이션을 작성한다고 해서 클래스 이름의 앞에 GSD를 붙이지는 말자. G를 입력하고 자동완성을 누를 경우 모든 클래스가 나타나는 등 효율적이지 못하다.  
위 처럼 접두어를 붙이는 것은 모듈의 재사용 관점에서도 좋지 못하다. 재사용하려면 이름을 바꿔야 한다.(eg, `GSDAccountAddress` 대신 `Address`라고만 해도 충분하다.)
- 의미가 분명한 경우에 한해 일반적으로 짧은 이름이 긴 이름이 좋다. 이름에 불필요한 맥학을 추가하지 않도록 한다.


> 두려워하지 말고 서로의 명명을 지적하고 고치자. 그렇게 하면 이름을 외우는 것에 시간을 빼앗기지 않고 "자연스럽게 읽히는 코드"를 짜는 데에 더 집중할 수 있다.
> 위에 코드대로 짜는건 연습많이 살길일듯 싶다.
















