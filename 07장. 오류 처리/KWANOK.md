# 오류 처리



## 오류 코드보다 예외를 사용합시다.

중간중간 오류를 하나하나 확인하는 건 까다로운 작업입니다.

코드의 가독성도 떨어지고, 호출자 코드가 복잡해집니다.

```java
public class DeviceController {
    public void sendShutDown() {
        if (상태점검문제없다면) {
            // 레코드 필드에 디바이스 상태를 저장
            retrieveDeviceRecord(handle);
            
            if (디바이스가 일시정지 상태가 아니라면) {
                pauseDevice(handle);
                clearDeviceWorkQueue(handle);
            } else {
                logger.log("비상")
            }
        } else {
            logger.log("비상")
        }
    }
}
```

### 예외 사용을 한다면

코드가 더 깔끔해집니다.

오류를 처리 과정과 디바이스 종료 과정을 분리했기 때문에 신경이 분산되지 않는 것 같습니다.

```java
public class DeviceController {
    public void sendShutDown() {
        try {
            tryToShutDown();
        } catch (DeviceShutDownError e) {
            logger.log(e);
        }
    }
    
    private void tryToShutDown() throws DeviceShutDownError {
        DeviceHandle handle = getHandle(DEV1);
        ...
        
        pauseDevice(handle);
        clearDeviceWorkQueue(handle);
        closeDevice(handle)
    }
    
    private Device Handle getHandle(DeviceID id) {
        ...
        throw new DeviceShutDownError("Invalid handle for: " + id.toString());
    }
}
```

## Try-Catch-Finally

## + Golang은 SHE(Structured Error Handling)를 지원하지 않습니다.

앞에 설명한 내용은 모두 try catch를 사용했습니다.

하지만 이런 예외처리는 성능 이슈도 있고,

오류를 인지하기 어렵다고 합니다. 이로 인해 처리를 하지 않고, 방치하는 경우가 생긴다고 합니다.

## 예외에 의미를 제공합시다

예외를 던질 때는 그냥 오류라고만 던지는 것이 아니라

전후 상황에 대한 내용을 충분하게 담아서 던지는 게 좋습니다.

그렇게 하면, 오류가 생겼을 때 더 쉽게 원인을 파악할 수 있기 때문입니다.

## Error를 잘 Wrapping 해서 사용합시다.

이렇게 나열하는 건 보기 별로 좋지 않습니다.

그리고 오류를 처리하는 방식은 비교적 일정합니다.

```java
ACMEPort port = new ACMEPort(12);

try {
    port.open();
} catch (DeviceResponseException e) {
    reportPortError(e):
    logger.log("Device response exception", e)
} catch (ATM1212UnlockedException e) {
    reportPortError(e):
    logger.log("Unlock exception", e)
} catch (GMXError e) {
    reportPortError(e):
    logger.log("Device response exception", e)
} finally {
    ...
}
```

이럴 땐 `ACMEPort` 클래스를 한 번 감싸는 것이 좋은 방법입니다.

```java
LocalPort port = new LocalPort(12);
try {
    port.open();
} catch (PortDeviceFailure e) {
    reportError(e)
    logger.log(e.getMessage(), e)
} finally {
    ...
}
```

```java
public class LocalPort {
    private ACMEPort innerPort;
    
    public LocalPort(int portNumber) {
        innerPort = new ACMEPort(portNumber);
    }
    
    public void open() {
        try {
            innerPort.open();
        } catch (DeviceResponseException e) {
            throw new PortDeviceFailure(e)
        } catch (ATM1212UnlockedException e) {
            throw new PortDeviceFailure(e)
        } catch (GMXError e) {
            throw new PortDeviceFailure(e)
        } 
    }
}
```

## NULL을 반환하지 마세요

오류가 생기는 이유도 방지해야 합니다.

null을 반환하는 습관은 우리가 항상 null값을 확인해야하게 만듭니다.

굳이 null을 반환하고 싶다면 차라리 빈 객체나 예외를 던지는 것이 더 좋다고 합니다.





