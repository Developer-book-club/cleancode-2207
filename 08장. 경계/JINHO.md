## 경계

### 외부 코드 사용하기
### 경계 탐험하고 공부하기
### log4j 공부하기
### 학습 테스트는 꽁자 이상이다
### 아직 존재하지 않는 코드 사용하기
### 깨끗한 경계

## 외부 코드 사용하기

- 패키지 제공자는 적용성을 최대한 넓히려 하고, 사용자는 자신의 요구에 집중하는 인터페이스를 원한다.
- 이런 차이로 인해 시스템의 경계에서 문제가 생길 소지가 많다.

```java
Map sensors = new HashMap();

Sensor s = (Sensor) sensors.get(sensorId);

Map<Sensor> sensors = new HashMap<Sensor>();

Sensor s = sensors.get(sensorId);
```

```java
public class Sensors {
    private Map sensors = new HashMap();
    
    public Sensor getById(String id) {
        return (Sensor)sensors.get(id);
    }
}
```
- java에서 Map의 경우를 예시로 드는데 제네릭 문법을 써서 코드를 간략하게 할수 있다.
- 두번째 예시처럼 Wrapping를 써서 Sensors 클래스 안에 넣으면 코드를 이해하기 쉽고 오용할 일이 없다.
- Sensors 클래스를 설계 규치고가 비즈니스 규칙을 따르도록 강제할 수 있다.


## 경계를 탐험하고 공부하기
- 외부패키지 테스트가 우리 책임은 아니다. 하지만 우리 자신을 위해 우리가 사용할
코드를 테스트하는 편이 바람직하다.
-  우리쪽 코드를 작성해 외부 코드를 호출하는 대신 먼저 간단한 테스트 케이스를 작성해 외부 코드를 익힌다. 이것을 '학습 테스트'라 부른다.


## log4j 공부하기 ##
```java
    @Test
    public void testLogCreate() {
        Logger logger = Logger.getLogger("MyLogger");
        logger.info("hello");
    }

    @Test
    public void testLogAddAppender() {
        Logger logger = Logger.getLogger("MyLogger");
        ConsoleAppender appender = new ConsoleAppender();
        logger.addAppender(appender);
        logger.info("hello");
    }

    @Test
    public void testLogAddAppender() {
        Logger logger = Logger.getLogger("MyLogger");
        logger.removeAllAppenders();
        logger.addAppender(new ConsoleAppender(
            new PatternLayout("%p %t %m%n"),
            ConsoleAppender.SYSTEM_OUT));
        logger.info("hello");
    }
```
```java
public class LogTest {
    private Logger logger;
    
    @Before
    public void initialize() {
        logger = Logger.getLogger("logger");
        logger.removeAllAppenders();
        Logger.getRootLogger().removeAllAppenders();
    }
    
    @Test
    public void basicLogger() {
        BasicConfigurator.configure();
        logger.info("basicLogger");
    }
    
    @Test
    public void addAppenderWithStream() {
        logger.addAppender(new ConsoleAppender(
            new PatternLayout("%p %t %m%n"),
            ConsoleAppender.SYSTEM_OUT));
        logger.info("addAppenderWithStream");
    }
    
    @Test
    public void addAppenderWithoutStream() {
        logger.addAppender(new ConsoleAppender(
            new PatternLayout("%p %t %m%n")));
        logger.info("addAppenderWithoutStream");
    }
}
```
- Log4j는 Java/Kotlin/Scala/Groovy/Clojure 코딩 도중에 프로그램의 로그를 기록해주는 라이브러리로, 이클립스, IntelliJ IDEA, 안드로이드 스튜디오 등에 추가해서 프로그램 실행 시 자동으로 지정한 경로에 로그를 저장해주는 기능을 한다.
- 최근에 Log4j 관련하여 우회 공격시 취약점에 노출된다는 기사가 있다.
- 버전에 따라 예방 가능한 것으로 보인다.
- [Log4j 취약점](https://www.whatap.io/ko/blog/100//)

## 학습 테스트는 꽁자 이상이다
- 오히려 필요한 지식만 확보하는 손쉬운 방법이다. 
- 외부코드를 사용시 실제 코드와 동일한 방식으로 인터페이스를 사용하는 테스트 케이스가 필요하다. 새로운 버전으로 이전하기가 수월해진다.
- 낡은 버전을 오랫동안 사용하지 말자.

## 아직 존재하지 않는 코드를 사용하기
- 부족한 모듈을 사용하여 개발을 할 상황이 오는데, 이때 자체적으로 인터페이스를 정의하고
어댑터 패턴으로 API 사용을 캡슐화하여 API가 수정할 코드를 한곳으로 모은다.
- 이렇게 하면 테스트도 아주 편리해 진다.


## 깨끗한 경계

- 소프트웨어 설계가 우수하다면 변경하는데 많은 투자와 재작업이 필요하지 않다. 엄청난 시
간과 노력과 재작업을 요구하지 않는다. 
- 통제하지 못하는 코드를 사용할 때는 너무 많은 투자를 하거나 향후 변경 비용이 지나치게 커지지 않도록 각별히 주의해야 한다.
- 외부 패키지를 사용시 통제 가능한 우리 코드에 의존하여 사용한다. 가능한 코드를 줄여 경계를 환리한다.