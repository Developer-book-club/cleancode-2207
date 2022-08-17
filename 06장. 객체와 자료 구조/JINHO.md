# 객체와 자료구조 정리

## 자료 추상화

```java
public class Point { 
  public double x; 
  public double y;
}
```

```java
public interface Point {
  double getX();
  double getY();
  void setCartesian(double x, double y); 
  double getR();
  double getTheta();
  void setPolar(double r, double theta); 
}
```

- 추상적인 Point는 해당 점이 내부적으로 직교좌표계를 사용하는지, 극 좌표계를 사용하는지 알 수 없다.
- 구체적인 Point는 직교 좌표계로 구현되어 있고, 이를 직접 사용할 수 있게 한다.
- 반적으로 변수를 private으로 많이 선언을 하는데, 각 값마다 get과 set 함수를 제공한다면 이는 결과적으로 내부 구조를 노출하는 구조가 된다.
- 추상 인터페이스를 제공해 사용자가 구현을 모른 채 자료의 핵심을 조작할 수 있어야 진정한 의미의 클래스다.

##자료/객체 비대칭


```java
public class Square { 
  public Point topLeft; 
  public double side;
}

public class Rectangle { 
  public Point topLeft; 
  public double height; 
  public double width;
}

public class Circle { 
  public Point center; 
  public double radius;
}

public class Geometry {
  public final double PI = 3.141592653589793;
  
  public double area(Object shape) throws NoSuchShapeException {
    if (shape instanceof Square) { 
      Square s = (Square)shape; 
      return s.side * s.side;
    } else if (shape instanceof Rectangle) { 
      Rectangle r = (Rectangle)shape; 
      return r.height * r.width;
    } else if (shape instanceof Circle) {
      Circle c = (Circle)shape;
      return PI * c.radius * c.radius; 
    }
    throw new NoSuchShapeException(); 
  }
}
```
```java
public class Square implements Shape { 
  private Point topLeft;
  private double side;
  
  public double area() { 
    return side * side;
  } 
}

public class Rectangle implements Shape { 
  private Point topLeft;
  private double height;
  private double width;

  public double area() { 
    return height * width;
  } 
}

public class Circle implements Shape { 
  private Point center;
  private double radius;
  public final double PI = 3.141592653589793;

  public double area() {
    return PI * radius * radius;
  } 
}
```

- 절차지향적 구현
- 새로운 도형의 추가는 class에 영향이 없음
- 도형의 둘레를 구하는 perimeter() 함수를 추가하려면 Geometry 클래스만 변경하면 됨

- 객체지향적 구현
- Gemotery 클래스 불필요, 새로운 도형의 추가로 인해 area()가 변경될 필요가 없음
- 도형에 대해 둘레를 구하는 perimeter()를 추가하려면 모든 클래스가 변경되어야 함

##디미터의 법칙

- 모듈은 자신이 조작하는 객체의 속사정을 몰라야 한다는 법칙이다.
- 소프트웨어 모듈 사이의 결합도를 줄여서 코드의 품질을 높이는 가이드 라인이다.

- 클래스 C와, 해당 클래스의 메서드 f
- 메서드 f가 생성한 객체
- 메서드 f에 인자로 전달 된 객체
- 클래스 C의 인스턴스 변수 안에 있는 객체

- 객체의 메서드들만 호출할 것을 말하고 있다. 해당 객체에 직접적으로 연관된 객체들만 사용할 것을 말하고 있다.
-  메서드 내부에서만 다른 객체를 사용하면 각 객체간의 결합도를 알 수 없기 때문인 것으로 보인다.

###기차충돌

```java
final String outputDir = ctxt.getOptions().getScratchDir().getAbsolutePath();
```

```java
Options opts = ctxt.getOptions();
File scratchDir = opts.getScratchDir();
final String outputDir = scratchDir.getAbsolutePath();
```

- 함수 호출이 연속되면, 조잡한 방식이고 객체의 결과값을 사용하므로 이런식의 코드는 좋지 않다.
- 디미터의 법칙 위반 여부는 결과 객체인 FIle, Options가 객체/자료구조에 따라 다르다.
- 자료구조라면 당연히 데이터를 노출하고, 객체면 내부구조를 숨겨야 한다.

### 잡종구조

- 위에 기차 충돌의 코드의 결과값을 자료구조라 생각하면 디미터의 법칙을 지켰으나, 기존의 코드를 보면 getter 형태로 메서드가 호출되었다.
- 이 중간 객체들은 함수를 호출하는 형태로 봤을때, 객체임에도 디미터의 법칙을 지키기 위해 내부를 개방하면 객체이자 자료구조가 된다.

### 구조체 감추기

```java
String outFile = outputDir + "/" + className.replace('.', '/') + ".class"; 
FileOutputStream fout = new FileOutputStream(outFile); 
BufferedOutputStream bos = new BufferedOutputStream(fout);
```

```java
ufferedOutputStream bos = ctxt.createScratchFileStream(classFileName);```

- 디미터의 법칙은, 중간 변수들은 드러나서 안된다.
- 하지만 이렇게 되면 처음 호출한 객체가 하위 객체의 정보를 모두 알아야 하는데 결합도가 높아지는 이슈가 생긴다.
- 허나 이름으로 의도를 유추하면 임시 경로를 얻는것은 아닌걸로 보인다.
- 해당 변수가 사용되는 코드를 보니 임시 디렉토리에 임시 파일을 생성하는 것이면, 이와 관련된 부분을 더 숨길 수 있다.

### 자료 전달 객체

- 데이터베이스의 값을 가져오기 위해 자료구조만 표현하는 형태의 객체가 사용되는데, 이를 DTO(Data Transfer Object)라 한다.


```java
public class Address {
    private String street;
    private String streetExtra;
    private String city;
    private String state;
    private String zip;
    
    public Address(String street, String streetExtra,
                    String city, String state, String zip) {
        this.street = street;
        this.streetExtra = streetExtra;
        this.city = city;
        this.state = state;
        this.zip = zip;
    }

    public String getStreet() {
        return street;
    }

    public String getStreetExtra() {
        return streetExtra;
    }

    public String getCity() {
        return city;
    }
    
    public String getState() {
        return state;
    }
    
    public String getZip() {
        return zip;
    }
}
```

### 활성 레코드

- 데이터베이스의 테이블을 직접 변환하여 기존 DTO에 find, save등의 레코드 작어블 수행하는 형테를 활성 레코드라 한다.
- Java에서 toString 메서드도 사용
- 이런 자료구조에 비즈니스로직에는 객체와 자료구조만을 사용하는 것을 지양한다.

### 결론

- 객체지항은 완전이 이해하기 어려운 주제이다.
-   디자인 패턴이 코드의 구조를 규격화하는 데 유용한 도구지만, 무분별하게 패턴을 적용할 필요가 없듯, 디미터의 법칙도 적절히 적용하는 것이 맞을 것 같다.