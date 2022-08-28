# 5장. 형식을 맞추기

## 형식을 맞추는 목적

- 너무나도 당연하다. 형식을 맞추는거는 의사소통을 하는데 있어서 가장 중요하다.
- 우리가 대화를 할때 형식이랑 문법을 맞추듯, 외국어도 그렇고..
- 이해 가능한 수준의 문법으로 말해야 이해가 쉬워진다고 생각하면 된다. 


## 적절한 행 길이를 유지하라 (코드의 세로 길이)

- 대부분 200줄 정도인 파일로도 커다란 시스템을 구축할 수 있다는 사실이다. 
- 자바의 경우, 파일 단위가 클래스를 기준으로 구분되기 때문에, 파일의 크기는 클래스에 크기에 비례한다.(여기서는 라인수를 기준으로 정함)

#### 신문 기사처럼 작성하라

- 헤드라인부터 top-down 방식으로 읽기 좋게 구성
- 목적을 알 수 있는 제목부터 상세한 구현부 까지 서술하라는 방식으로 보인다.

#### 개념은 빈 행으로 분리하
```java
// 빈 행을 넣지 않을 경우
package fitnesse.wikitext.widgets;
import java.util.regex.*;
public class BoldWidget extends ParentWidget {
	public static final String REGEXP = "'''.+?'''";
	private static final Pattern pattern = Pattern.compile("'''(.+?)'''",
		Pattern.MULTILINE + Pattern.DOTALL);
	public BoldWidget(ParentWidget parent, String text) throws Exception {
		super(parent);
		Matcher match = pattern.matcher(text); match.find(); 
		addChildWidgets(match.group(1));}
	public String render() throws Exception { 
		StringBuffer html = new StringBuffer("<b>"); 		
		html.append(childHtml()).append("</b>"); 
		return html.toString();
	} 
}
```

```java
// 빈 행을 넣을 경우
package fitnesse.wikitext.widgets;

import java.util.regex.*;

public class BoldWidget extends ParentWidget {
	public static final String REGEXP = "'''.+?'''";
	private static final Pattern pattern = Pattern.compile("'''(.+?)'''", 
		Pattern.MULTILINE + Pattern.DOTALL
	);
	
	public BoldWidget(ParentWidget parent, String text) throws Exception { 
		super(parent);
		Matcher match = pattern.matcher(text);
		match.find();
		addChildWidgets(match.group(1)); 
	}
	
	public String render() throws Exception { 
		StringBuffer html = new StringBuffer("<b>"); 
		html.append(childHtml()).append("</b>"); 
		return html.toString();
	} 
}
```

- 개념별로 빈 행을 두어 코드끼리 그룹을 지을 수 있게 한다.
- 글을 작성시 문단 단위로 글을 나누는 것과 같은 맥락이라 생각한다.

#### 세로 밀집도

```java
public class ReporterConfig {
	/**
	* The class name of the reporter listener 
	*/
	private String m_className;
	
	/**
	* The properties of the reporter listener 
	*/
	private List<Property> m_properties = new ArrayList<Property>();
	public void addProperty(Property property) { 
		m_properties.add(property);
	}
```

```java
public class ReporterConfig {
	private String m_className;
	private List<Property> m_properties = new ArrayList<Property>();
	
	public void addProperty(Property property) { 
		m_properties.add(property);
	}
```

- 코드 개념별로 빈 행을 두어 관리하듯, 관련 있는 코드들은 붙여 작성한다.(관계가 있음을 나타낸다.)

#### 수직 거리 

- 변수나 함수의 정의가 사용되는 고승로 부터 얼마나 멀리 정의되어있는지를 표현하는 것이다.
- 현재 맥락과 관련된 변수, 함수의 선언을 가까운 곳에 둬서 커서를 최소한으로 움직이면서 맥락을 확인할 수 있게 배치하란 뜻이다.  

###### 변수선언
```java
private static void readPreferences() {
	InputStream is = null;
	try {
		is = new FileInputStream(getPreferencesFile()); 
		setPreferences(new Properties(getPreferences())); 
		getPreferences().load(is);
	} catch (IOException e) { 
		try {
			if (is != null) 
				is.close();
		} catch (IOException e1) {
		} 
	}
}
```
- 지역변수는 사용위치로부터 최대한 가까이 선언.

