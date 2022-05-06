package app

type DataRepository interface {
	GetStudentById(id int32) (*Student, error)
	GetStudentByLastName(last_name string) ([]*Student, error)
	GetAllStudents() ([]*Student, error)
	GetPopulation() (int, int, error)
}
