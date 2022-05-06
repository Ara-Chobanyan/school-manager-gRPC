package app

type ManagerRepository interface {
	Behavior() ([]*Student, error)
	Exceptional() ([]*Student, error)
	FlagHelper(bool, string) error
}
