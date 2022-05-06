package app

type managerService struct {
	managerRepo ManagerRepository
}

func NewManagerService(managerRepo ManagerRepository) *managerService {
	return &managerService{
		managerRepo,
	}
}

func (d *managerService) Behavior() ([]*Student, error) {
	return d.managerRepo.Behavior()
}

func (d *managerService) Exceptional() ([]*Student, error) {
	return d.managerRepo.Exceptional()
}

func (d *managerService) FlagHelper(status bool, who string) error {
	return d.managerRepo.FlagHelper(status, who)
}