```java
public int countTestCases() { 
	int count = 0;
	for (Test each : tests)
		count += each.countTestCases(); 
	return count;
}
```

- 루프제어 변수는 루프 문 내부에 선언.

```java
for (XmlTest test : m_suite.getTests()) {
	TestRunner tr = m_runnerFactory.newTestRunner(this, test);
	tr.addListener(m_textReporter); 
	m_testRunners.add(tr);

	invoker = tr.getInvoker();
	
	for (ITestNGMethod m : tr.getBeforeSuiteMethods()) { 
		beforeSuiteMethods.put(m.getMethod(), m);
	}

	for (ITestNGMethod m : tr.getAfterSuiteMethods()) { 
		afterSuiteMethods.put(m.getMethod(), m);
	} 
}
```

- 긴 함수에서는 블록 상단 또는 루프 직전에 변수를 선언 할 수도 있다.

###### 인스턴스 변수

```java
public class TestSuite implements Test {
	static public Test createTest(Class<? extends TestCase> theClass,
									String name) {
		... 
	}

	public static Constructor<? extends TestCase> 
	getTestConstructor(Class<? extends TestCase> theClass) 
	throws NoSuchMethodException {
		... 
	}

	public static Test warning(final String message) { 
		...
	}
	
	private static String exceptionToString(Throwable t) { 
		...
	}
	
    // 글에 작성된 코드는 중간쯤에 인스턴스 변수를 선언한다.
	private String fName;

	private Vector<Test> fTests= new Vector<Test>(10);

	public TestSuite() { }
	
	public TestSuite(final Class<? extends TestCase> theClass) { 
		...
	}

	public TestSuite(Class<? extends TestCase> theClass, String name) { 
		...
	}
	
	... ... ... ... ...
}
```

- 클래스 맨 처음에 선언한다.
- 변수 간에 세로로 공백을 두지 않는다.
- C++의 경우에는 마지막에 선언하는 것이 일반적이다. 어느 곳이든 잘 알려진 위치에 인스턴스 변수를 모으는 것이 중요하다.

###### 종속 함수

```java
public class WikiPageResponder implements SecureResponder { 
	protected WikiPage page;
	protected PageData pageData;
	protected String pageTitle;
	protected Request request; 
	protected PageCrawler crawler;
	
	public Response makeResponse(FitNesseContext context, Request request) throws Exception {
		String pageName = getPageNameOrDefault(request, "FrontPage");
		loadPage(pageName, context); 
		if (page == null)
			return notFoundResponse(context, request); 
		else
			return makePageResponse(context); 
		}

	private String getPageNameOrDefault(Request request, String defaultPageName) {
		String pageName = request.getResource(); 
		if (StringUtil.isBlank(pageName))
			pageName = defaultPageName;

		return pageName; 
	}
	
	protected void loadPage(String resource, FitNesseContext context)
		throws Exception {
		WikiPagePath path = PathParser.parse(resource);
		crawler = context.root.getPageCrawler();
		crawler.setDeadEndStrategy(new VirtualEnabledPageCrawler()); 
		page = crawler.getPage(context.root, path);
		if (page != null)
			pageData = page.getData();
	}
	
	private Response notFoundResponse(FitNesseContext context, Request request)
		throws Exception {
		return new NotFoundResponder().makeResponse(context, request);
	}
	
	private SimpleResponse makePageResponse(FitNesseContext context)
		throws Exception {
		pageTitle = PathParser.render(crawler.getFullPath(page)); 
		String html = makeHtml(context);
		SimpleResponse response = new SimpleResponse(); 
		response.setMaxAge(0); 
		response.setContent(html);
		return response;
	} 
...
```

- 한 함수가 다른 함수를 호출한다면 두 함수는 세로로 가까이 배치한다.
- 가능하면 호출되는 함수를 호출하는 함수보다 뒤에 배치한다.

###### 개념의 유사성

