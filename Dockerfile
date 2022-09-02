#Temporary container building the app with gradle.
FROM gradle:7-jdk18-alpine AS BUILD
ENV APP_HOME=/usr/local/dev/OpenLife/

WORKDIR $APP_HOME

COPY gradle ./gradle/
COPY settings.gradle.kts gradlew ./
COPY app ./app/

RUN ./gradlew build 

#Creating the openjdk container for the app
FROM openjdk:18-alpine
ENV APP_HOME=/usr/local/dev/OpenLife/

WORKDIR $APP_HOME

#Define the version of the app, this should better be done with an .env file.
ENV APP_VERSION=0.0.1-SNAPSHOT

#honestly, i dont know where a good place is to get the final app. altough it seems logical to put it in /var/www
COPY --from=BUILD $APP_HOME/app/build/libs/app-0.0.1-SNAPSHOT.jar $APP_HOME/app/build/libs/app-0.0.1-SNAPSHOT.jar 

#RUN addgroup -S spring && adduser -S spring -G spring
#USER spring:spring

ENTRYPOINT ["java", "-jar", "/usr/local/dev/OpenLife/app/build/libs/app-0.0.1-SNAPSHOT.jar"]

