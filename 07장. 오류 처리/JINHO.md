 # 오류처리

### 오류 코드보다 예외를 사용하라

```java
// Bad
public class DeviceController {
  ...
  public void sendShutDown() {
    DeviceHandle handle = getHandle(DEV1);
    // Check the state of the device
    if (handle != DeviceHandle.INVALID) {
      // Save the device status to the record field
      retrieveDeviceRecord(handle);
      // If not suspended, shut down
      if (record.getStatus() != DEVICE_SUSPENDED) {
        pauseDevice(handle);
        clearDeviceWorkQueue(handle);
        closeDevice(handle);
      } else {
        logger.log("Device suspended. Unable to shut down");
      }
    } else {
      logger.log("Invalid handle for: " + DEV1.toString());
    }
  }
  ...
}
```
```java
// Good
public class DeviceController {
  ...
  public void sendShutDown() {
    try {
      tryToShutDown();
    } catch (DeviceShutDownError e) {
      logger.log(e);
    }
  }
    
  private void tryToShutDown() throws DeviceShutDownError {
    DeviceHandle handle = getHandle(DEV1);
    DeviceRecord record = retrieveDeviceRecord(handle);
    pauseDevice(handle); 
    clearDeviceWorkQueue(handle); 
    closeDevice(handle);
  }
  
  private DeviceHandle getHandle(DeviceID id) {
    ...
    throw new DeviceShutDownError("Invalid handle for: " + id.toString());
    ...
  }
  ...
}
```

- 위와 같은 방법을 사용하면 코드가 복잡해진다.
- 오류가 발생하면 예외처리를 하는 것이 코드가 더 깔끔해진다.
- 디바이스를 종료하는 알고리즘과 오류를 처리하는 알고리즘을 분리했다.

### Try-Catch-Finally 문부터 작성하라.

```java
  @Test(expected = StorageException.class)
  public void retrieveSectionShouldThrowOnInvalidFileName() {
    sectionStore.retrieveSection("invalid - file");
  }
  
  public List<RecordedGrip> retrieveSection(String sectionName) {
    // dummy return until we have a real implementation
    return new ArrayList<RecordedGrip>();
  }
```
```java
  public List<RecordedGrip> retrieveSection(String sectionName) {
    try {
      FileInputStream stream = new FileInputStream(sectionName)
    } catch (Exception e) {
      throw new StorageException("retrieval error", e);
    }
  return new ArrayList<RecordedGrip>();
}
```
```java
  public List<RecordedGrip> retrieveSection(String sectionName) {
    try {
      FileInputStream stream = new FileInputStream(sectionName);
      stream.close();
    } catch (FileNotFoundException e) {
      throw new StorageException("retrieval error", e);
    }
    return new ArrayList<RecordedGrip>();
  }
```

 - try-catchfinally 문에서 try 블록에 들어가는 코드를 실행하면 어느 시점에서든 실행이 중단된 후 catch 블록으로 넘어갈 수 있다.
- 강제로 예외를 일으키는 테스트 케이스를 작성한 후 테스트를 통과하게 코드를 작성하는 방법을 권장한다.
- try 블록의 트랜잭션 범위부터 구현하게 되므로 Scope의 정의와 트랜잭션의 본질을 유지하기 쉬워진다.

### 미확인 예외를 사용하라.

- 확인된 오류가 치르는 비용에 상응하는 이익을 제공하는지 (철저히) 따져봐야한다.
- checked 예외 는 컴파일 단계에서 확인되며 반드시 처리해야 하는 예외.
- IoException
- SQLException
- Unchecked 예외 는 실행 단계에서 확인되며 명시적인 처리를 강제하지는 않는 예외.
- NullPointer, IllegalArgument, SystemException
- 확인된 예외는 예상되는 모든 예외를 사전에 처리할 수 있다는 장점이 있지만, 일반적인 애플리케이션은 의존성이라는 비용이 이익보다 더 크다.
- 해당 함수 뿐만이 아니라 호출하는 함수도 수정을 해줘야 하기 때문에 OCP 를 위반하게 된다.

### 예외의 의미를 제공하라.

- 오류 메시지에 정보를 담아 예외와 함께 던진다.
- 오류가 발생한 원인과 위치를 찾기가 쉬워진다.
- 로깅을 사용하면 catch블록에서 오류를 기록하도록 충분한 정보를 남겨준다.

### 호출자를 고려해 예외 클래스를 정의하라.

```java
  ACMEPort port = new ACMEPort(12);
  try {
    port.open();
  } catch (DeviceResponseException e) {
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
```java
  LocalPort port = new LocalPort(12);
  try {
    port.open();
  } catch (PortDeviceFailure e) {
    reportError(e);
    logger.log(e.getMessage(), e);
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
      } catch (ATM1212UnlockedException e) {
        throw new PortDeviceFailure(e);
      } catch (GMXError e) {
        throw new PortDeviceFailure(e);
      }
    }
    ...
  }
```

- 위 두개를 비교하면 예외처리를 외부 API로 감싸고 아니고의 차이가 있다.
- 2번쨰 방법시
> 에러처리가 간결해짐
> 외부 라이브러리와 프로그램 사이의 의존성이 크게 줄어듦
> 라이브러리의 API가 종속적이지 않고 외부 API 방식에 의존하지 않아도 된다.

### 정상 흐름을 정의하라.

- catch문을 사용하지 않고 예외적인 상황을 캡슐화하여 클라이언트 코드가 처리할 필요가 없도록 할 수있다.


### null을 반환하지 마라.

```java
  try {
    MealExpenses expenses = expenseReportDAO.getMeals(employee.getID());
    m_total += expenses.getTotal();
  } catch(MealExpensesNotFound e) {
    m_total += getMealPerDiem();
  }
```
```java
  ...
  MealExpenses expenses = expenseReportDAO.getMeals(employee.getID());
  m_total += expenses.getTotal();
  ...
  
  public class PerDiemMealExpenses implements MealExpenses {
    public int getTotal() {
      // return the per diem default
    }
  }
```

- null을 반환하지 말고 차라리 예외를 던지거나 특수 사례로 반환하는것이 낫다.
- NullPoint 에러 발생할수 있음.
- null을 체크할 일이 계속 생김

### null을 전달하지 마라

- 정상적인 인수로 null을 기대하는 API가 아니라면 메서드로 null을
전달하는 코드는 최대한 피한다
- 일반적으로 대다수의 프로그래밍 언어들은 파라미터로 들어온 null에 대해 적절한 방법을 제공하지 못한다.
- 가장 이성적인 해법은 null을 파라미터로 받지 못하게 하는 것이다.
- null을 넘기지 못하도록 금지하는 정책이 합리적이다. 
