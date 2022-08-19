### 목차

- 자료 추상화 🥸
- 자료/객체 비대칭 🤔
- 디미터 법칙 📐
- 자료 전달 객체 📞
- 결론 👨‍🎓





### 자료 추상화 🥸

---

- 변수를 비공개(Private)로 정의하는 이유는 남들이 변수에 의존하지 않게 만들고 싶어서다.
- 변수를 Private로 선언하더라도 각 값마다 조회(getter), 함수와 설정(setter) 함수를 제공한다면 구현을 외부로 노출하는 셈이다.
- 변수 사이에 함수라는 계층을 넣는다고 구현이 저절로 감춰지지는 않는다.
  - 구현을 감추려면 추상화가 필요하다!!
  - 조회 함수와 설정 함수로 변수를 다룬다고 클래스가 되지는 않는다. 그보다는 추상 인터페이스를 사용하여 사용자가 구현을 모른 채 자료의 핵심을 조작할 수 있어야 진정한 의미의 클래스다.

###### 목록 6-1 구체적인 Point 클래스

```swift
public class Point {
  
  public var x: Double
  public var y: Double
  	
 		init(x: Double, y: Double) {
      self.x = x
      self.y = y
    }
}
```



###### 목록 6-2 추상적인 Point 클래스

```swift
Protocol Point {
	func getX() -> Double
  func getY() -> Double
  func setCartesian(x: Double, y: Double)
  func getR() -> Double
  func getTheta() -> Double
  func setPolar(r: Double, theta: Double)
}
```





### 자료/객체 비대칭 🤔

---

- 객체는 추상화 뒤로 자료를 숨긴 채 자료를 다루는 함수만 공개한다. 반면 자료 구조는 자료를 그대로 공개하며 별다른 함수는 제공하지 않는다.

###### 목록 6-5 절차적인 도형

```swift
public class Square {
  public var topLeft: Point
  public var side: Double
}

public class Rectangle {
  public var topLeft: Point
  public var height: Double
  public var width: Double
}

public class Circle {
  public var center: Point
  public var radius: Double
}

public class Geometry {
  public var PI: Double = 3.141592653589793
  
  public func area(shape: AnyObject) throws -> Double {
    if shape === Square {
      var s: Square = shape
      return s.side * s.side
    } else if shape === Rectangle {
      var r: Rectangle = shape
      return r.height * r.width
    } else if shape === Circle {
      var c: Circle = shape
      return PI * c.radius * c.radius
    }
  }
}
```

- 만약 Geometry 클래스에 둘레 길이를 구하는 perimeter() 함수를 추가하고 싶다면? 도형 클래스는 아무 영향도 받지 않는다.
  - 도형 클래스에 의존하는 다른 클래스도 마찬 가지다. 반대로 새 도형을 추가 하고 싶다면? Geometry 클래스에 속한 함수를 모두 고쳐야 한다.



###### 목록 6-6 다형적인 도형

```swift
public class Square: Shape {
  private var topLeft: Point
  private side: Double
  
  public func area() -> Double {
    return side * side
  }
}


public class Reactangle: Shape {
  private topLeft: Point
  private height: Double
  private width: Double
  
  public area() -> Double {
    return height * width
  }
}
```

- 이번 예제는 객체 지향적인 도형 클래스다. 여기서 area()는 다형 메서드다,
  - Geometry 클래스는 필요없으며, 새 도형을 추가해도 기존 함수에 아무런 영향을 미치지 않는다. 반면 새 함수를 추가하고 싶다면 도형 클래스 전부를 고쳐야한다.
- 절차적인 코드는 새로운 자료구조를 추가하기 어렵다. 반면 객체 지향 코드는 새로운 함수를 추가하기 어렵다.
- 절차적인 코드는 기존 자료구조를 변경하지 않으면서 새 함수를 추가하기 쉽다. 반면 객체 지향 코드는 기존 함수를 변경하지 않으면서 새 클래스를 추가하기 쉽다.
- 복잡한 시스템을 짜다 보면 새로운 함수가 아니라 새로운 자료 타입이 필요한 경우가 생긴다. 이때는 클래스와 객체 지향 기법이 가장 적합하다.
  - 새로운 자료 타입이 아니라 새로운 함수가 필요한 경우도 생긴다. 이때는 절차적인 코드와 자료 구조가 좀 더 적합하다.
