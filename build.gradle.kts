plugins {
	java
	id("org.springframework.boot") version "3.4.2"
	id("io.spring.dependency-management") version "1.1.7"
	id("com.github.bjornvester.xjc") version "1.8.1"
}

val mockitoAgent = configurations.create("mockitoAgent")

group = "de.dwcaesar"
version = "0.0.1-SNAPSHOT"

java {
	toolchain {
		languageVersion = JavaLanguageVersion.of(21)
	}
}

xjc {
	xjcVersion.set("3.0.2") // jakarta
	outputJavaDir.set(layout.projectDirectory.dir("/src/main/java/xjc"))
	markGenerated.set(true)
}

repositories {
	mavenCentral()
}

dependencies {
	compileOnly("org.projectlombok:lombok")
	annotationProcessor("org.projectlombok:lombok")
	implementation("com.fasterxml.jackson.dataformat:jackson-dataformat-xml")
	implementation("org.springframework.boot:spring-boot-starter")
	implementation("org.springframework.boot:spring-boot-starter-web")
	implementation("org.springframework.boot:spring-boot-starter-actuator")

	testCompileOnly("org.projectlombok:lombok")
	testAnnotationProcessor("org.projectlombok:lombok")
	testImplementation("org.springframework.boot:spring-boot-starter-test")
	// https://javadoc.io/doc/org.mockito/mockito-core/latest/org/mockito/Mockito.html#mockito-instrumentation
	// explicitly setting up instrumentation for inline mocking
	mockitoAgent("org.mockito:mockito-core") { isTransitive = false }
	testRuntimeOnly("org.junit.platform:junit-platform-launcher")

	// improves integration with IDEs, for faster restarts and easier debugging
	developmentOnly("org.springframework.boot:spring-boot-devtools")
}

tasks.withType<Test> {
	jvmArgs = listOf("-javaagent:${mockitoAgent.asPath}")
	useJUnitPlatform()
}
