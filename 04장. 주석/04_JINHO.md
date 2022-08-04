# 4장. 주석

## 주석은 나쁜 코드를 보완하지 못한다 ##

- 보통 주석을 작성하는 이유는 코드가 더럽기 때문.
- 차라리 코드를 정리하자.

## 코드를 의도로 표현하라! ##

```java
if ((employee.flags & HOURLY_FLAG) && (employee.age > 65)
```

```java
if (employee.isEligibleForFullBenefits())
```

- 직관적인 함수 명으로 주석을 대체할 수 있다.
- 함수로 구분해야 할 만큼 자주 쓰이면 더 좋을듯 하다. 


## 좋은 주석 ##

#### 법적인 주석: 각 소스 파일 첫머리에 들어가는 저작권 정보와 소유권 정보 등

- 오픈 소스의 경우, 라이선스에 관한 경고문이나, 회사의 법적 소유권을 증명하기 위해 필요하다.
- 개발자에 입장에서는 유용하진 않으나 사업적으로 필요한 주석이다.

#### 정보를 제공하는 주석 

```java
// 테스트 중인 Responder 인스턴스를 반환
protected abstract Responder responderInstance();
```

```java
// kk:mm:ss EEE, MMM dd, yyyy 형식이다.
Pattern timeMatcher = Pattern.compile("\\d*:\\d*\\d* \\w*, \\w*, \\d*, \\d*");
```

- 기본적인 정보를 제공하기 위해 작성된 주석이다.
- 정규 표현식의 경우, 아래의 의미를 명료하게 밝히는 주석이랑 겹치는 부분이 있어보인다.

#### 의도를 설명하는 주석

```java
// 스레드를 대량 생성하는 방법으로 어떻게든 경쟁 조건을 만들려 시도한다. 
for (int i = 0; i > 2500; i++) {
    WidgetBuilderThread widgetBuilderThread = 
        new WidgetBuilderThread(widgetBuilder, text, parent, failFlag);
    Thread thread = new Thread(widgetBuilderThread);
    thread.start();
}
```
- 직관적으로 이해되지 않는 코드 흐름에 다는 것은 좋다고 생각한다.
- TODO 대체가능해보임


#### 결과를 경고하는 주석    

```java
// 여유 시간이 충분하지 않다면 실행하지 마십시오.
public void _testWithReallyBigFile() {

}
```

- TODO 대체가능해보임

#### TODO 주석    

```java
// TODO-MdM 현재 필요하지 않다.
// 체크아웃 모델을 도입하면 함수가 필요 없다.
protected VersionInfo makeVersion() throws Exception {
    return null;
}
```
- 코드의 부족한 부분, 필요한 정보를 정확하게 알려준다.
- 태그에 따라 주석의 역활이 명확하게 나뉜다.
- TODO, FIXME는 공통적으로 활용된다.

> TODO : 아직 해당 기능이 작성되지 않았지만, 코드 틀을 작성하느라 임시로 생성한 함수를 표현할 때
> FIXME : 임시 방편으로 작성된 코드거나(하드 코딩), 수정이 필요하다고 논의된 코드에 메모를 남길 때
> HACK : 문제를 회피하는 기법에 대한 메모, 직관적으로 이해하기 힘들지만 효율적인 코드에 대한 간단한 설명에 사용
> XXX : 경고를 표현할 때 사용, WARN 등으로도 표현


#### 중요성을 강조하는 주석    

```java
String listItemContent = match.group(3).trim();
// 여기서 trim은 정말 중요하다. trim 함수는 문자열에서 시작 공백을 제거한다.
// 문자열에 시작 공백이 있으면 다른 문자열로 인식되기 때문이다. 
new ListItemWidget(this, listItemContent, this.level + 1);
return buildList(text.substring(match.end()));
```

#### 공개 API에서 Javadocs

- 잘 작성된 API문서는 큰 도움이 된다. 

## 나쁜 주석 ##

#### 주절거리는 주석    

- 주석을 써야 한다면, 시간을 들여서 필요하고 중요한 주석만 작성하는 것이 좋다.