- 분별 있는 프로그래머는 모든 것이 객체라는 생각이 미신임을 잘안다.!! 떄로는 단순한 자료 구조와 절차적인 코드가 가장 적합한 상황도 있다.



### 디미터 법칙 📐

---

- 디미터 법칙은 잘 알려진 휴리스틱으로, 모듈은 자신이 조작하는 객체의 속사정을 몰라야 한다는 법칙이다.
  - 휴리스틱: **휴리스틱**(heuristics) 또는 **발견법**(發見法)이란 불충분한 시간이나 정보로 인하여 합리적인 판단을 할 수 없거나, 체계적이면서 합리적인 판단이 굳이 필요하지 않은 상황에서 사람들이 빠르게 사용할 수 있게 보다 용이하게 구성된 간편추론의 방법이다.
- 객체는 조회 함수로 내부 구조를 공개하면 안 된다는 의미다. 그러면 내부 구조를 노출하는 셈이니까.
  - 디미터 법칙 : 클래스 C의 메서드 F는 다음과 같은 객체의 메서드만 호출해야 한다
    - 클래스 C
    - f 가 생성한 객체
    - f 인수로 넘어온 객체
    - C 인스턴스 변수에 저장된 객체
- 위 객체에서 허용된 메서드가 반환하는 객체의 메서드는 호출하면 안 된다. 다시 말해 낯선 사람은 경계하고 친구랑만 놀라는 의미다.



###### 기차 충돌 Code

```swift
var outputDir: String = ctxt.getOptions().getScratchDir().getAbsolutePate()
```

- 여러 객차가 한 줄로 이어진 기차처럼 보이기 때문이다. 일방적으로 조잡하다 여겨지는 방식이므로 피하는 편이 좋다.

###### 기차 충돌 Code Refactoring

```swift
var opts = ctxt.getOptions()
var scratchDir = opts.getScratchDir()
var outputDir = scratchDir.getAbsolutePate()
```

- 위 예제가 디미터 법칙을 위반 할까?
  - 디미터 법칙을 위반하는지 여부는 ctxt, Options, ScratchDir이 객체인지 아니면 자료 구조인지에 달렸다.
    - 객체 라면 내부 구조를 숨겨야 하므로 확실히 디미터 법칙을 위반한다. 반면 자료 구조라면 당연히 내부 구조를 노출하므로 디미터 법칙이 적용되지 않는다.

###### 디미터 법칙 통과 Code

```swift
var outputDir: String = ctxt.options.scratchDir.absolutePath 
```

- 자료 구조는 무조건 함수 없이 공개 변수만 포함하고 객체는 비공개 변수와 공개 함수를 포함한다면, 문제는 훨씬 간단하리라
  - 단순한 자료 구조에도 조회 함수와 설정 함수를 정의하라 요구하는 프레임워크와 표준(bean)이 존재한다.



### 자료 전달 객체 📞

---

- 자료 구조체의 전형적인 형태는 공개 변수만 있고 함수가 없는 클래스다.
  - 이런 자료 구조체를 때로는 자료 전달 객체(Data Transfer Object) DTO라 한다.  DTO는 굉장히 유용한 구조체다.
- 활성 레코드는 DTO의 특수한 형태다. 공개 변수가 있거나 비공개 변수에 조회/설정 함수가 있는 자료 구조지만, 대개 save나 find와 같은 탐색 함수도 제공한다.
- 활성 레코드에 비즈니스 규칙 메서드를 추가해 이런 자료 구조를 객체로 취급하는 개발자가 흔하다.
  - 하지만 이는 바람직하지 않다. 그러면 자료구조도 아니고 객체도 아닌 잡종 구조가 나오기 때문이다.



### 결론 👨‍🎓

---

- 객체는 동작을 공개하고 자료를 숨긴다. 그래서 기존 동작을 변경하지 않으면서 새 객체 타입을 추가하기는 쉬운 반면, 기존 객체에 새 동작을 추가하기는 어렵다.
- 자료구조는 별다른 동작 없이 자료를 노출한다. 그래서 기존 자료 구조에 새 동작을 추가하기는 쉬우나, 기존 함수에 새 자료 구조를 추가하기는 어렵다.

