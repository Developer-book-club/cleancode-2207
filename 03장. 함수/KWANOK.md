# 함수

## 한 가지만 해야 합니다.

> 함수는 한 가지를 해야 한다. 그 한 가지를 잘 해야 한다. 그 한 가지만을 해야 한다.

### 함수의 추상화 수준을 최대한 낮추면 좋습니다.

아래 코드의 추상화 수준은 하나입니다.

함수는 TO 문단으로 기술할 수 있습니다.

to render page with setups and teardowns

페이지가 테스트페이지인지 확인할 호 테스트페이지라면 설정 페이지와 해제 페이지를 넣는다

테스트페이지든 아니든 페이지를 HTML로 렌더링한다.

```go
func renderPageWithSetupsAndTeardowns(pageData PageData, isSuite bool) string {
	if isTestPage(pageData) {
		includeSetupAndTeardownPages(pageData, isSuite)
	}
	
	return pageData.getHtml()
}
```

위 코드처럼 추상화 수준이 하나인 단계만 수행한다면 그 함수는 한 가지 작업만 하는 겁니다.

단순히 다른 표현이 아니라 의미있는 이름으로 다른 함수를 추출할 수 있다면,

그 함수는 여러 작업을 한다고 봐도 됩니다.

### 코드는 위에서 아래로 이야기처럼 읽혀야 좋습니다.

어떤 함수 다음에 오는 함수는 추상화가 한 단계 낮은 함수가 옵니다.

위에서 아래로 프로그램을 읽는다면 추상화 수준이 한 단계씩 낮아지는 겁니다.

but 어렵다.

## Switch 문

아래는 직원 유형에 따라 다른 값을 계산해 반환하는 함수입니다.

```go
const (
	COMMISSIONED = "T1"
	HOURLY       = "T2"
	SALARIED     = "T3"
)

type employee struct {
	Type string
}

func calculatePay(e employee) *Money {
	switch e.Type {
	case COMMISSIONED:
		return calculateCommissionedPay(e)
	case HOURLY:
		return calculateHourlyPay(e)
	case SALARIED:
		return calculateSalariedPay(e)
	default:
		return nil
	}
}
```

위 코드의 문제는

- 깁니다.
- 직원 유형이 추가되면 길어집니다.
- 하나의 작업만 수행하지 않습니다.
- SRP(Single Responsibility Principle) 를 위반합니다.
- OCP(Open Closed Principle) 를 위반합니다.

그래서 아래처럼 스위치문을 추상 팩토리에 숨기는 방법이 있습니다.

그래서 `EmployeeRecord` 를 통해 `makeEmployee` 메서드를 호출하면 `Employee` 인터페이스에 맞게 구현된 적절한 Employee 를 반환하게 됩니다.

```go
type EmployeeFactory interface {
	makeEmployee() Employee
}

type Employee interface {
	isPayday() bool
	calculatePay() Money
	deliverPay(money Money)
}

type EmployeeRecord struct {
	Type string
}

func (e *EmployeeRecord) makeEmployee() Employee {
	switch e.Type {
	case COMMISSIONED:
		return &CommissionEmployee{}
	case HOURLY:
		return &HourlyEmployee{}
	case SALARIED:
		return &SalariedEmployee{}
	default:
		return nil
	}
}
```

## 서술적인 이름을 사용해주세요.

> 한 가지만 하는 작은 함수에 좋은 이름을 붙인다면 이런 원칙을 달성함에 있어 이미 절반은 성공했다.

함수가 작고 단순할수록 서술적인 이름을 고르기 쉬워집니다.

함수 이름은

- 길어도 됩니다.
- 서술적인 이름이 길고 서술적인 주석보다 좋습니다.
- 이름을 정할 땐 여러 단어가 쉽게 읽히는 명명법을 사용합니다.
- 여러 단어를 사용해 함수 기능을 잘 표현하는 이름을 선택합니다.

서술적인 이름을 사용하면, 개발자 머릿속에서도 설계가 뚜렷해져 코드를 개선하기 쉬워집니다.

## 함수 인수

> 함수에서 이상적인 인수 개수는 0개(무항)입니다.

다음은 1개, 다음은 2개입니다.

3개 이상부터는 피하는 게 좋습니다.

### 코드를 읽는 사람 관점

`includeSetupPageInto(newPageContent)` 보다 `includeSetupPage()` 가 더 이해하기 쉽습니다.

