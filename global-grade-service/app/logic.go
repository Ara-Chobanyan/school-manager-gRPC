package app

type globalService struct {
	db GlobalRepo
}

func NewGlobalService(data GlobalRepo) GlobalService {
	return &globalService{
		db: data,
	}
}

func (s *globalService) GetGlobalAverage() (GlobalStatus, error) {
	return s.db.GlobalAverage()
}
