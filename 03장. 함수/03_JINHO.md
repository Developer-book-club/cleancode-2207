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












