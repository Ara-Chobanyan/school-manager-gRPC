package app

type dataService struct {
	dataRepo DataRepository
	// dataSer  DataService
}

func NewDataService(dataRepo DataRepository /* dataSer DataService */) DataService {
	return &dataService{
		dataRepo,
		// dataSer,
	}
}

func (d *dataService) GetStudentById(id int32) (*Student, error) {
	return d.dataRepo.GetStudentById(id)
}

func (d *dataService) GetStudentByLastName(last_name string) ([]*Student, error) {
	return d.dataRepo.GetStudentByLastName(last_name)
}

func (d *dataService) GetAllStudents() ([]*Student, error) {
	return d.dataRepo.GetAllStudents()
}

func (d *dataService) GetPopulation() (int, int, error) {
	return d.dataRepo.GetPopulation()
}

// func (d *dataService) DoSomething(name *Student) (*Student, error) {
// 	return d.dataSer.DoSomething(name)
// }

// The commented out method/return are going be exerpimented later on.
