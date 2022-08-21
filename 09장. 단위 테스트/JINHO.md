## 단위테스트


## TDD 법칙 세 가지
## 깨끗한 테스트 코드 유지하기
## 깨끗한 테스트 코드
## 테스트당 assert 하나
## F.I.R.S.T.

### TDD 법칙 세 가지
- 첫째 법칙: 실패하는 단위 테스트를 작성할 때까지 실제 코드를 작성하지 않는다.
- 둘째 법칙: 컴파일은 실패하지 않으면서 실행이 실패하는 정도로만 단위 테스트를 작성한다.
- 셋째 법칙: 현재 실패하는 테스트를 통과할 정도로만 실제 코드를 작성한다.

### 깨끗한 테스트 코드 유지하기
- 지저분한 테스트코드는 안하니만 못하다.
- 테스트 코드는 실제 코드 못지 않게 중요하고 깨끗하게 작성한다.

### 테스트는 유연성, 유지보수성, 재사용성을 제공한다

-  테스트는 유연성, 유지보수성, 재사용성을 제공한다. 테스트 케이스가 있으면 변경이 쉬워지기 때문이다
- 테스트 코드가 지저분하면 코드를 변경하는 능력이 떨어지며 코드 구조를 개선하는 능력도 떨어진다.


### 깨끗한 테스트 코드

- 깨끗한 코드를 만들려면 가독성이 중요하다.
- 가독성을 높이기 위해선 명료성, 단순성, 풍부한 표현력이 필요하다.

#### 목록 9-1 SerializedPageResponderTest.java
```java
public void testGetPageHieratchyAsXml() throws Exception {
  crawler.addPage(root, PathParser.parse("PageOne"));
  crawler.addPage(root, PathParser.parse("PageOne.ChildOne"));
  crawler.addPage(root, PathParser.parse("PageTwo"));

  request.setResource("root");
  request.addInput("type", "pages");
  Responder responder = new SerializedPageResponder();
  SimpleResponse response =
    (SimpleResponse) responder.makeResponse(new FitNesseContext(root), request);
  String xml = response.getContent();

  assertEquals("text/xml", response.getContentType());
  assertSubString("<name>PageOne</name>", xml);
  assertSubString("<name>PageTwo</name>", xml);
  assertSubString("<name>ChildOne</name>", xml);
}

public void testGetPageHieratchyAsXmlDoesntContainSymbolicLinks() throws Exception {
  WikiPage pageOne = crawler.addPage(root, PathParser.parse("PageOne"));
  crawler.addPage(root, PathParser.parse("PageOne.ChildOne"));
  crawler.addPage(root, PathParser.parse("PageTwo"));

  PageData data = pageOne.getData();
  WikiPageProperties properties = data.getProperties();
  WikiPageProperty symLinks = properties.set(SymbolicPage.PROPERTY_NAME);
  symLinks.set("SymPage", "PageTwo");
  pageOne.commit(data);

  request.setResource("root");
  request.addInput("type", "pages");
  Responder responder = new SerializedPageResponder();
  SimpleResponse response =
    (SimpleResponse) responder.makeResponse(new FitNesseContext(root), request);
  String xml = response.getContent();

  assertEquals("text/xml", response.getContentType());
  assertSubString("<name>PageOne</name>", xml);
  assertSubString("<name>PageTwo</name>", xml);
  assertSubString("<name>ChildOne</name>", xml);
  assertNotSubString("SymPage", xml);
}

public void testGetDataAsHtml() throws Exception {
  crawler.addPage(root, PathParser.parse("TestPageOne"), "test page");

  request.setResource("TestPageOne"); request.addInput("type", "data");
  Responder responder = new SerializedPageResponder();
  SimpleResponse response =
    (SimpleResponse) responder.makeResponse(new FitNesseContext(root), request);
  String xml = response.getContent();

  assertEquals("text/xml", response.getContentType());
  assertSubString("test page", xml);
  assertSubString("<Test", xml);
}
```

#### 목록 9-2 (refactored)
```java
public void testGetPageHierarchyAsXml() throws Exception {
  makePages("PageOne", "PageOne.ChildOne", "PageTwo");

  submitRequest("root", "type:pages");

  assertResponseIsXML();
  assertResponseContains(
    "<name>PageOne</name>", "<name>PageTwo</name>", "<name>ChildOne</name>");
}

public void testSymbolicLinksAreNotInXmlPageHierarchy() throws Exception {
  WikiPage page = makePage("PageOne");
  makePages("PageOne.ChildOne", "PageTwo");

  addLinkTo(page, "PageTwo", "SymPage");

  submitRequest("root", "type:pages");

  assertResponseIsXML();
  assertResponseContains(
    "<name>PageOne</name>", "<name>PageTwo</name>", "<name>ChildOne</name>");
  assertResponseDoesNotContain("SymPage");
}

public void testGetDataAsXml() throws Exception {
  makePageWithContent("TestPageOne", "test page");

  submitRequest("TestPageOne", "type:data");

  assertResponseIsXML();
  assertResponseContains("test page", "<Test");
}
```

- 9-1 코드는 읽는 사람을 고려하지 않는다.
- BUILD-OPERATE-CHECK 패턴2이 위와 같은 테스트 구조에 적합하다.

### 도메인에 특화된 테스트 언어

- API 위에다 함수와 유틸리티를 구현한 후 그 함수와 유틸리티를 사용하므로 테스트 코드를 짜기도 읽기도 쉬워진다. 
- 테스트를 구현하는 당사자와 나중에 테스트를 읽어볼 독자를 도와주는 테스트 언어다.

### 이중 표준

- 테스트 API코드에 적용하는 표준은 실제 코드에 적용하는 표준과 확실히 다르다.
- 실제 환경과 달리 테스트 환경은 자원이 제한적일 가능성이 낮고,테스트 코드가 실제 코드만큼 효율적일 필요는 없다.
- 이중표준의 본질은 실제 환경에서는 절대로 안 되지만 테스트 환경에서는 전혀 문제없는 방식이 있다.

### 테스트 당 assert 하나

- JUnit으로 테스트 코드를 짤 때 함수마다 assert를 단 하나만 사용해야 하면 가독성이 빠르고 이해하기 쉽다.
- assert의문의 개수는 최대한 줄이는게 좋다.

### 테스트당 개념 하나

- 테스트 함수마다 한 개념만 테스트하라
- 개념당 assert문을 하나로 줄여라

### F.I.R.S.T

- 깨끗한 테스트는 다음 다섯 가지 규칙을 따르는데, 각 규칙에서 첫 글자를 따오면 FIRST가 된다. 

**빠르게Fast:**  
테스트는 빨리 돌아야 한다.

**독립적으로Independent:**  
각 테스트는 독립적으로 그리고 어떤 순서로 실행해도 괜찮아야 한다. 

**반복가능하게Repeatable:**  
테스트는 어떤 환경에서도 반복 가능해야 한다.

**자가검증하는Self-Validating:**  
테스트는 bool값으로 결과를 내야 한다. 성공 아니면 실패다.
