package app

type ParentRepository interface {
	CreateStudent(student Student) error
	DeleteStudent(id int32) error
	EditStudent(student Student) error
}
