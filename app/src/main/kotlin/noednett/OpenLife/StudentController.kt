package noednett.OpenLife

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class StudentController(val StudentService: StudentService) {

	@GetMapping("/hello")
	fun helloKotlin(): String {
		return "Hello from the controller in the Kotlin-Spring app"
	}
}
