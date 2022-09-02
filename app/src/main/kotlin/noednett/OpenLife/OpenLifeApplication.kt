package noednett.OpenLife

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.boot.SpringApplication
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController


@SpringBootApplication
class OpenLifeApplication

fun main(args: Array<String>) {
	SpringApplication.run(OpenLifeApplication::class.java, *args)
	//runApplication<OpenLifeApplication>(*args)
}
