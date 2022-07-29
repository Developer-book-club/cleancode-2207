# 1장 깨끗한 코드

### 코드는 항상 존재하리라

- 코드는 요구사항을 표현하는 언어라는 사실을 명심한다.
    - 요구사항에 가까운 언어를 맞추고 정형 구조를 뽑아내는 도구를 만들 수 있지만, 
    - 어느순간에는 정밀한 표현이 필요하다.

### 나쁜 코드

- 르블랑의 법칙(나중은 절대 돌아오지 않는다.)
- 회사가 망한 원인은 나쁜 코드 탓.

### 나쁜코드로 큰 코 다치는 비용

##### 나쁜코드가 생기는 과정

##### 태도

##### 원초적 난제

- 나쁜코드가 업무 속도를 늦춘다.
- 기한을 맞추는 유일한 방법은 “언제나 클린코드 유지하는 습관”

##### [펌글] 본 실력없는 팀장

> 혼자 서버구축, UI, UX, 웹 앱 혼자서 해내고 100억 투자를 받음
> 혼자 많은걸 구현하다보니 테스트코드 X, 유지보수 X가 되어버림
> 팀원들이 보충되고 한 파트씩 업무를 가져가는데 위 내용처럼 나쁜 코드가 되어버림


### 깨끗한 코드

- 잘그린 그림을 구분하는 능력 = 그림을 잘 그리는 능력
- 빈 캔퍼스에 우아한 작품을 만들어내는 감각. 코드 감각을 길러야 된다.

<!-- -->

- **Bjarne Stroustrup**
    - 클린코드는 한 가지에 집중한다. **클린코드는 한 가지 일을 잘 한다.**
- **Grady Booch**
    - 클린코드는 하나의 잘 쓰여진 산문처럼 읽혀야 한다. 소설의 기승전결처럼 문제를 제시하고 명쾌한 해답을 제시해야 한다.
    - 명백한 추상 : 코드는 추측 대신 실제를 중시, 필요한 것만 포함하며 독자로 하여금 결단을 내렸다고 생각하게 해야 한다.
- **“Big” Dave Thomas**
    - 다른 이가 수정하기 쉬워야 한다.
    - 테스트를 해야 한다.
    - 코드는 간결할 수록, 세련될 수록 좋다.
- **Michael Feathers**
    - 코드에 주의, 관심을 가지고 작성해라
- **Ron Jeffries**
    - 모든 테스트 통과
    - 중복제거
    - 클래스/메서드는 최소화
    - 메서드의 이름으로 코드의 표현력 강화
- **Ward Cunningham**
    - "코드가 그 문제를 풀기 위한 언어처럼 보인다면 아름다운 코드라 불러도 되겠다."
    - 그 문제를 풀기위한 언어처럼 보이는 코드를 작성한다.

### 보이스카우트 규칙
- 시간이 지나도 언제나 깨끗하게 유지해야 한다.(지속적인 개선)
- 캠프장은 처음 왔을 때보다 더 깨끗하게 해놓고 떠난다.

### 프리퀄의 원칙
- SRP(Single Responsibility Principle): 클래스에는 단 한 가지 변경 이유만 존재해야 한다.
- OCP(Open Closed Principle): 클래스는 확장에 열려있어야 하며 변경에 닫혀있어야 한다.
- LSP(Liskou Subsitution Principle): 상속 받은 클래스는 기초 클래스를 대체할 수 있어야 한다.
- DIP(Depending Inversion Principle): 추상화에 의존 O, 구체화에 의존 X
- ISP(Interface Seregation Principle): 클라이언트에 밀접하게 작게 쪼개진 인터페이스 유지

<!-- -->
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



# 3장 함수

### 작게 만들어라!

#### 함수를 만들 때 최대한 ‘작게!’ 만들어라.

