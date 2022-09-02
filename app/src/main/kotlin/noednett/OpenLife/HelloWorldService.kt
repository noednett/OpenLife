package noednett.OpenLife

import org.springframework.stereotype.Service

@Service
class HelloService {

	fun getHello(): String {
		return "Hello from the spring boot-kotlin app"
	}
}