### 테스트 관점

갖가지 인수 조합으로 함수를 검증하는 테스트를 짠다고 가정했을 때

인수가 많아지면 많아질수록 검증해야하는 복잡도가 급상승합니다.

### 인수 객체

인수가 2-3개 필요할 경우 일부를 독자적인 클래스로 넘기는 게 더 좋습니다.

```go
func makeCircle(x float64, y float64, double radius)
func makeCircle(center Point, double radius)
```

### 동사와 키워드

함수의 의도나 인수의 순서와 의도를 잘 표현하려면 좋은 이름이 필요합니다.

당항 함수는 인수가 동사 / 명사 쌍을 이뤄야 합니다.

ex) `write(name)`, `writeField(name)`

그리고 함수 이름에 키워드를 추가하는 방법도 있습니다.

아래 두 예시중 아래처럼 바꾸면 인수의 순서를 기억할 필요가 없어집니다.

```go
func assertEquals(expected, actual)
```

```go
func assertExpectedEqualsActual(expected, actual)
```

## 부수 효과를 일으키면 안됩니다.

부수효과를 일으킨다는 것은

- 예상치 못하게 클래스 변수를 수정하고,
- 함수로 넘어온 임수나 시스템 전역 변수를 수정합니다.

아래 예시를 보면 간단히 패스워드를 검증하는 것 같지만, 세션을 새로고침하는 걸 확인할 수 있습니다.

함수의 이름은 `checkPassword` 이지만, 세션을 초기화한다는 정보는 없습니다.

그래서 함수 이름만 보고 호출한다면, 사용자의 세션을 지워버릴 수도 있는 위험이 있습니다.

```go
func checkPassword(userName string, password string) bool {
	var userGateway UserGateway
	user := userGateway.FindByName(userName)
	if user != nil {
		codedPharase := user.getPhraseEncodedByPassword()
		phrase := bcrypt.CompareHashAndPassword([]byte(codedPharase), []byte(password))
		if phrase == nil {
			Session.initialize()
		}
		return true
	}

	return false
}
```

그리고 이 함수가 기존에 세션을 초기화해도 되는 상황에만 구현돼있을 경우,

다음에 세션을 초기화해선 안 되는 상황에 이 함수를 호출했다간 큰일납니다.

이런 부수 효과로 숨겨진 함수는 큰 혼란을 가져올 수 있습니다.

### 출력 인수

잘 이해가 가지 않지만..

아래 예시보단

```go
func main() {
    r := Report
    appendFooter(&r)
}
```

이 예시가 낫다는 것 같습니다.

```go
func main() {
    r := Report
    r.appendFooter()
}
```

## 명령과 조회를 분리하면 좋습니다.

함수는 뭔사를 수행하거나 뭔가에 답하거나 둘 중 하나만 해야 합니다.

둘 다 하면 안됩니다.

객체 상태를 변경하거나, 객체 정보를 반환하거나 둘 중 하나만 해야합니다.

아래 함수는 이름이 `attribute` 인 속성을 찾아 값을 `value` 로 설정한 뒤 성공하면 true, 실패하면 false를 반환합니다.

```go
func set(attribute string, value string) bool {
    ...
}
```

그러면 다음과 같은 코드가 나옵니다

```go
if set("username", "kwanok") ...
```

여기서 일단 `set` 함수가 왜 if문 안에 들어오는지 의문이 듭니다.

그래서 차라리 아래와 같이 `set` 함수는 설정 기능에만 집중하고 오류만 반환하는 쪽으로 가는 게 좋다고 생각합니다.

```go
func set(attr string, val string) error {
	var db *sql.DB
	_, err := db.Query("INSERT ...")
	return err
}

func main() {
	err := set("username", "kwanok")
	if err != nil {
		...
	}
}
```

## 반복하지 맙시다.

중복은 문제입니다.

왜냐면 수정 사항이 생기면 중복된 만큼 수정을 해야하기 때문입니다.

그리고 중복이 있는 경우 하나하나 꼼꼼하게 수정해야하기 때문에 그걸 수정하고 테스트하는 일은 정말 끔찍합니다.

## 마무리

함수는 처음엔 길고 복잡합니다.

이걸 다듬고, 분리하고, 중복을 제거하면 좋은 함수가 나온다고 합니다.



