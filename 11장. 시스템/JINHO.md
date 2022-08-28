##　11장　시스템
도시를 세운다면?
시스템 제작과 시스템 사용을 분리하라
확장
자바 프록시
순수 자바 AOP 프레임워크
AspectJ 관점
테스트 주도 시스템 아키텍처 구축
의사 결정을 최적화하라
명백한 가치가 있을 때 표준을 현명하게 사용하라
시스템은 도메인 특화 언어가 필요하다
결론
### 도시를 세운다면?

－　흔히 소프트웨어 팀도 도시처럼 구성한다. 그런데 막상 팀이 제작하는 시스템은 비슷한 수준으로 관심사를 분리하거나 추상화를 이뤄내지 못한다.
－　소프트웨어 또한 이와 비슷한 방식으로 구성되기는 하나 도시의 모듈화 만큼의 추상화를 이루지 못하는 경우가 많다.

### 시스템의 생성과 사용을 분리하라

```java
public Service getService() {
    if (service == null)
        service = new MyServiceImpl(...); // 모든 상황에 적합한 기본값일까?
    return service;
}
```

- 소프트웨어 시스템은 (애플리케이션 객체를 제작하고 의존성을 서로 ‘연결’하는)
준비 과정과 (준비 과정 이후에 이어지는) 런타임 로직을 분리해야 한다.
-  getService 메서드가 MyServiceImpl과 (위에서는 생략한) 생성자 인수에 명시적으로 의존한다. 런타임 로직에서 MyServiceImpl 객체를 전혀 사용하지 않더라도 의존성을 해결하지 않으면 컴파일이 안 된다.
- 체계적이고 탄탄한 시스템을 만들고 싶다면 흔히 쓰는 좀스럽고 손쉬운 기법으로 모듈성을 깨서는 절대로 안 된다.

#### Main 분리

- 시스템 생성과 시스템 사용을 분리하는 한 가지 방법으로, 생성과 관련한 코드
는 모두 main이나 main이 호출하는 모듈로 옮긴다.
-  main 함수에서 시스템에 필요한 객체를 생성한 후 이를 애플리케이션에 넘긴다. 
- 애플리케이션은 main이나 객체가 생성되는 과정을 전혀 모른다는 뜻이다. 


#### 팩토리

- 물론 때로는 객체가 생성되는 시점을 애플리케이션이 결정할 필요도 생긴다. 
- 책에서는 추상팩토리를 사용하여 main과 애플리케이션을 분리했다.

#### 의존성 주입

- 밥을 지을때 필요한 쌀을 직접 만들지 말고 가져와서 사용하라.
- 의존성 주입은 제어 역전Inversion of Control, IoC 기법3을 의존성 관리4에 적용한 메커니즘이다.
- 제어 역전에서는 한 객체가 맡은 보조 책임을 새로운 객체에게 전적으로 떠넘긴다.
- 초기 설정은 시스템 전체에서 필요하므로 대개 ‘책임질’ 메커니즘으로 ‘main’ 루틴이나 특수 컨테이너를 사용한다.

```java
    MyService myService = (MyService)(jndiContext.lookup(“NameOfMyService”));
```

- 호출하는 객체는 (반환되는 객체가 적절한 인터페이스를 구현하는 한) 실제로 반환되는 객체의 유형을 제어하지 않는다. 
- 진정한 의존성 주입은 여기서 한 걸음 더 나간다. 클래스가 의존성을 해결하려 시도하지 않는다. 클래스는 완전히 수동적이다. 

### 확장

- ‘처음부터 올바르게’ 시스템을 만들 수 있다는 믿음은 미신이다. 대신에 우리
는 오늘 주어진 사용자 스토리에 맞춰 시스템을 구현해야 한다.
- 내일은 새로운스토리에 맞춰 시스템을 조정하고 확장하면 된다. 
- 이것이 반복적이면서 점진적인 애자일 개발 방법론이다.
- 테스트 주도 개발Test-driven Development, TDD, 리팩터링, 깨끗한 코드는 코드 수준에서 시스템을 조정하고 확장하기 쉽게 만든다.
- 소프트웨어 시스템은 물리적인 시스템과 다르다. 관심사를 적절히 분리해 관리한다면 소프트웨어 아키텍처는 점진적으로 발전할 수 있다.

