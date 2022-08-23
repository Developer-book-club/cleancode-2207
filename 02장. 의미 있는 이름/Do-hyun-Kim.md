### 목차

- 의도를 분명히 밝혀라 🤔
- 그릇된 정보를 피하라 😫
- 의미 있게 구분하라 🥸
- 발음하기 쉬운 이름을 사용하라 👂
- 검색하기 쉬운 이름을 사용하라 🔍
- 인코딩을 피하라 💾
- 멤버 변수 접두어 💤
- 인터페이스 클래스와 구현 클래스 ☘️
- 자신의 기억력을 자랑하지 마라 🧠
- 메서드 이름 ✏️
- 기발한 이름은 피하라 🚫
- 한 개념에 한 단어를 사용하라 🙆‍♀️
- 말장난 하지 마라 😵‍💫
- 해법 영역에서 가져온 이름을 사용하라 🫶
- 문제 영역에서 가져온 이름을 사용하라 😮‍💨
- 의미 있는 맥락을 추가하라 😀
- 불필요한 맥락을 없애라 ❗️

### 의도를 분명히 밝혀라 🤔

---

- 좋은 이름을 지으려면 시간이 걸리지만 좋은 이름으로 절약하는 시간이 훨씬더 많다.
- 따로 주석이 필요 하다면 의도를 분명히 드러 내지 못했다는 말이다.

<span style="color: red; font-weight:bold; "> WRONG </span>

```swift
var d: Date = Date() //경과 시간(단위: 날짜)
```

<span style="color: green; font-weight:bold; ">Right </span>

```swift
var elapsedTimeInDays: Int?
var daysSinceCreation: Int?
var daysSinceModification: Int?
var fileAgeInDays: Int?
```

- 의도가 드러나는 이름을 사용하면 코드 이해와 변경이 쉬워진다. 



<span style="color: red; font-weight:bold; "> WRONG </span>

- list에 무엇이 들어 있는가?
- 값 4에는 무엇을 의미하는가?
- 함수를 반환하는 list1은 어떻게 사용하는가
- 단순히 list라는 이름만 부여했을뿐 어디에 사용하는지를 부여하지 않아 독자에게 많은 혼란을 일으킬 수 있다.

```swift
public func getThem(list: [Int]) -> [Int] {
    var list1:Array<Int> = []
    for a in list {
        if a == 4 {
            list1.append(a)
        }
    }
    return list1
}
```

<span style="color: green; font-weight:bold; ">Right </span>

- 변수와 인자값에 명확한 이름을 부여하며, a라는 값이 게임 보드의 칸을 의미하는 것을 알 수 있다.
- 이처럼 변수와 인자값에 의미있는 이름을 부여함으로써 코드 분석 시간을 단축 시킬 수 있다.

```swift
public func getFlaggedCells(gameBoard: [Int]) -> [Int] {
  var falggedCells: [Int] = []
  
  for cell in gameBoard {
    if cell == FLAGGED {
      falggedCells.append(cell)
    }
  }
  return falggedCells
}
```





### 그릇된 정보를 피하라 😫

---

- 프로그래머는 코드에 그릇된 단서를 남겨서는 안 된다. 그릇된 단서는 코드 의미를 흐린다.
- 실제 컨테이너가 List가 아니라면 List라 명명하지 않는다. 프로그래머에게 List라는 단어는 특수한 의미다.
- 실제로 List라 하더라도 컨테이너에 List 유형을 사용하지 않는 것이 바람직하다.

> accountGroup, bunchOfAccounts, Accounts, 등의 이름으로 List를 대처한다.

- 서로 흡사한 이름을 사용하지 않도록 주의하라!!

> XYZControllerForEfficientHandlingOfStrings 을 사용하고, 조금 떨어진 모듈에서 XYZControllerForEfficientStorageOfStrings라는 이름을 사용하면 차이를 알수 없며, 혼란을 일으킨다.



### 의미 있게 구분하라 🥸

---

- 컴파일러나 인터프리터만 통과하려는 생각으로 코드를 구현하는 프로그래머는 스스로 문제를 일으킨다.
- 이름이 달라야 한다면 의미 또한 달라져야 한다.
- Info나 Data는 a, an, the 와 마찬가지로 의미가 불분명한 불용어다. 불용어를 추가한 이름 역시 아무런 정보도 제공하지 못한다.
- moneyAmount는 money와 구분이 안되며, customerInfo는 customer와, accountData는 account와 theMessage는 message와 구분이 안된다. 이처럼 불용어는 아무런 정보를 제공하지 못한다.



### 발음하기 쉬운 이름을 사용하라 👂

---

- 사람들은 단어에 능숙하다. 우리 두뇌에서 상당 부분은 단어라는 개념만 전적으로 처리한다.

```swift
class DtaRcrd102 {
  private var genymdhms: Date
  private var modymdhms: Date
  private var pszqint: String = "102"
  
  init(genymdhms: Date, modymdhms: Date) {
    self.genymdhms = genymdhms
    self.modymdhms = modymdhms
  }
}

class Customer {
  private var generationTimestamp: Date
  private var modificationTimestamp: Date
  private var recordId: String = "102"
  
 init(generationTimestamp: Date, modificationTimestamp: Date) {
   self.generationTimestamp = generationTimestamp
   self.modificationTimestamp = modificationTimestamp
 } 
  
}
```



### 검색하기 쉬운 이름을 사용하라 🔍

---

- 긴 이름이 짧은 이름보다 좋다. 검색하기 쉬운 이름이 상수 보다 좋다. 이름 길이는 범위 크기에 비례한다.
- 이름을 의미있게 지으면 함수가 길어진다. 