- 함수를 작게, 나눠 작성할수록 가독성이 올라간다.
- 설계시 SOLID의 원칙을 따르는것도 좋을 것 이다.
```java
public static String renderPageWithSetupsAndTeardowns( PageData pageData, boolean isSuite) throws Exception {
	boolean isTestPage = pageData.hasAttribute("Test"); 
	if (isTestPage) {
		WikiPage testPage = pageData.getWikiPage(); 
		StringBuffer newPageContent = new StringBuffer(); 
		includeSetupPages(testPage, newPageContent, isSuite); 
		newPageContent.append(pageData.getContent()); 
		includeTeardownPages(testPage, newPageContent, isSuite); 
		pageData.setContent(newPageContent.toString());
	}
	return pageData.getHtml(); 
}
```
 위 코드도 길다. 되도록 한 함수당 3~5줄 이내로 줄이는 것을 권장한다
 ```java
public static String renderPageWithSetupsAndTeardowns( PageData pageData, boolean isSuite) throws Exception { 
	if (isTestPage(pageData)) 
		includeSetupAndTeardownPages(pageData, isSuite); 
	return pageData.getHtml();
}
```

#### 블록과 들여쓰기  
- 중첩구조(if/else, while문 등)에 들어가는 블록은 한 줄이어야 한다. 각 함수 별 들여쓰기 수준이 2단을 넘어서지 않고,  각 함수가 명백하다면 함수는 더욱 읽고 이해하기 쉬워진다.
- 코드 블럭이 한 화면을 넘어가면 함수로 추출하는 것이 적절하다 생각한다. (가독성이 어려워짐)


### 한 가지만 해라  

- SOLID의 단일 책임의 원칙

#### 함수 내 섹션

- 함수를 여러 섹션으로 나눌 수 있다면 그 함수는 여러작업을 하는 셈이다.


### 함수 당 추상화 수준은 하나로

- 코드라인의 표현이 개념을 나타내는 것인지, 구현을 나타내는 것인지 구별하기 힘들다.
- 함수가 ‘한가지’ 작업만 하려면 함수 내 모든 문장의 추상화 수준이 동일해야 된다.  
- 만약 한 함수 내에 추상화 수준이 섞이게 된다면 읽는 사람이 헷갈린다.

#### 위에서 아래로 코드 읽기:내려가기 규칙  

- 코드는 위에서 아래로 이야기처럼 읽혀야 좋다.  
- 함수 추상화 부분이 한번에 한단계씩 낮아지는 것이 가장 이상적이다.(내려가기 규칙)


### Switch문

- 당연하게도 switch문을 사용하는 함수에선 하나에 하나만의 동작을 수행하기 힘들다. 특정 변수의 값에 대해 조건문의 개수가 발산하는 상황에서 사용되기 때문이다. 책에서는 switch문을 통해 동작을 결정하는 코드를 추상 팩토리 패턴으로 숨기는 방법을 제안한다.

```java
public Money calculatePay(Employee e) throws InvalidEmployeeType {
	switch (e.type) { 
		case COMMISSIONED:
			return calculateCommissionedPay(e); 
		case HOURLY:
			return calculateHourlyPay(e); 
		case SALARIED:
			return calculateSalariedPay(e); 
		default:
			throw new InvalidEmployeeType(e.type); 
	}
}
```

```java
public abstract class Employee {
	public abstract boolean isPayday();
	public abstract Money calculatePay();
	public abstract void deliverPay(Money pay);
}
-----------------
public interface EmployeeFactory {
	public Employee makeEmployee(EmployeeRecord r) throws InvalidEmployeeType; 
}
-----------------
public class EmployeeFactoryImpl implements EmployeeFactory {
	public Employee makeEmployee(EmployeeRecord r) throws InvalidEmployeeType {
		switch (r.type) {
			case COMMISSIONED:
				return new CommissionedEmployee(r) ;
			case HOURLY:
				return new HourlyEmployee(r);
			case SALARIED:
				return new SalariedEmploye(r);
			default:
				throw new InvalidEmployeeType(r.type);
		} 
	}
}
```
하지만 switch문은 불가피하게 써야될 상황이 많으므로, 상황에 따라서는 사용 할 수도 있다.


### 서술적인 이름을 사용하라!  