```java
package com.example.banking;
import java.util.Collections;
import javax.ejb.*;

public interface BankLocal extends java.ejb.EJBLocalObject {
    String getStreetAddr1() throws EJBException;
    String getStreetAddr2() throws EJBException;
    String getCity() throws EJBException;
    String getState() throws EJBException;
    String getZipCode() throws EJBException;
    void setStreetAddr1(String street1) throws EJBException;
    void setStreetAddr2(String street2) throws EJBException;
    void setCity(String city) throws EJBException;
    void setState(String state) throws EJBException;
    void setZipCode(String zip) throws EJBException;
    Collection getAccounts() throws EJBException;
    void setAccounts(Collection accounts) throws EJBException;
    void addAccount(AccountDTO accountDTO) throws EJBException;
}
```

```java
package com.example.banking;
import java.util.Collections;
import javax.ejb.*;

public abstract class Bank implements javax.ejb.EntityBean {
    // Business logic...
    public abstract String getStreetAddr1();
    public abstract String getStreetAddr2();
    public abstract String getCity();
    public abstract String getState();
    public abstract String getZipCode();
    public abstract void setStreetAddr1(String street1);
    public abstract void setStreetAddr2(String street2);
    public abstract void setCity(String city);
    public abstract void setState(String state);
    public abstract void setZipCode(String zip);
    public abstract Collection getAccounts();
    public abstract void setAccounts(Collection accounts);
    
    public void addAccount(AccountDTO accountDTO) {
        InitialContext context = new InitialContext();
        AccountHomeLocal accountHome = context.lookup("AccountHomeLocal");
        AccountLocal account = accountHome.create(accountDTO);
        Collection accounts = getAccounts();
        accounts.add(account);
    }
    
    // EJB container logic
    public abstract void setId(Integer id);
    public abstract Integer getId();
    public Integer ejbCreate(Integer id) { ... }
    public void ejbPostCreate(Integer id) { ... }
    
    // The rest had to be implemented but were usually empty:
    public void setEntityContext(EntityContext ctx) {}
    public void unsetEntityContext() {}
    public void ejbActivate() {}
    public void ejbPassivate() {}
    public void ejbLoad() {}
    public void ejbStore() {}
    public void ejbRemove() {}
}
```

위 코드와 같은 전형적인 EJB2 객체 구조는 아래와 같은 문제점을 가지고 있다.  
- 비즈니스 논리는 EJB 컨테이너와 강하게 결합되어 클래스를 생성할 때는 컨테이너에서 파생해야 하며 컨테이너가 요구하는 다양한 생명주기 메서드도 제공해야 한다.
- 이렇듯 비즈니스 논리가 덩치 큰 컨테이너와 밀접하게 결합된 탓에 독자적인 단위 테스트가 어렵다.

#### 횡단(cross-cutting) 관심사

- 원론적으로는 모듈화되고 캡슐화된 방식으로 영속성 방식을 구상할 수 있다. 하지만 현실적으로는 영속성 방식을 구현한 코드가 온갖 객체로 흩어진다. 여기서 횡단 관심사라는 용어가 나온다.
- 사실 EJB 아키텍처가 영속성, 보안, 트랜잭션을 처리하는 방식은 관점 지향
프로그래밍Aspect-Oriented Programming, AOP을 예견했다고 보인다. 
- AOP의 개념은 "특정 관심사를 지원하려면 시스템에서 특정 지점들이 동작하는 방식을 일관성 있게 바꿔야 한다"

### 자바 프록시

