import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
	id("org.springframework.boot") version "2.7.3"
	id("io.spring.dependency-management") version "1.0.13.RELEASE"
	kotlin("jvm") version "1.6.21"
	kotlin("plugin.spring") version "1.6.21"

	application
}

group = "noednett"
version = "0.0.1-SNAPSHOT"
java.sourceCompatibility = JavaVersion.VERSION_17

repositories {
	mavenCentral()
}

dependencies {
	implementation("org.springframework.boot:spring-boot-starter-web")
	implementation("com.fasterxml.jackson.module:jackson-module-kotlin")
	implementation("org.jetbrains.kotlin:kotlin-reflect")
	implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")

	//adding dependencies for flyway
	implementation("org.flywaydb:flyway-core")

	//adding dependencies for postgres and the required drivers to use it with spring boot
	implementation("org.springframework.boot:spring-boot-starter-jdbc")
	implementation("org.springframework.boot:spring-boot-starter-web")
	implementation("org.postgresql:postgresql")

	testImplementation("org.springframework.boot:spring-boot-starter-test")

}

tasks.withType<KotlinCompile> {
	kotlinOptions {
		freeCompilerArgs = listOf("-Xjsr305=strict")
		jvmTarget = "17"
	}
}

//disabling '-plain.jar' version in build.
tasks.getByName<Jar>("jar") {
	enabled = false
}

tasks.withType<Test> {
	useJUnitPlatform()
}

application {
	mainClass.set("noednett.OpenLife.OpenLifeApplicationKt")
}