- 함수가 무슨일을 하는가? 를 표현해야 한다.
- 작은 함수는 그 기능이 명확하므로 이름을 붙이기가 더 쉬우며, 일관성 있는 서술형 이름을 사용한다면 코드를 순차적으로 이해하기도 쉬워진다.
- 서술적인 이름은 테스트코드에서 좀 더 사용하는듯 하다.

### 함수 인수  
- 함수에서 이상적인 인수 개수는 0개(무항).
- 인수는 코드 이해에 방해가 되는 요소이므로 최선은 0개이고, 차선은 1개뿐인 경우이다.
- 출력인수(함수의 반환 값이 아닌 입력 인수로 결과를 받는 경우)는 이해하기 어려우므로 왠만하면 쓰지 않는 것이 좋겠다.

#### 많이 쓰는 단항 형식  

> 함수의 인자가 한개 들어가는 대표적인 예시로는 boolean fileExists("MyFile")처럼 인자에 대한 질문을 하는  
> 경우, InputStream fileOpen("MyFile")처럼 인자를 다른 형태로 변환하는 경우가 있다고 한다. 각 용도를 잘 
> 구분하기 위해 함수 이름을 잘 지어야 한다는데, 보통 질문에 해당하는 함수는 isFileExists() 같은 식으로 앞에 
> is, has 같이 boolean한 결과라는 것을 예측할 수 있는 이름으로 구분하면 충분할 것 같다.

#### 플래그 인수

- boolean 형의 플래그 인자를 사용하는 함수는 한 가지 일만 하지 않는다는 뜻 이므로 좋지 않다고 한다.

#### 이항 함수, 삼항 함수, 인수 객체

- 핵심은 인자의 개수가 많아질수록 이해하기 어렵다.
- 최대한 줄이고, 인자들의 관계가 하나의 객체로 묶을 수 있다면, 객체료 포현해 인자 개수를 줄이는 것이 좋다.
- circle(x, y) , circle(point) 

#### 동사와 키워드

- 단항 함수는 함수와 인수가 동사/명사 쌍을 이뤄야한다.  
`writeField(name);`  
- 함수이름에 키워드(인수 이름)을 추가하면 인수 순서를 기억할 필요가 없어진다.  
`assertExpectedEqualsActual(expected, actual);`  

### 부수 효과를 일으키지 마라!

- 부수효과는 거짓말이다. 함수에서 한가지를 하겠다고 약속하고는 남몰래 다른 짓을 하는 것이므로, 한 함수에서는 딱 한가지만 수행할 것!
- 부수 효과를 완전히 배제하면서 프로그래밍을 하는 것은 거의 불가능에 가깝겠지만, 일반적인 함수의 경우 순수 함수(pure function)로 작성하여 부수 효과를 없앰으로서 의도치 않은 버그를 없앤다는 것이 핵심이다.

아래 코드에서 `Session.initialize();`는 함수명과는 맞지 않는 부수효과이다.
```java
public class UserValidator {
	private Cryptographer cryptographer;
	public boolean checkPassword(String userName, String password) { 
		User user = UserGateway.findByName(userName);
		if (user != User.NULL) {
			String codedPhrase = user.getPhraseEncodedByPassword(); 
			String phrase = cryptographer.decrypt(codedPhrase, password); 
			if ("Valid Password".equals(phrase)) {
				Session.initialize();
				return true; 
			}
		}
		return false; 
	}
}
```
- checkPassword() 함수의 결과를 바탕으로 Session.initalize()가 호출되어야 한다.

#### 출력인수  

- 일반적으로 출력 인수는 피해야 한다.   
- 함수에서 상태를 변경해야 한다면 함수가 속한 객체 상태를 변경하는 방식을 택하라.

### 명령과 조회를 분리하라

함수는 뭔가 객체 상태를 변경하거나, 객체 정보를 반환하거나 둘 중 하나다. 둘 다 수행해서는 안 된다.  
`public boolean set(String attribute, String value);`같은 경우에는 속성 값 설정 성공 시 true를 반환하므로 괴상한 코드가 작성된다.  
`if(set(“username”, “unclebob”))...` 그러므로 명령과 조회를 분리해 혼란을 주지 않도록 한다.  

## 오류코드보다 예외를 사용하라!

