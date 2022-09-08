package noednett.OpenLife

import javax.persistence.Table;
import javax.persistence.Column;
import javax.persintence.Entity;
import javax.persistence.GeneratedType;
import javax.persistence.GeneratedValue;
import javax.persistence.id;



@Entity
@Table(name = "Student")
data class Student(
	@Id @GeneratedValue(strategy = GenerationType.IDENTITY) var id: Long,
	@Column(name = "email") val email: String,
	@Column(name = "first_name") val firstName: String,
	@Column(name = "last_name") val lastName: String,
	@Column(name = "course") val course: String,
	@Column(name = "registration_number") val registrationNumber: String
)
	
