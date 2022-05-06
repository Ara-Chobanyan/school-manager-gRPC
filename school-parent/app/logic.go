package app

// Create a struct of our services so we can make methods of it
type parentService struct {
	parentRepo ParentService
}

// Init our interface so it can be injected into
func NewParentService(parentRepo ParentRepository) ParentService {
	return &parentService{
		parentRepo,
	}
}

// creates a new student
func (r *parentService) CreateStudent(student Student) error {
	return r.parentRepo.CreateStudent(student)
}

// Deletes a student
func (r *parentService) DeleteStudent(id int32) error {
	return r.parentRepo.DeleteStudent(id)
}

// Makes edits to a students data
func (r *parentService) EditStudent(student Student) error {
	return r.parentRepo.EditStudent(student)
}