try/catch를 사용하면 오류 처리 코드가 원래 코드에서 분리되므로 코드가 깔끔해 진다.

#### Try/Catch 블록 뽑아내기

- 들여쓰기를 필요이상 하게되고, 이 블록 들여쓰기 상에 생기는 더러움을 최소화 하기 위해 try문에서 시도하는 부분도 함수로 추출해야 할 필요가 있다.

```java
if (deletePage(page) == E_OK) {
	if (registry.deleteReference(page.name) == E_OK) {
		if (configKeys.deleteKey(page.name.makeKey()) == E_OK) {
			logger.log("page deleted");
		} else {
			logger.log("configKey not deleted");
		}
	} else {
		logger.log("deleteReference from registry failed"); 
	} 
} else {
	logger.log("delete failed"); return E_ERROR;
}
```

정상 작동과 오류 처리 동작을 뒤섞는 추한 구조이므로 if/else와 마찬가지로 블록을 별도 함수로 뽑아내는 편이 좋다.

```java
public void delete(Page page) {
	try {
		deletePageAndAllReferences(page);
  	} catch (Exception e) {
  		logError(e);
  	}
}

private void deletePageAndAllReferences(Page page) throws Exception { 
	deletePage(page);
	registry.deleteReference(page.name); 
	configKeys.deleteKey(page.name.makeKey());
}

private void logError(Exception e) { 
	logger.log(e.getMessage());
}
```

오류 처리도 한가지 작업이다.

Error.java 의존성 자석

```java
public enum Error { 
	OK,
	INVALID,
	NO_SUCH,
	LOCKED,
	OUT_OF_RESOURCES, 	
	WAITING_FOR_EVENT;
}
```
- 예외를 처리하지 않고, 오류 코드를 사용하다 보면 오류의 종류에 따라 오류 코드를 선언해야 하고, 결국 오류 코드들을 선언한 (선언할) Error.java 파일에 점점 의존성이 추가될 것이라는 것이다. 간접적으로 SOLID의 단일 책임 원칙을 위반할 가능성이 높다 해석할 수 있다.

### 반복하지 마라!

- 중복은 모든 소프트웨어에서 모든 악의 근원이므로 늘 중복을 없애도록 노력해야한다.
- vue에서 mixins (중복코드 처리기능)

### 구조적 프로그래밍  

> 다익스트라의 구조적 프로그래밍의 원칙을 따르자면 모든 함수와 함수 내 모든 블록에 입구와 출구가 하나여야 된다. 즉, 
> 함수는 return문이 하나여야 되며, **루프 안에서 break나 continue를 사용해선 안된며 goto는 절대로, 절대로 사용하지 말자.** 
> 함수가 클 경우에만 상당 이익을 제공하므로, 함수를 작게 만든다면 오히려 여러차례 사용하는 것이 함수의 의도를 표현하기 쉬워진다.

> 그런데 구조적 프로그래밍의 목표와 규율은 공감하지만 함수가 작다면 위 규칙은 별 이익을 제공하지 못한다. 함수가 
> 아주 클 때만 상당한 이익을 제공한다. 그러므로 함수를 작게 만든다면 간혹 return, break, continue를 사용해도 
> 괜찮다. 오히려 때로는 단일 입/출구 규칙보다 의도를 표현하기 쉬워진다.

- Java에서 가비지 컬렉터가 지원되므로 자원 할당에 대한 해제작업을 매번 해줄 필요는 없다고 본다.
- 그 외 자원관리인 작업 쓰레드를 종료시키거나, 파일 입출력을 위해 열었던 파일들을 닫는 작업등 들이 명시적으로 수행되어야 할 것 같다.


### 함수를 어떻게 짜죠?  

- 처음에는 길고 복잡하고, 들여쓰기 단계나 중복된 루프도 많다. 인수목록도 길지만, 이 코드들을 빠짐없이 
- 테스트하는 단위 테스트 케이스도 만들고, 
- 코드를 다듬고, 함수를 만들고, 이름을 바꾸고, 중복을 제거한다. 처음부터 탁 짜지지는 않는다.












