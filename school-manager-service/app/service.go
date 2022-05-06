package app

type ServiceManager interface {
	Behavior() ([]*Student, error)
	Exceptional() ([]*Student, error)
	FlagHelper(bool, string) error
}
