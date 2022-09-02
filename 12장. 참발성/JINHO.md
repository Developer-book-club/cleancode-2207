## 목차  

### 창발적 설계로 깔끔한 코드를 구현하자

- 우리들 대다수는 켄트 벡이 제시한 단순한 설계 규칙 네 가지3가 소프트웨어 설계 품질을 크게 높여준다고 믿는다.

> 모든 테스트를 실행한다.
> 중복을 없앤다.
> 프로그래머 의도를 표현한다.
> 클래스와 메서드 수를 최소로 줄인다.

### 단순한 설계 규칙 1: 모든 테스트를 실행하라

- 설계는 의도한 대로 돌아가는 시스템을 내놓아야 한다.
- 시스템이 의도한 대로 돌아가는지 검증할 간단한 방법이 없다면, 문서 작성을 위해 투자한 노력에 대한 가치는 인정받기 힘들다.
-  검증이 불가능한 시스템은 절대 출시하면 안 된다.
- 철저한 테스트가 가능한 시스템을 만들면 더 나은 설계가 얻어진다.
-  테스트 케이스를 많이 작성할수록 개발자는 DIP와 같은 원칙을 적용하고 의존성 주입Dependency Injection, 인터페이스, 추상화 등과 같은 도구를 사용해 결합도를 낮춘다. 따라서 설계 품질은 더욱 높아진다.

### 단순한 설계 규칙 2~4: 리팩터링

- 테스트 케이스를 모두 작성했다면 이제 코드와 클래스를 정리해도 괜찮다. 구체적으로는 코드를 점진적으로 리팩터링 해나간다.
- 코드를 정리하면서 시스템이 깨질까 걱정할 필요가 없다. 테스트 케이스가 있기 때문이다.
- 리팩터링 단계에서는 소프트웨어 설계 품질을 높이는 기법이라면 무엇이든 적용해도 괜찬다.
- 응집도를 높이고, 결합도를 낮추고, 관심사를 분리하고, 시스템 관심 모듈을 나누고, 함수와 클래스 크기를 줄이는 등 다양한 기법을 동원한다.
- 또한 단순한 설계 규칙중 나머지 3개를 적용해 중복을 제거하고, 프로그래머 의도를 표현하고, 클래스와 메서드 수를 최소로 줄이는 단계이기도 하다.

### 중복을 없애라

- 우수한 설계에서 중복은 커다란 적이다. 복잡도를 뜻하기 때문이다.
- 비슷한 코드는 더 비슷하게 고치면 리팩토링이 쉬워진다.

```java
int size() {}
boolean isEmpty{}
```

```java
boolean isEmpty() {
  return 0 == size();
}
```
- size를 개수를 반환하는 로직이므로 중복해서 구현할 필요 x 

```java
public void scaleToOneDimension(float desiredDimension, float imageDimension) {
  if (Math.abs(desiredDimension - imageDimension) < errorThreshold)
    return;
  float scalingFactor = desiredDimension / imageDimension;
  scalingFactor = (float)(Math.floor(scalingFactor * 100) * 0.01f);
  
  RenderedOpnewImage = ImageUtilities.getScaledImage(image, scalingFactor, scalingFactor);
  image.dispose();
  System.gc();
  image = newImage;
}

public synchronized void rotate(int degrees) {
  RenderedOpnewImage = ImageUtilities.getRotatedImage(image, degrees);
  image.dispose();
  System.gc();
  image = newImage;
}
```

```java
public void scaleToOneDimension(float desiredDimension, float imageDimension) {
  if (Math.abs(desiredDimension - imageDimension) < errorThreshold)
    return;
  float scalingFactor = desiredDimension / imageDimension;
  scalingFactor = (float) Math.floor(scalingFactor * 10) * 0.01f);
  replaceImage(ImageUtilities.getScaledImage(image, scalingFactor, scalingFactor));
}

public synchronized void rotate(int degrees) {
  replaceImage(ImageUtilities.getRotatedImage(image, degrees));
}

private void replaceImage(RenderedOpnewImage) {
  image.dispose();
  System.gc();
  image = newImage;
}
```
- 중복코드를 따로 뺀듯하다.

```java
public class VacationPolicy {
  public void accrueUSDDivisionVacation() {
    // 지금까지 근무한 시간을 바탕으로 휴가 일수를 계산하는 코드
    // ...
    // 휴가 일수가 미국 최소 법정 일수를 만족하는지 확인하는 코드
    // ...
    // 휴가 일수를 급여 대장에 적용하는 코드
    // ...
  }
  
  public void accrueEUDivisionVacation() {
    // 지금까지 근무한 시간을 바탕으로 휴가 일수를 계산하는 코드
    // ...
    // 휴가 일수가 유럽연합 최소 법정 일수를 만족하는지 확인하는 코드
    // ...
    // 휴가 일수를 급여 대장에 적용하는 코드
    // ...
  }
}
```

```java
abstract public class VacationPolicy {
  public void accrueVacation() {
    caculateBseVacationHours();
    alterForLegalMinimums();
    applyToPayroll();
  }
  
  private void calculateBaseVacationHours() { /* ... */ };
  abstract protected void alterForLegalMinimums();
  private void applyToPayroll() { /* ... */ };
}

public class USVacationPolicy extends VacationPolicy {
  @Override protected void alterForLegalMinimums() {
    // 미국 최소 법정 일수를 사용한다.
  }
}

public class EUVacationPolicy extends VacationPolicy {
  @Override protected void alterForLegalMinimums() {
    // 유럽연합 최소 법정 일수를 사용한다.
  }
}
```

- 다형성을 이용해 중복코드를 따로 분리하였다.

### 표현하라

- 소프트웨어 프로젝트 비용 중 대다수는 장기적인 유지보수에 들어간다.
- 하지만 시스템이 점차 복잡해지면서 유지보수 개발자가 시스템을 이해하느라 보내는 시간은 점점 늘어나고 동시에 코드를 오해할 가능성도 점점 커진다.
- 좋은 이름을 선택해 개발자를 이해시킨다.
- 함수와 클래스 크기를 가능한 줄인다.
- 표준 명칭을 사용한다.
- 단위 테스트 케이스를 꼼꼼히 작성한다.

### 클래스와 메서드 수를 최소로 줄여라

- SRP 준수한다는 기본적인 개념도 극단으로 치달으면 득보다 실이 많아진다.
- 때로는 무의미하고 독단적인 정책 탓에 클래스 수와 메서드 수가 늘어나기도 한다.
- 목표는 함수의 클래스 크기를 작게 유지하면서 동시에 시스템 크기도 작게 유지하는 데 있다.