```swift
for a in 0..<34 {
   s += (t[a]*4) / 5
}

private var realDaysPerIdealDay = 4
private let WORK_DAYS_PER_WEEK = 5
private var sum = 0

for i in 0..<NUMBER_OF_TASKS {
  var realTasksDays = taskEstimate[i] * realDaysPerIdealDay
  var realTasksWeeks = (realTasksDays / WORK_DAYS_PER_WEEK)
 	sum += realTaskWeeks
}

```



### 인코딩을 피하라 💾

---

- 개발자에게 인코딩은 불필요한 정신적 부담이나. 인코딩한 이름은 거의가 발음하기 어려우며 오타가 생기기도 쉽다.
- 유형이나 범위 정보 까지 인코딩에 넣으면 그만큼 이름을 해독하기 어려워진다.



### 멤버 변수 접두어 💤

---

- 이제는 멤버 변수에 m_ 이라는 접두어를 붙일 필요도 없다. 클래스와 함수는 접두어가 필요없을 정도로 작아야 마땅하다.
- 사람들은 접두어(접미어)를 무시하고 이름을 해독하는 방식을 재빨리 익히며, 코드를 읽을 수록 접두어는 관심 밖으로 밀려난다.
- 결국 접두어는 옛날에 작성한 구닥다리 코드라는 징표가 되버린다.

```swift
public class Part {
  private var m_dsc: String // 설명 문자열
  
  init(m_dsc: String) {
    self.m_dsc = m_dsc
  }
  
  public func setName(name: String) -> Void {
    self.m_dsc = name
  }
}
```





### 인터페이스 클래스와 구현 클래스 ☘️

---

- 인터페이스 클래스(Interface Class)의 이름은 접두어를 붙이지 않는 편이 좋다.

  > ShapeFactory, AccountFactory, DocumentFacotry

### 자신의 기억력을 자랑하지 마라 🧠

----

- 코드를 읽으면서 변수 이름을 자신이 아는 이름으로 변환해야 한다면 그 변수 이름은 바람직 하지 못하다는 것이다.
- 문자 하나만 사용하는 변수 이름은 문제가 있다 (단 루프 에서 반복 횟수를 세는 변수 i,j,k)는 괜찮다. 

> l은 안된다!! 숫자 1가 혼동이 올수 있다 또한 O 역시 0와 혼동이 올 수 있기 때문에 대도록이면 i,j,k 등을 사용하는 것이 좋다.

- 전문가 프로그래머는 명료함이 최고라는 사실을 이해하며 개발해 나아간다.

### 메서드 이름 ✏️

----

- 메서드 이름은 동사나 동사구가 적합하다.

> PostPayMent, deletePage, save 등이 좋은 예다

- 접근자(Accessor), 변경자(Mutator), 조건자(Predicate)는 javaBean 표준에 따라 값 앞에 get, set, is를 붙인다.



### 기발한 이름은 피하라 🚫

----

- 이름이 너무 기발하면 저자와 유머 감각이 비슷한 사람만, 그리고 농담을 기억하는 동안만, 이름을 기억한다.
- 특정 문화에서만 사용하는 농담은 피하는 편이 좋다. 의도를 분명하고 솔직하게 표현하라.



### 한 개념에 한 단어를 사용하라 🙆‍♀️

---

- 추상적인 개념 하나에 단어 하나를 선택해 이를 고수한다.

> 똑같은 메서드를 클래스 마다 Fetch, retrieve, get으로 제각각 부르면 혼란스럽다. 

- 이름이 다르면 당연히 클래스도 다르고 타입도 다르다 생각한다.
- 일관성 있는 어휘는 코드를 사용할 프로그래머가 반갑게 여길 선물이다.



### 말장난 하지 마라 😵‍💫

---

- 한 단어를 두가지 목적으로 사용하지 마라. 다른 개념에 같은 단어를 사용한다면 그것은 말장난에 불과 하다.
- 기존 메서드가 add라는 단어를 선택 했다 해서 '일관성'을 고려해 add라는 단어를 선택하는 것은 잘못된 것이다.

> 두 개를 더하거나 새로운 값을 만든 다고 하더라도 기존 add 메서드와 맥락이 다르기에 Insert와 Append라는 이름이 적합하다.

- 프로그래머는 코드를 최대한 이해하기 쉽게 짜야 한다.



### 해법 영역에서 가져온 이름을 사용하라 🫶

---

- 코드를 읽을 사람도 프로그래머이다. 전산 용어, 알고리즘 이름, 패턴 이름, 수학 용어 등을 사용해도 괜찮다.

  > AccountVisitor, JobQueue, AccountEntity



### 문제 영역에서 가져온 이름을 사용하라 😮‍💨

----

- 적절한 '프로그래밍 용어'가 없다면 문제 영역에서 이름을 가져온다.



### 의미 있는 맥락을 추가하라 😀

----

- 클래스, 함수, 이름 공간에 넣어 맥락을 부여한다. 모든 방법이 실패하면 접두어를 붙인다.

> firstName,lastName,street,houseNumber 에 접두어를 부여하여 맥락을 좀 더 분명하게 전달 한다.
>
> addrFirstName,addrLastName,addrStreet 처럼 접두어를 부여하여 맥락을 좀더 분명하게 전달 할 수 있다. 



### 불필요한 맥락을 없애라 ❗️

----

- 일반적으로는 짧은 이름이 긴 이름보다 좋다. 단, 으미가 분명한 경우에 한해서다. 이름에 불필요한 맥락을 추가하지 않도록 주의 한다.

> accountAddress와 customerAddress는 Address 클래스 인스턴스로는 좋은 이름이나 클래스 이름으로는 적합하지 못하다. Address는 클래스 이름으로 적합하다.
>
> - accountAddress 인스턴스를 찾으려면 address에 관한 것들이 함께 검색 되기에 accountId 가 적합하다.

