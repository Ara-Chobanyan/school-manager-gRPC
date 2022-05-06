package app

type ParentService interface {
	CreateStudent(student Student) error
	DeleteStudent(id int32) error
	EditStudent(student Student) error
}