- 단순한 상황에서 개별 객체나 클래스에서 메서드 호출을 감싸는 경우가 좋은 예다.
- 클래스 프록시를 사용하려면 CGLIB, ASM, Javassist10 등과 같은 바이트 코드 처리 라이브러리가 필요하다.

```java
// Bank.java (패키지 이름을 감춘다)
import java.util.*;

// 은행 추상화
public interface Bank {
    Collection<Account> getAccounts();
    void setAccounts(Collection<Account> accounts);
}

// BankImpl.java
import java.utils.*;

// 추상화를 위한 POJO("Plain Old Java Object") 구현
public class BankImpl implements Bank {
    private List<Account> accounts;

    public Collection<Account> getAccounts() {
        return accounts;
    }
    
    public void setAccounts(Collection<Account> accounts) {
        this.accounts = new ArrayList<Account>();
        for (Account account: accounts) {
            this.accounts.add(account);
        }
    }
}
// BankProxyHandler.java
import java.lang.reflect.*;
import java.util.*;

// 프록시 API가 필요한 "InvocationHandler"
public class BankProxyHandler implements InvocationHandler {
    private Bank bank;
    
    public BankHandler (Bank bank) {
        this.bank = bank;
    }
    
    // InvocationHandler에 정의된 메서드
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
        String methodName = method.getName();
        if (methodName.equals("getAccounts")) {
            bank.setAccounts(getAccountsFromDatabase());
            
            return bank.getAccounts();
        } else if (methodName.equals("setAccounts")) {
            bank.setAccounts((Collection<Account>) args[0]);
            setAccountsToDatabase(bank.getAccounts());
            
            return null;
        } else {
            ...
        }
    }
    
    // 세부사항은 여기에 이어진다.
    protected Collection<Account> getAccountsFromDatabase() { ... }
    protected void setAccountsToDatabase(Collection<Account> accounts) { ... }
}

// 다른 곳에 위치하는 코드
Bank bank = (Bank) Proxy.newProxyInstance(
    Bank.class.getClassLoader(),
    new Class[] { Bank.class },
    new BankProxyHandler(new BankImpl())
);
```
- 위에서는 프록시로 감쌀 인터페이스 Bank와 비즈니스 논리를 구현하는 POJO Plain Old Java Object BankImpl을 정의했다.
- 프록시 API에는 InvocationHandler를 넘겨 줘야 한다. 넘긴 InvocationHandler는 프록시에 호출되는 Bank 메서드를 구현하는 데 사용된다. 
- BankProxyHandler는 자바 리플렉션 API를 사용해 제네릭스 메서드를 상응하는 BankImpl 메서드로 매핑한다.

- 코드가 상당히 많으며 제법 복잡하다.(프록시 사용시 깨끗한 코드 작성이 어려움)
- 프록시는 시스템 단위로 실행 ‘지점’을 명시하는 메커니즘도 제공하지 않는다.

### 순수 자바 AOP 프레임워크

- 대부분의 프록시 코드는 판박이라 도구로 자동화할 수 있다. 
- 자바 프레임워크는 내부적으로 프록시를 사용한다.
- 영속성, 트랜잭션, 보안, 캐시, 장애조치 등과 같은 필수적인 애플리케이션 기반 구조를 구현한다.

```java
<beans>
    ...
    <bean id="appDataSource"
        class="org.apache.commons.dbcp.BasicDataSource"
        destroy-method="close"
        p:driverClassName="com.mysql.jdbc.Driver"
        p:url="jdbc:mysql://localhost:3306/mydb"
        p:username="me"/>
    
    <bean id="bankDataAccessObject"
        class="com.example.banking.persistence.BankDataAccessObject"
        p:dataSource-ref="appDataSource"/>
    
    <bean id="bank"
        class="com.example.banking.model.Bank"
        p:dataAccessObject-ref="bankDataAccessObject"/>
    ...
</beans>
```

- 클라이언트는 Bank에 접근하고 있다고 생각하지만 사실은 가장 외곽부터 접근하고 있는 것이다.
- 구조 정의를 위한 xml은 다소 장황하고 읽기 힘들 수는 있지만 Java Proxy보다는 훨씬 간결하다.

