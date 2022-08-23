### 목차

- 형식을 맞추는 목적 🥸
- 적절한 행 길이를 유지하라 🙇‍♂️
- 팀 규칙 ✏️
- 밥 아저씨의 형식 규칙 🤔



### 형식을 맞추는 목적 🥸

---

- 구현 스타일과 가독성 수준은 유지보수 용이성과 확장성에 계속 영향을 미친다.
- 코드 형식은 너무나도 중요하므로 맹목적으로 따르면 안된다! 코드 형식은 의사 소통의 일환이다.



### 적절한 행 길이를 유지하라 🙇‍♂️

---

- 일반적으로 큰 파일 보다 작은 파일이 코드를 이해하기 쉽다.
- 소스 파일도 신문 기사와 비슷하게 작성한다!

> - 이름은 간단하면서도 설명이 가능하게 짓는다.
> - 소스 파일 첫 부분은 고차원 개념(중요한 함수 Method)를 구현 하며,  아래로 내려갈수록 의도를 세세하게 묘사한다.
> - 마지막에는 가장 저차원 개념(Method 등을) 구현 한다. 

- 개념은 빈 행으로 분리하라!

> Example) 각 행은 수식이나 절을 나타내고, 일련의 행 묶음은 완결된 생각 하나를 표현한다. 생각 사이는 빈 행을 넣어 분리해야 마땅하다.
>
> - 빈 행을 빼버린 코드는 코드 가독성이 현저하게 떨어져 버린다.



<span style="color: green; font-weight:bold; "> Right </span>

```swift
import UIKit

public class BoldWidget {
  public static var REGEXP: String = ""
  
  public func BoldWidget(parent: ParentWidget, text: String) throws -> Void {
  
	}

	public func render() throws -> String {
  
	}
  
}

```



<span style="color: red; font-weight:bold; "> WRONG </span>

```swift
import UIKit
public class BoldWidget {
  public static var REGEXP = ""
  public func BoldWidget(parent: ParentWidget, text: String) throws -> Void {
    
  }
  public func render() throws -> String {
    
  }
  
}
```



- 세로 밀집도

> 줄바꿈이 개념을 분리한다면 세로 밀집도는 연관성을 의미한다.  즉, 서로 밀접한 코드 행은 세로로 가까이 놓여야 한다는 뜻이다.

```swift
public class ReporterConfig {
  /**
  * 리포터 리스너의 클래스 이름
  */
  private var m_className: String = ""
  
  /**
  * 리포터 리스너의 속성
  */
  
  private var m_properties: Array[Property] = []
  public func addProperty(property: Property) -> Void {
    m_properties.append(property)
  }
  
}
```



- 수직 거리

  > - 서로 밀접한 개념은 세로로 가까이 둬야 한다.
  > - 두 개념이 서로 다른 파일에 속한다면 규칙이 통하지 않는다.
  > - 타당한 근거가 없다면 서로 밀접한 개념은 한 파일에 속해야 마땅하다. 이게 바로(protected 변수를 피해야 하는 이유 중 하나다.)
  > - 같은 파일에 속할 정도로 밀접한 두 개념은 세로 거리로 연관성을 표현한다.

  - 변수선언

    - 변수는 사용하는 위치에 최대한 가까이 선언한다.
    - 지역 변수는 각 함수 맨 처음에 선언 한다.

  - 인스턴스 변수

    - 인스턴스 변수는 클래스 맨 처음에 선언한다.
    - 변수 간에 세로로 거리를 두지 않는다. 잘 설계한 클래스는 메서드가 인스턴스 변수를 사용하기 때문이다.

  - 종속 함수

    - 한 함수가 다른 함수를 호출 한다면 두 함수는 세로로 가까이 배치한다.
    - 가능하다면 호출하는 함수를 호출 되는 함수보다 먼저 배치한다. 그러면 자연스럽게 읽힌다.

    > - 즉 함수의 호출 순서를 코드 선언을 통해 보여주면 읽기가 쉬워진다.

```swift
public class WikiPageResponder: SecureResponder  {
  private var page: WikiPage
  private var pageData: PageData
  private var pageTitle: String
  private var request: Request
  private var crawler: PageCrawler
  
  
  public func makeResponse(context: FitNesseContext, request: Request) throws -> Void {
    var pageName: String = getPageNameOrDefault(request, "FrontPage")
    loadPage(pageName, context)
    
    if page == nil {
      return notFoundResponse(context, reqeust)
    } else {
      return makePageResponse(context)
    }
  }
 
  private func getPageNameOrDefault(request: Reqeust, defaultPageName: String) -> Void {
    var pageName: String = request.getResource()
    
    if StringUtil.isBlank(pageName) {
      pageName = defaultPageName
    }
    
    return pageName
  }
 
  private func loadPage(resource: String, context: FitNesseContext) throws -> Void {
    var path: WikiPagePath = PathParser.parse(resource)
    crawler = context.root.getPageCrawler()
   	crawler.setDeadEndStrategy(VirtualEnabledPageCrawler())
    page = crawler.getPage(context.root, path)
    if page != nil {
      pageData = page.getData()
    }
  }
  
  private func notFoundResponse(context: FitNesseContext, reqeust: Request) throws -> Response {
    return NotFoundResponder().makeResponse(context, reqeust);
  }
  
  private func makePageResponse(context: FitNesseContext) throws -> SimpleResponse {
    var pageTitle = PathParser.render(crawler.getFullPath(page))
    var html: String = makeHtml(context)
    
    var response: SimpleResponse = SimpleResponse()
    
    response.setMaxAge(0)
    response.setContent(html)
    return response
  }
  
  
  
  
}
```





- 개념적 유사성

  - 친화도가 높을수록 코드를 가까이 배치한다.

  >친화도가 높은 요인은 여러가지다. 한 함수가 다른 함수를 호출해 생기는 직접적인 종속성과,  변수와 변수를 사용하는 함수도 한 예다.

- 세로 순서

  - 함수 호출 종속성은 아래 방향으로 유지한다.
  - 호출되는 함수를 호출하는 함수보다 나중에 배치한다. 그러면 소스 코드 모듈이 고차원에서 저차원으로 자연스럽게 내려간다.

>신문 기사와 마찬가지로 가장 중요한 개념을 가장 먼저 표현한다.

- 가로 공백과 밀집도

  - 가로로는 공백을 사용해 밀접한 개념과 느슨한 개념을 표현한다.

- 들여쓰기

  - 범위(Scope)로 이뤄진 계층을 표현하기 위해 우리는 코드를 들여쓴다.
  - 들여쓰는 정도는 계층에서 코드가 자리잡는 수준에 비례한다.

  > 들여쓰기한 파일은 구조가 한눈에 들어온다. 변수, 생성자 함수, 접근자 함수, 메서드가 금방 보인다.





### 팀 규칙 ✏️

---

- 팀은 한 가지 규칙에 합의해야 한다. 그리고 모든 팀원은 그 규칙을 따라야 한다.
- 좋은 소프트웨어 시스템은 읽기 쉬운 문서로 이뤄진다는 사실을 기억하라

> 어디에 괄호를 넣을지, 들여쓰기는 몇 자로 할지, 클래스와 변수와 메서드 이름은 어떻게 지을지 결정 한다.



### 밥 아저씨의 형식 규칙 🤔

---

- 주석을 사용하지 않는다!!
- 한 함수가 다른 함수를 호출 한다면 두 함수는 세로로 가까이 배치한다!!
- 인스 턴스 변수는 상단에 배치한다!!