#### 같은 이야기를 중복하는 주석    

- 같은 맥락의 주석이 반복되면 시간만 오래걸린다.

#### 오해할 여지가 있는 주석

- 주석으로 인해 코드를 이해를 방해하는 주석은 불필요하다.

#### 의무적으로 다는 주석    

```java
/**
 *
 * @param title CD 제목
 * @param author CD 저자
 * @param tracks CD 트랙 숫자
 * @param durationInMinutes CD 길이(단위: 분)
 */
public void addCD(String title, String author, int tracks, int durationInMinutes) {
    CD cd = new CD();
    cd.title = title;
    cd.author = author;
    cd.tracks = tracks;
    cd.duration = durationInMinutes;
    cdList.add(cd);
}
```

- 코드 작성할 때 주석에 대한 피로도를 높이는 주범이다. 모든 함수/클래스에 문서화 작업을 강제하면 불필요한 주석을 작성하게 만든다.

#### 이력을 기록하는 주석    

- 형상 관리 시스템을 잘 사용하면 다 알 수 있다. 

#### 있으나 마나 한 주석

```java
/*
 * 기본 생성자
 */
protected AnnualDateRule() {

}
```

#### 무서운 잡음

```java
return 0; // 0을 반환한다.
return 0; // 정상 종료

return EXIT_SUCCESS;
```

- EXIT_SUCCESS같이 리팩토링하여 주석을 제거 할 수 있다.

#### 함수나 변수로 표현할 수 있다면 주석을 달지 마라    

```java
// 전역 목록 <smodule>에 속하는 모듈이 우리가 속한 하위 시스템에 의존하는가?
if (module.getDependSubsystems().contains(subSysMod.getSubSystem()))
```

```java
ArrayList moduleDependencies = smodule.getDependSubSystems();
String ourSubSystem = subSysMod.getSubSystem();
if (moduleDependees.contains(ourSubSystem))
```

- 함수를 통해 무슨 일을 하는지 설명한다.
- 변수를 통해 어떤 의도로 값을 저장하는지 나타낸다.

#### 위치를 표시하는 주석    

- 소스 코드의 길이가 길어질 때, 역할 별로 코드를 구분하기 위해 구분자 같은 주석을 사용하기도 한다.
- 반드시 필요할 때는 드물게 사용하는 것이 좋다.

#### 닫는 괄호에 다는 주석

- 들여쓰기를 줄이면 어느 부분이 닫는 괄호인지 알 수 있다.

#### 공로를 돌리거나 저자를 표시하는 주석

````java
/* 릭이 추가함 */
````

- 형상 관리를 사용하면 위 주석은 필요가 없다.

#### 주석으로 처리한 코드    

- 1960년대 즈음에는 주석으로 처리한 코드가 유용했었지만 우리는 우수한 소스 코드 관리 시스템을 사용하기 때문에 우리를 대신에 코드를 기억해준다. 그냥 삭제하라. 잃어버릴 염려는 없다. 약속한다. 

#### 전역 정보

```java
/**
 * 적합성 테스트가 동작하는 포트: 기본값은 <b>8082</b>.
 *
 * @param fitnessePort
 */
public void setFitnessePort(int fitnessePort) {
    this.fitnewssePort = fitnessePort;
}
```

- 주석을 작성한다면, 해당 주석 근처의 코드만 서술해야한다. 너무 멀리 있는 내용을 서술하면 주석 및 코드를 이해하기 위해 계속 소스 코드를 오고 가며 읽어야 하기 때문에 불편하기 때문이다.

#### 함수 헤더

- 짧은 함수는 설명이 많이 필요하지 않다. 함수 이름만 잘 선택해도 주석이 불필요하다.

#### 비공개 코드에서 Javadocs    

- 공개 API는 Javadocs가 유용하지만 공개하지 않을 코드라면 Javadocs는 쓸모가 없다. 코드만 보기싫고 산만해질 뿐이다. 

#### 결론

- 위에서 말한 대로 나쁜 주석으로 인해 코드의 혼동이 생길수는 있으나, 부득이하게 사용할 경우도 분명히 존재할것 같다.