```java
/* Code 3-4(Listing 11-5): An EBJ3 Bank EJB */

package com.example.banking.model;

import javax.persistence.*;
import java.util.ArrayList;
import java.util.Collection;

@Entity
@Table(name = "BANKS")
public class Bank implements java.io.Serializable {
    @Id @GeneratedValue(strategy=GenerationType.AUTO)
    private int id;
    
    @Embeddable // Bank의 데이터베이스 행에 '인라인으로 포함된' 객체
    public class Address {
        protected String streetAddr1;
        protected String streetAddr2;
        protected String city;
        protected String state;
        protected String zipCode;
    }
    
    @Embedded
    private Address address;
    @OneToMany(cascade = CascadeType.ALL, fetch = FetchType.EAGER, mappedBy="bank")
    private Collection<Account> accounts = new ArrayList<Account>();
    public int getId() {
        return id;
    }
    
    public void setId(int id) {
        this.id = id;
    }
    
    public void addAccount(Account account) {
        account.setBank(this);
        accounts.add(account);
    }
    
    public Collection<Account> getAccounts() {
        return accounts;
    }
    
    public void setAccounts(Collection<Account> accounts) {
        this.accounts = accounts;
    }
}
```

- 일부 상세한 엔티티 정보는 애너테이션에 포함되어 그대로 남아있지만, 모든 정보가 애너테이션 속에 있으므로 코드 자체는 깔끔하고 깨끗하다.


### AspectJ

- AspectJ는 언어 차원에서 관점을 모듈화 구성으로 지원하는 자바 언어 확장이다.
- 관심사를 관점으로 분리하는 가장 강력한 도구
- 새 도구를 사용하고 새 언어 문법과 사용법을 익혀야 한다는 단점이 있다.
- 최근에 나온 AspectJ ‘애너테이션 폼’은 새로운 도구와 새로운 언어라는 부담을 어느 정도 완화한다.

### 테스트 주도 시스템 아키텍처 구축

- 애플리케이션 도메인 논리를 POJO로 작성할 수 있다면, 즉 코드 수준에서 아키텍처 관심사를 분리할 수 있다면, 진정한 테스트 주도 아키텍처 구축이 가능해진다.
- 그때그때 새로운 기술을 채택해 단순한 아키텍처를 복잡한 아키텍처로 키워갈 수도 있다. 

### 의사 결정을 최적화하라

- 모듈을 나누고 관심사를 분리하면 지엽적인 관리와 결정이 가능해진다.
- 큰 시스템에서는 한 사람이 모든 결정을 내리기 어렵다.
- 최대한 정보를 모아 최선의 결정을 내리기 위해서 가능한 마지막 순간까지 결정을 미루는 방법이 최선이라는 사실을 까먹곤 한다.

### 명백한 가치가 있을 때 표준을 현명하게 사용하라

- 표준을 사용하면 아이디어와 컴포넌트를 재사용하기 쉽고, 적절한 경험을 가진사람을 구하기 쉬우며, 좋은 아이디어를 캡슐화하기 쉽고, 컴포넌트를 엮기 쉽다. 
- 하지만 때로는 표준을 만드는 시간이 너무 오래 걸려 업계가 기다리지 못한다. 

### 시스템에는 DSL(도메인 영역 언어)이 필요하다

- 건축 분야 역시 필수적인 정보를 명료하고 정확하게 전달하는 어휘, 관용구, 패턴이 풍부하다.
- DSL은 간단한 스크립트 언어나 표준 언어로 구현한 API를 가리킨다. 
- 좋은 DSL은 도메인 개념과 그 개념을 구현한 코드 사이에 존재하는 ‘의사소통 간극’을 줄여준다.
- 도메인 전문가가 사용하는 언어로 도메인 논리를 구현하면 도메인을 잘못 구현할 가능성이 줄어든다.