```java
public class Assert {
	static public void assertTrue(String message, boolean condition) {
		if (!condition) 
			fail(message);
	}

	static public void assertTrue(boolean condition) { 
		assertTrue(null, condition);
	}

	static public void assertFalse(String message, boolean condition) { 
		assertTrue(message, !condition);
	}
	
	static public void assertFalse(boolean condition) { 
		assertFalse(null, condition);
	} 
...
```

- 개념적인 친화도가 높을 수록 코드를 서로 가까이 배치한다.  

#### 세로 순서

- 종속성에 맞춰 고차원 ==> 저차원으로 쓴다.
- 관련 있는 순서부터 코드를 작성한다.


## 가로 형식 맞추기

- 코드를 볼때 커서는 위 - 아래만으로도 충분하다.
- 가독성이 너무 떨어진다.

#### 가로 공백과 밀집도
```java
private void measureLine(String line) { 
	lineCount++;
	int lineSize = line.length();
	totalChars += lineSize; 
	
	lineWidthHistogram.addLine(lineSize, lineCount);
	recordWidestLine(lineSize);
}
```

- 연산자 사이에는 공백을 두어 무슨 연산이 일어나는지 알기 쉽게 한다.
- 함수 이름가 괗로 사이에는 공백을 없애 밀접함을 보여준다.

#### 가로 정렬
```java
public class CodeAnalyzer implements JavaFileAnalysis {
	private int lineCount;
	private int maxLineWidth;
	private int widestLineNumber;
	private LineWidthHistogram lineWidthHistogram;
	private int totalChars;

	public CodeAnalyzer() {
		lineWidthHistogram = new LineWidthHistogram();
	}

	public static List<File> findJavaFiles(File parentDirectory) {
		List<File> files = new ArrayList<File>();
		findJavaFiles(parentDirectory, files);
		return files;
	}

	private static void findJavaFiles(File parentDirectory, List<File> files) {
		for (File file : parentDirectory.listFiles()) {
			if (file.getName().endsWith(".java"))
				files.add(file);
			else if (file.isDirectory())
				findJavaFiles(file, files);
			}
	}

	public void analyzeFile(File javaFile) throws Exception {
		BufferedReader br = new BufferedReader(new FileReader(javaFile));
		String line;
		while ((line = br.readLine()) != null)
			measureLine(line);
	}

	private void measureLine(String line) {
		lineCount++;
		int lineSize = line.length();
		totalChars += lineSize;
		lineWidthHistogram.addLine(lineSize, lineCount);
		recordWidestLine(lineSize);
	}

	private void recordWidestLine(int lineSize) {
		if (lineSize > maxLineWidth) {
			maxLineWidth = lineSize;
			widestLineNumber = lineCount;
		}
	}

	public int getLineCount() {
		return lineCount;
	}

	public int getMaxLineWidth() {
		return maxLineWidth;
	}

	public int getWidestLineNumber() {
		return widestLineNumber;
	}

	public LineWidthHistogram getLineWidthHistogram() {
		return lineWidthHistogram;
	}

	public double getMeanLineWidth() {
		return (double)totalChars/lineCount;
	}

	public int getMedianLineWidth() {
		Integer[] sortedWidths = getSortedWidths();
		int cumulativeLineCount = 0;
		for (int width : sortedWidths) {
			cumulativeLineCount += lineCountForWidth(width);
			if (cumulativeLineCount > lineCount/2)
				return width;
		}
		throw new Error("Cannot get here");
	}
	private int lineCountForWidth(int width) {
		return lineWidthHistogram.getLinesforWidth(width).size();
	}
	private Integer[] getSortedWidths() {
		Set<Integer> widths = lineWidthHistogram.getWidths();
		Integer[] sortedWidths = (widths.toArray(new Integer[0]));
		Arrays.sort(sortedWidths);
		return sortedWidths;
	}
}


```
- 정렬하기 어렵고 효율적이지도 않다.
- 위에 공백과 밀집도에 따라 변수의 관계성이 약해보인다.

#### 들여쓰기  

- 가독성이 달라진다. 개발자 스타일마다 다른듯한다.(2자, 4자, 8자)

## 팀 규칙

- 현재 업무에 사용하는 형식을 따르는 것이 가장 중요하다.
- 코드를 통해 의사소통을 할 때 서로 지정도니 규정대로 하는 것이 좋다.